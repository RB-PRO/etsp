package etsp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Структура ответа на авторизацию [Logon]
//
// [Logon]: https://ws.etsp.ru/Help/v2/Security/Logon.aspx
type LogonResponse struct {
	Errors   []string `json:"Errors"`
	Success  bool     `json:"Success"`
	Warnings []string `json:"Warnings"`
	Data     string   `json:"Data"`
}

func (user *User) Logon() (LogonResponse, error) {
	// Запаковать в json
	bytesRepresentation, err := json.Marshal(user)
	if err != nil {
		return LogonResponse{}, err
	}

	// Выполнить запрос
	resp, err := http.Post(URL+"/v2/json/Security.svc/Logon", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return LogonResponse{}, err
	}

	// Преобразовать данные в массив byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LogonResponse{}, err
	}

	// Распарсить данные
	var LogonRes LogonResponse
	responseErrorUnmarshal := json.Unmarshal(body, &LogonRes)
	if responseErrorUnmarshal != nil {
		return LogonResponse{}, responseErrorUnmarshal
	}

	// Проверка на отсутствие авторизации
	if !LogonRes.Success {
		return LogonResponse{}, errors.New(LogonRes.Errors[0])
	}

	// Заполнение HashSession
	user.HashSession = LogonRes.Data

	return LogonRes, nil
}
