package plugins

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

// ssh验证引擎
func SshVerify(burstCase BurstCase) bool{

	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		err          error
	)

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(burstCase.Password))
	clientConfig = &ssh.ClientConfig{
		User:    burstCase.Username,
		Auth:    auth,
		Timeout: 5 * time.Second,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", burstCase.Ip, burstCase.Port)
	if _, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return false
	}
	return true
}