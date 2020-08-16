package main

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	err := _sshConnect("10.10.10.72", "root", "roo11", 22)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("connnect success!!")
	}
}

func BenchmarkSSHScan_Scan(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			ssc := SSHScan{}
			ssc.Ip="10.10.10.232"
			ssc.Port=47013
			ssc.UsernamePath="./resource/username_ssh.txt"
			ssc.PassWordPath="./resource/password_ssh.txt"
			ssc.Concurrent=20
			ssc.Scan()
			fmt.Println("Have Found Password Num:  ", len(ssc.PasswordBurst))
		}

	})

}