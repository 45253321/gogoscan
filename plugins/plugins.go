package plugins

import "fmt"

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

type verifyHandler func(burstCase BurstCase)bool


var pluginMap = map[string]verifyHandler{"mysql": MysqlVerify, "ssh": SshVerify}

func SearchHandler(protocol string) (verifyHandler, bool){
	handler, ok := pluginMap[protocol]
	if ok {
		return handler, true
	}else{
		return nil, false
	}
}