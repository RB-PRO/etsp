package etsp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Структура запроса на поиск
type GetPartsRemainsByCodeRequest struct {
	Code                      string `json:"Code"`                      // Код детали
	ShowRetailRemains         bool   `json:"ShowRetailRemains"`         // Признак показа остатков розничной сети (работает только по отдельному доступу*)
	ShowOutsideRemains        bool   `json:"ShowOutsideRemains"`        //  Признак показа товаров под заказ (работает только по отдельному доступу*)
	ExcludeSubstandardRemains bool   `json:"ExcludeSubstandardRemains"` // Признак исключения из вывода остатков некондиционного товара.
	HashSession               string `json:"HashSession"`               // Хэш-ключ сессии
}

// Структура ответа на метод простого поска [GetPartsRemainsByCode]
//
// [GetPartsRemainsByCode]: https://ws.etsp.ru/Help/v2/Search/GetPartsRemainsByCode.aspx
type GetPartsRemainsByCodeResponse struct {
	Errors   []string `json:"Errors"`
	Success  bool     `json:"Success"`
	Warnings []string `json:"Warnings"`
	Data     struct {
		OutsideRemains struct { // цены на определенный товар, в зависимости от резервируемого количества
		} `json:"OutsideRemains"`
		Part struct { // ??? Неизвестный параметр
		} `json:"Part"`
		QuantityDependence []struct { // остатки на внешних складах
		} `json:"QuantityDependence"`
		Remains []struct { // остатки на складах
			ApproximateIncomeDateInFreeStatus string  `json:"ApproximateIncomeDateInFreeStatus"` // Ориентировочная дата поступления товара в свободный остаток.
			GoodsCode                         string  `json:"GoodsCode"`                         // Артикул товара
			GoodsComment                      string  `json:"GoodsComment"`                      // Описание товара
			GoodsUnitID                       int     `json:"GoodsUnitId"`                       // Идентификатор товара
			HasPriceQuantityDependence        bool    `json:"HasPriceQuantityDependence"`        // Признак наличия цены от количества (если True, сведения будут в группе QuantityDependence, отфильтровывать по комбинации GoodsUnitId+StorageId)
			IsNotRefundable                   bool    `json:"IsNotRefundable"`                   // Признак невозвратности (True - товар нельзя вернуть)
			IsSubstandard                     bool    `json:"IsSubstandard"`                     // Признак некондиции (True - некондиция)
			ManufacturerName                  string  `json:"ManufacturerName"`                  // Производитель
			Price                             float64 `json:"Price"`                             // Цена клиента
			Quantity                          string  `json:"Quantity"`                          // Кол-во товара в виде строки**
			QuantityValue                     int     `json:"QuantityValue"`                     // Кол-во товара в виде числа**
			QuantityWithRestrictions          bool    `json:"QuantityWithRestrictions"`          // Признак ограниченного доступа к просмотру количества. Если доступ ограничен, то товар больше 8 показывается как >8
			RemainsStatusID                   int     `json:"RemainsStatusId"`                   // Идентификатор статуса товара
			RemainsStatusName                 string  `json:"RemainsStatusName"`                 // Статус товара (свободный остаток, оформление, в пути и др.)
			StorageID                         int     `json:"StorageId"`                         // Идентификатор склада
			StorageName                       string  `json:"StorageName"`                       // Название склада
			StoragePosition                   int     `json:"StoragePosition"`                   // Позиция склада (для сортировки)
			Weight                            float64 `json:"Weight"`                            // Вес (кг.)

			//ManufacturerNumber                string      `json:"ManufacturerNumber"`                // Номер производителя (только по отдельному доступу*)
			//PriceCost                         interface{} `json:"PriceCost"`                         // Цена себестоимости (только по отдельному доступу*)
			//PricePurchasing                   interface{} `json:"PricePurchasing"`                   // Цена закупки (только по отдельному доступу*)
			//PriceRetail                       interface{} `json:"PriceRetail"`                       // Цена розничная (только по отдельному доступу*)
			//PriceWholesale                    interface{} `json:"PriceWholesale"`                    // Цена оптовая (только по отдельному доступу*)
			//QuantityInCart                    interface{} `json:"QuantityInCart"`                    // Поставлено в корзину
			//QuantityOrdered                   interface{} `json:"QuantityOrdered"`                   // Поставлено в резерв (еще не отгружено)
			//RegularDeliveryInfo               interface{} `json:"RegularDeliveryInfo"`               // ??? Неизвестный параметр
			//SelfDeliveryInfo                  interface{} `json:"SelfDeliveryInfo"`                  // ??? Неизвестный параметр
		} `json:"Remains"`
	} `json:"Data"`
}

func (user User) GetPartsRemainsByCode(text string) (GetPartsRemainsByCodeResponse, error) {
	// Запаковать в json
	bytesRepresentation, err := json.Marshal(GetPartsRemainsByCodeRequest{Code: text, HashSession: user.HashSession})
	if err != nil {
		return GetPartsRemainsByCodeResponse{}, err
	}

	// Выполнить запрос
	resp, err := http.Post(URL+"/v2/json/PartsRemains.svc/GetPartsRemainsByCode", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return GetPartsRemainsByCodeResponse{}, err
	}

	// Преобразовать данные в массив byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetPartsRemainsByCodeResponse{}, err
	}

	//fmt.Println(string(body))

	// Распарсить данные
	var GetPartsRemainsByCodeRes GetPartsRemainsByCodeResponse
	responseErrorUnmarshal := json.Unmarshal(body, &GetPartsRemainsByCodeRes)
	if responseErrorUnmarshal != nil {
		return GetPartsRemainsByCodeResponse{}, responseErrorUnmarshal
	}

	// Проверка на отсутствие данных
	if !GetPartsRemainsByCodeRes.Success {
		return GetPartsRemainsByCodeResponse{}, errors.New(GetPartsRemainsByCodeRes.Errors[0])
	}

	return GetPartsRemainsByCodeRes, nil
}
