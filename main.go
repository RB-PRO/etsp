package main

import "fmt"

const URL string = "https://ws.etsp.ru"

func main() {

	//Create user struct which need to post.
	userLogin := User{
		Login:    "ya.detal@yandex.ru_IP",
		Password: "12345W",
	}
	fmt.Println(Logon(userLogin))
}
