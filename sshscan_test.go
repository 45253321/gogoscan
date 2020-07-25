package main

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	_, err := _sshConnect("10.10.10.72", "root", "roo11", 22)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("connnect success!!")
	}
}