package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

func Logon(userLogin User) string {

	bytesRepresentation, err := json.Marshal(userLogin)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://ws.etsp.ru/v2/json/Security.svc/Logon", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}
