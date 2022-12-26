package main

import "fmt"

const URL string = "https://ws.etsp.ru/api-docs/swagger.json"

func main() {

	//Create user struct which need to post.
	userLogin := User{
		Login:    "ya.detal@yandex.ru_hbsk",
		Password: "detal1",
	}
	fmt.Println(Logon(userLogin))
}
