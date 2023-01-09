package etsp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

func Logon(userLogin User) (string, error) {

	bytesRepresentation, err := json.Marshal(userLogin)
	if err != nil {
		return "", err
	}

	fmt.Println("Request:", string(bytesRepresentation))

	resp, err := http.Post("https://ws.etsp.ru/v2/json/Security.svc/Logon", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
