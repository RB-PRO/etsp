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
type searchBasicRequest struct {
	Text              string `json:"Text"`              // Строка поиска. Можно делать сложный поиск (например: вал колен)
	WithAnalogsSearch bool   `json:"WithAnalogsSearch"` // Флаг, устанавливающий возможность использования поиска с предложением аналогов (поиск аналогов производится только в том случае, если ничего не найдено по прямому запросу)
	HashSession       string `json:"HashSession"`       // Хэш-ключ сессии
}

// Структура ответа на метод простого поска [SearchBasic]
//
// [SearchBasic]: https://ws.etsp.ru/Help/v2/Search/SearchBasic.aspx
type SearchBasicResponse struct {
	Errors   []string `json:"Errors"`
	Success  bool     `json:"Success"`
	Warnings []string `json:"Warnings"`
	Data     struct {
		Items []struct {
			Code             string `json:"Code"`             // Код детали (используется для показа остатков на складах/магазинах, , в GetPartsRemainsByCode2 и других методах PartsRemains);
			Name             string `json:"Name"`             // Название детали;
			Note             string `json:"Note"`             // Описание детали;
			UniqueNumber     string `json:"UniqueNumber"`     // Уникальный номер;
			OmegaNumber      string `json:"OmegaNumber"`      // Код омеги (только по отдельному доступу*);
			SkubaNumber      string `json:"SkubaNumber"`      // Код скубы (только по отдельному доступу*);
			Group            string `json:"Group"`            // Группа (применяемость);
			Subgroup         string `json:"Subgroup"`         // Подгруппа (код подгруппы);
			CodeImagePart    string `json:"CodeImagePart"`    // Код изображения детали (используется для получения изображения детали);
			HasPartAttendant bool   `json:"HasPartAttendant"` //  Наличие сопутствующих деталией (признак используется для вызова списка сопустствующих деталей);
			IsSklad          bool   `json:"IsSklad"`          // Наличие на складе (True - имеется);
			IsShops          bool   `json:"IsShops"`          // Наличие в розничной сети (True - имеется);
			IsShipment       bool   `json:"IsShipment"`       // Наличие в пути (True - имеется);
			IsOutside        bool   `json:"IsOutside"`        // Наличие на внешних складах (True - имеется);
			ClientArticle    string `json:"ClientArticle"`    // Артикул клиента (только по отдельному доступу*);
			IsAnalog         bool   `json:"IsAnalog"`         // Признак аналога (True - аналог).
		} `json:"Items"`
		Total int `json:"Total"`
	} `json:"Data"`
	// * Доступ выдает ответственный менеджер.
	//** Флаги IsSklad, IsShops, IsShipment, IsOutside требуют специального разрешения, т.к. их расчет сказывается на производительности сервера. Если нет разрешения, то будет установлено значение "0" независимо от наличия на складе. В таком случае для просмотра остатков необходимо воспользоваться методом GetPartsRemainsByCode или GetPartsRemainsBySkubaNumber сервиса PartsRemains.
}

func (user User) SearchBasic(text string) (SearchBasicResponse, error) {
	// Запаковать в json
	bytesRepresentation, err := json.Marshal(searchBasicRequest{Text: text, WithAnalogsSearch: false, HashSession: user.HashSession})
	if err != nil {
		return SearchBasicResponse{}, err
	}

	// Выполнить запрос
	resp, err := http.Post(URL+"/v2/json/Search.svc/SearchBasic", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return SearchBasicResponse{}, err
	}

	// Преобразовать данные в массив byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchBasicResponse{}, err
	}

	fmt.Println(string(body))

	// Распарсить данные
	var SearchBasicRes SearchBasicResponse
	responseErrorUnmarshal := json.Unmarshal(body, &SearchBasicRes)
	if responseErrorUnmarshal != nil {
		return SearchBasicResponse{}, responseErrorUnmarshal
	}

	// Проверка на отсутствие данных
	if !SearchBasicRes.Success {
		return SearchBasicResponse{}, errors.New(SearchBasicRes.Errors[0])
	}

	return SearchBasicRes, nil
}
