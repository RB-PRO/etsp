package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/RB-PRO/etsp/pkg/etsp"
)

func RunForArray() {
	SearchArray := make([]string, 2)
	SearchArray[0] = "1261-2919010"
	SearchArray[1] = "1262-2919010"
	errorSearch := Run(SearchArray)
	if errorSearch != nil {
		log.Fatal(errorSearch)
	}
}

func Run(SearchArray []string) error {
	fileOut, errorMakeXlsx := etsp.MakeWorkBook()
	if errorMakeXlsx != nil {
		return errorMakeXlsx
	}
	etsp.WriteHead(fileOut, "main")

	// Получение логина и пароля из файлов
	login, ErrorFile := dataFile("Login")
	if ErrorFile != nil {
		return ErrorFile
	}
	password, ErrorFile := dataFile("Password")
	if ErrorFile != nil {
		return ErrorFile
	}

	// Объявление пользователя
	user := etsp.User{
		Login:    login,
		Password: password,
	}

	// Авторизация
	_, errorAuf := user.Logon()
	if errorAuf != nil {
		return errorAuf
	}
	time.Sleep(100 * time.Microsecond)

	// ************************************************
	var count int = 2

	// Проходим по исходному массиву
	for _, SearchArrayVal := range SearchArray {
		// Простой поиск
		SearchBasicRes, SearchBasicError := user.SearchBasic(SearchArrayVal)
		if SearchBasicError != nil {
			return SearchBasicError
		}
		time.Sleep(100 * time.Microsecond)

		if len(SearchBasicRes.Data.Items) != 0 {
			for indexSearchBasic, valueSearchBasic := range SearchBasicRes.Data.Items {
				fmt.Println("Code:", valueSearchBasic.Code)

				// Поиск по коду товара
				GetPartsRemainsByCodeRes, GetPartsRemainsByCodeError := user.GetPartsRemainsByCode(valueSearchBasic.Code) //SearchBasicRes.Data.Items[0].Code)
				if GetPartsRemainsByCodeError != nil {
					return GetPartsRemainsByCodeError
				}

				if len(GetPartsRemainsByCodeRes.Data.Remains) != 0 {
					for indexGetPartsRemainsByCode, valueGetPartsRemainsByCode := range GetPartsRemainsByCodeRes.Data.Remains {

						// проверка на Хабаровск
						if valueGetPartsRemainsByCode.StorageName == "Хабаровск" {

							etsp.WriteOneLine(fileOut, "main", count, SearchBasicRes, indexSearchBasic, GetPartsRemainsByCodeRes, indexGetPartsRemainsByCode)

							count++

						}
					}
				}
				//fmt.Println(GetPartsRemainsByCodeRes.Data.Remains[0].StorageName)
				time.Sleep(100 * time.Microsecond)
			}
		}
	}
	// ************************************************

	// Деавторизация
	_, errorLogout := user.Logout()
	if errorLogout != nil {
		return errorLogout
	}

	// ************************************************ EXCEL SAVE ************************************************
	fileCloseError := etsp.CloseXlsx(fileOut)
	if fileCloseError != nil {
		return fileCloseError
	}
	return nil
}

// Получение значение из файла
func dataFile(filename string) (string, error) {
	// Открыть файл
	fileToken, errorToken := os.Open(filename)
	if errorToken != nil {
		return "", errorToken
	}

	// Прочитать значение файла
	data := make([]byte, 64)
	n, err := fileToken.Read(data)
	if err == io.EOF { // если конец файла
		return "", errorToken
	}
	fileToken.Close() // Закрытие файла

	return string(data[:n]), nil
}
