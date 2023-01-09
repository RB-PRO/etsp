package etsp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Структура запроса на деавторизацию
type logoutRequest struct {
	HashSession string `json:"HashSession"` //Хэш-ключ сессии
}

// Структура ответа на авторизацию [Logout]
//
// [Logout]: https://ws.etsp.ru/Help/v2/Security/Logout.aspx
type LogoutResponse struct {
	Errors   []string `json:"Errors"`
	Success  bool     `json:"Success"`
	Warnings []string `json:"Warnings"`
}

func (user User) Logout() (LogoutResponse, error) {
	// Запаковать в json
	bytesRepresentation, err := json.Marshal(logoutRequest{user.HashSession})
	if err != nil {
		return LogoutResponse{}, err
	}

	// Выполнить запрос
	resp, err := http.Post(URL+"/v2/json/Security.svc/Logout", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return LogoutResponse{}, err
	}

	// Преобразовать данные в массив byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LogoutResponse{}, err
	}

	// Распарсить данные
	var LogoutRes LogoutResponse
	responseErrorUnmarshal := json.Unmarshal(body, &LogoutRes)
	if responseErrorUnmarshal != nil {
		return LogoutResponse{}, responseErrorUnmarshal
	}

	// Проверка на отсутствие авторизации
	if !LogoutRes.Success {
		return LogoutResponse{}, errors.New(LogoutRes.Errors[0])
	}

	return LogoutRes, nil
}
