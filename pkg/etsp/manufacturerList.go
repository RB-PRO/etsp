package etsp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Структура запроса на поиск
type ManufacturerList struct {
	HashSession string `json:"HashSession"` // Хэш-ключ сессии
}

// Структура ответа на метод простого поска [ManufacturerList]
//
// [ManufacturerList]: https://ws.etsp.ru/Help/v2/Search/ManufacturerList.aspx
type ManufacturerListResponse struct {
	Errors   []string `json:"Errors"`
	Success  bool     `json:"Success"`
	Warnings []string `json:"Warnings"`
	Data     []struct {
		Id   int    `json:"Id"`   // Идентификатор производителя
		Name string `json:"Name"` // Идентификатор производителя
	} `json:"Data"`
	// * Доступ выдает ответственный менеджер.
	//** Флаги IsSklad, IsShops, IsShipment, IsOutside требуют специального разрешения, т.к. их расчет сказывается на производительности сервера. Если нет разрешения, то будет установлено значение "0" независимо от наличия на складе. В таком случае для просмотра остатков необходимо воспользоваться методом GetPartsRemainsByCode или GetPartsRemainsBySkubaNumber сервиса PartsRemains.
}

func (user User) ManufacturerList() (ManufacturerListResponse, error) {
	// Запаковать в json
	bytesRepresentation, err := json.Marshal(ManufacturerList{HashSession: user.HashSession})
	if err != nil {
		return ManufacturerListResponse{}, err
	}

	// Выполнить запрос
	resp, err := http.Post(URL+"/v2/json/Search.svc/ManufacturerList", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return ManufacturerListResponse{}, err
	}

	// Преобразовать данные в массив byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ManufacturerListResponse{}, err
	}

	//fmt.Println(string(body))

	// Распарсить данные
	var SearchBasicRes ManufacturerListResponse
	responseErrorUnmarshal := json.Unmarshal(body, &SearchBasicRes)
	if responseErrorUnmarshal != nil {
		return ManufacturerListResponse{}, responseErrorUnmarshal
	}

	// Проверка на отсутствие данных
	if !SearchBasicRes.Success {
		return ManufacturerListResponse{}, errors.New(SearchBasicRes.Errors[0])
	}

	return SearchBasicRes, nil
}
