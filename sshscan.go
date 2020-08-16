package main

import (
	"fmt"
	"gogoscan/resource"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)


type SSHScan struct {
	Ip string
	Port int
	UsernamePath string
	PassWordPath string
	Concurrent int
	PasswordBurst []BurstCase
}

type BurstCase struct {
	Ip string
	Port int
	Username string
	Password string
	Success bool
}

func (bc BurstCase) String() string {
	return fmt.Sprintf("<ip: %v, port: %v username: %v, password: %v, success: %v>",
		bc.Ip, bc.Port, bc.Username, bc.Password, bc.Success)
}

func _sshConnect(host, user, password string, port int) error{

	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		err          error
	)

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 5 * time.Second,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if _, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return err
	}

	return nil
}
// 消费者
func sshConnect(input <-chan BurstCase, result chan <- BurstCase){

	for {
		burstCase := <- input
		// 进行ssh请求
		err := _sshConnect(burstCase.Ip, burstCase.Username, burstCase.Password, burstCase.Port)
		if err != nil{
			burstCase.Success = false
		}else{
			burstCase.Success = true
		}
		fmt.Println(burstCase)
		result <- burstCase
	}

}

func (self *SSHScan) Scan(){
	usernameLines := resource.ReadLines(self.UsernamePath)
	passwordLines := resource.ReadLines(self.PassWordPath)
	burstUserCase := ItertoolsProduct(*usernameLines, *passwordLines)


	result := make(chan BurstCase, len(*burstUserCase))
	burstCaseChan := make(chan BurstCase, self.Concurrent)

	for i := 0; i<self.Concurrent; i++{
		go sshConnect(burstCaseChan, result)
	}

	for _, userCase := range *burstUserCase{
		username, password := userCase[0], userCase[1]
		burstCaseChan <- BurstCase{Ip:self.Ip, Port:self.Port, Username:username, Password:password, Success:false}
	}

	for i:=0; i<len(*burstUserCase); i++{
		burstCase := <- result
		if burstCase.Success == true{
			self.PasswordBurst = append(self.PasswordBurst, burstCase)
		}
	}
}

func main()  {
	sc := SSHScan{Ip:"10.10.10.232",
		Port:47013,
		UsernamePath:"./resource/username_ssh.txt",
		PassWordPath:"./resource/password_ssh.txt",
		Concurrent:50,
		PasswordBurst: []BurstCase{}}
	sc.Scan()
	fmt.Println("The ParseResult: ", sc.PasswordBurst)
}