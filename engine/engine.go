package engine

import (
	"fmt"
	"gogoscan/plugins"
	"gogoscan/utils"
)

type ScanEngine struct {
	Ip string
	Port int
	Protocol string
	UsernameTextPath string
	PasswordTextPath string
	Concurrency int
	scanHandler func(burstCase plugins.BurstCase) bool
	PasswordBurst []plugins.BurstCase
}

// 协程处理模型
func engine(input <-chan plugins.BurstCase, result chan <- plugins.BurstCase, handle func(burstCase plugins.BurstCase)bool ){

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

// 核心执行引擎
func (self *ScanEngine) Run() error {

	handler, ok := plugins.SearchHandler(self.Protocol)
	if ok {
		self.scanHandler = handler
	}else{
		return nil
	}

	usernameLines := utils.ReadLines(self.UsernameTextPath)
	passwordLines := utils.ReadLines(self.PasswordTextPath)
	burstUserCase := utils.ItertoolsProduct(*usernameLines, *passwordLines)

	result := make(chan plugins.BurstCase, len(*burstUserCase))
	burstCaseChan := make(chan plugins.BurstCase, self.Concurrency)

	for i := 0; i<self.Concurrency; i++{
		go engine(burstCaseChan, result, self.scanHandler)
	}

	for _, userCase := range *burstUserCase{
		username, password := userCase[0], userCase[1]
		burstCaseChan <- plugins.BurstCase{Ip: self.Ip, Port:self.Port, Username:username, Password:password, Success:false}
	}

	for i:=0; i<len(*burstUserCase); i++{
		burstCase := <- result
		if burstCase.Success == true{
			self.PasswordBurst = append(self.PasswordBurst, burstCase)
		}
	}

	return nil
}

