package etsp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LogonResponse struct {
	Errors   stringOrNull `json:"Errors"`
	Success  bool         `json:"Success"`
	Warnings stringOrNull `json:"Warnings"`
	Data     stringOrNull `json:"Data"`
}

func (user User) Logon() (LogonResponse, error) {
	bytesRepresentation, err := json.Marshal(user)
	if err != nil {
		return LogonResponse{}, err
	}

	resp, err := http.Post(URL+"/v2/json/Security.svc/Logon", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return LogonResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LogonResponse{}, err
	}

	fmt.Println(string(body))

	// Распарсить данные
	var LogonRes LogonResponse
	responseErrorUnmarshal := json.Unmarshal(body, &LogonRes)
	if responseErrorUnmarshal != nil {
		return LogonResponse{}, responseErrorUnmarshal
	}

	return LogonRes, nil
}
