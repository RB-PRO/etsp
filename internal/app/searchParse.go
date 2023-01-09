package app

import (
	"fmt"

	"github.com/rb-pro/etsp/pkg/etsp"
)

const URL string = "https://ws.etsp.ru"

func Run() {

	//Create user struct which need to post.
	userLogin := etsp.User{
		Login:    "ya.detal@yandex.ru_IP",
		Password: "12345W",
	}

	Response, _ := etsp.Logon(userLogin)
	fmt.Print("Response", Response)
}
