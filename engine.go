package main

import (
	"fmt"
	"gogoscan/resource"
)

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

type ScanEngine struct {
	Ip string
	Port int
	Protocol string
	UsernameTextPath string
	PasswordTextPath string
	Concurrency int
	scanHandler func(burstCase BurstCase) bool
	PasswordBurst []BurstCase
}

// 消费者
func engine(input <-chan BurstCase, result chan <- BurstCase, handle func(burstCase BurstCase)bool ){

	for {
		burstCase := <- input
		// 进行ssh请求
		ok := handle(burstCase)
		if ok {
			burstCase.Success = true
		}else{
			burstCase.Success = false
		}
		fmt.Println(burstCase)
		result <- burstCase
	}

}

func (self *ScanEngine) Run() error {

	if self.Protocol == "ssh"{
		self.scanHandler = SshVerify
	}

	usernameLines := resource.ReadLines(self.UsernameTextPath)
	passwordLines := resource.ReadLines(self.PasswordTextPath)
	burstUserCase := ItertoolsProduct(*usernameLines, *passwordLines)

	result := make(chan BurstCase, len(*burstUserCase))
	burstCaseChan := make(chan BurstCase, self.Concurrency)

	for i := 0; i<self.Concurrency; i++{
		go engine(burstCaseChan, result, self.scanHandler)
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

	return nil
}

