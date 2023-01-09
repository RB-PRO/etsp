package etsp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Структура запроса на поиск
type PartAttendantByCodeRequest struct {
	Code        string `json:"Code2"`       // Код детали
	HashSession string `json:"HashSession"` // Хэш-ключ сессии
}

// Структура ответа на метод простого поска [PartAttendantByCode]
//
// [PartAttendantByCode]: https://ws.etsp.ru/Help/v2/Search/PartAttendantByCode.aspx
type PartAttendantByCodeResponse struct {
	Errors   []string `json:"Errors"`
	Success  bool     `json:"Success"`
	Warnings []string `json:"Warnings"`
	Data     struct {
		Items []struct {
			Code         string `json:"Code"`         // Код детали (используется для показа остатков на складах/магазинах, , в GetPartsRemainsByCode2 и других методах PartsRemains);
			Name         string `json:"Name"`         // Название детали;
			Note         string `json:"Note"`         // Описание детали;
			UniqueNumber string `json:"UniqueNumber"` // Уникальный номер;
			OmegaNumber  string `json:"OmegaNumber"`  // Код омеги (только по отдельному доступу*);
			SkubaNumber  string `json:"SkubaNumber"`  // Код скубы (только по отдельному доступу*);
			Group        string `json:"Group"`        // Группа (применяемость);

			CodeImagePart bool `json:"CodeImagePart"` // Код изображения детали (используется для получения изображения детали);

			IsPartAttendant bool `json:"IsPartAttendant"` // Если True - это сопутствующая деталь, если False - деталь, которую часто покупают со входящей деталью;

			IsSklad       bool   `json:"IsSklad"`       // Наличие на складе (True - имеется);
			IsShops       bool   `json:"IsShops"`       // Наличие в розничной сети (True - имеется);
			IsShipment    bool   `json:"IsShipment"`    // Наличие в пути (True - имеется);
			IsOutside     bool   `json:"IsOutside"`     // Наличие на внешних складах (True - имеется);
			ClientArticle string `json:"ClientArticle"` // Артикул клиента (только по отдельному доступу*);
		} `json:"Items"`
		Total int `json:"Total"`
	} `json:"Data"`
	// * Доступ выдает ответственный менеджер.
	//** Флаги IsSklad, IsShops, IsShipment, IsOutside требуют специального разрешения, т.к. их расчет сказывается на производительности сервера. Если нет разрешения, то будет установлено значение "0" независимо от наличия на складе. В таком случае для просмотра остатков необходимо воспользоваться методом GetPartsRemainsByCode или GetPartsRemainsBySkubaNumber сервиса PartsRemains.
}

func (user User) PartAttendantByCode(text string) (PartAttendantByCodeResponse, error) {
	// Запаковать в json
	bytesRepresentation, err := json.Marshal(PartAttendantByCodeRequest{Code: text, HashSession: user.HashSession})
	if err != nil {
		return PartAttendantByCodeResponse{}, err
	}

	// Выполнить запрос
	resp, err := http.Post(URL+"/v2/json/Search.svc/PartAttendantByCode", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return PartAttendantByCodeResponse{}, err
	}

	// Преобразовать данные в массив byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PartAttendantByCodeResponse{}, err
	}

	fmt.Println(string(body))

	// Распарсить данные
	var PartAttendantByCodeRes PartAttendantByCodeResponse
	responseErrorUnmarshal := json.Unmarshal(body, &PartAttendantByCodeRes)
	if responseErrorUnmarshal != nil {
		return PartAttendantByCodeResponse{}, responseErrorUnmarshal
	}

	// Проверка на отсутствие данных
	if !PartAttendantByCodeRes.Success {
		return PartAttendantByCodeResponse{}, errors.New(PartAttendantByCodeRes.Errors[0])
	}

	return PartAttendantByCodeRes, nil
}
