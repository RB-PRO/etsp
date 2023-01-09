package etsp

import (
	"encoding/json"
	"errors"
)

const URL string = "https://ws.etsp.ru"

type User struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

type stringOrNull string

// Кастомное декодирование JSON для ключа SearchID
func (str *stringOrNull) UnmarshalJSON(data []byte) error {
	//fmt.Println(">" + string(data) + "<")
	if string(data) == "null" {
		//newData := make([]byte, 1)
		//newData[0] = 34
		//newData = append(newData, data...)
		//newData = append(newData, 34)

		err := json.Unmarshal([]byte{34, 110, 117, 108, 108, 34}, &str)
		if err != nil {
			return errors.New("CustomString: UnmarshalJSON: enable null: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &str)
		if err != nil {
			return errors.New("CustomString: UnmarshalJSON: " + err.Error())
		}

	}
	return nil
}

/*
// Кастомное кодирование JSON для ключа SearchID
func (str stringOrNull) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(str)
	return json, err
}
*/
