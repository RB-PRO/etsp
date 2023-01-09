package app

import (
	"fmt"

	"github.com/RB-PRO/etsp/pkg/etsp"
)

func Run() {

	user := etsp.User{
		Login:    "ya.del@andex.ru_IP",
		Password: "12345W",
	}

	ResponseLogon, _ := user.Logon()

	fmt.Println(ResponseLogon)
}
