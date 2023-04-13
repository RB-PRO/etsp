package app

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/RB-PRO/etsp/pkg/etsp"
	"github.com/cheggaaa/pb"
)

func Run(user etsp.User, SearchArray []string) ([]string, []string, error) {
	fileOut, errorMakeXlsx := etsp.MakeWorkBook()
	if errorMakeXlsx != nil {
		return nil, nil, errorMakeXlsx
	}
	etsp.WriteHeadCustom(fileOut, "main")

	//manuf, errorManuf := user.ManufacturerList()
	//if errorManuf != nil {
	//	return errorManuf
	//}
	//fmt.Println(manuf)
	//fmt.Printf("%+#v", manuf)
	// ************************************************
	var count int = 2
	var errorsNumbers []string
	var errorsStrs []string

	// Проходим по исходному массиву
	bar := pb.StartNew(len(SearchArray))
	defer bar.Finish()
	for _, SearchArrayVal := range SearchArray {
		bar.Prefix(SearchArrayVal)
		bar.Increment()
		// Простой поиск
		SearchBasicRes, SearchBasicError := user.SearchBasic(SearchArrayVal)
		if SearchBasicError != nil {
			//log.Println(SearchBasicError)
			errorsNumbers = append(errorsNumbers, SearchArrayVal)
			errorsStrs = append(errorsStrs, SearchBasicError.Error())
			log.Println("Error 1:", SearchArrayVal, SearchBasicError)
			continue
		}
		time.Sleep(300 * time.Microsecond)

		if len(SearchBasicRes.Data.Items) != 0 {
			for indexSearchBasic, valueSearchBasic := range SearchBasicRes.Data.Items {
				//log.Println("Код:", valueSearchBasic.Code)

				// Поиск по коду товара
				GetPartsRemainsByCodeRes, GetPartsRemainsByCodeError := user.GetPartsRemainsByCode(valueSearchBasic.Code) //SearchBasicRes.Data.Items[0].Code)
				if GetPartsRemainsByCodeError != nil {
					log.Println("Error 2:", SearchArrayVal, GetPartsRemainsByCodeError)
				}

				if len(GetPartsRemainsByCodeRes.Data.Remains) != 0 {
					for indexGetPartsRemainsByCode, valueGetPartsRemainsByCode := range GetPartsRemainsByCodeRes.Data.Remains {
						if valueGetPartsRemainsByCode.StorageName == "Хабаровск" { // проверка на Хабаровск
							etsp.WriteOneLineCustom(fileOut, "main", count, SearchBasicRes, indexSearchBasic, GetPartsRemainsByCodeRes, indexGetPartsRemainsByCode)
							count++
						}
					}
				}
				time.Sleep(300 * time.Microsecond)
			}
		}
	}

	// ************************************************

	// Деавторизация
	_, errorLogout := user.Logout()
	if errorLogout != nil {
		return nil, nil, errorLogout
	}

	// ************************************************ EXCEL SAVE ************************************************
	fileCloseError := etsp.CloseXlsx(fileOut)
	//etsp.Filter(fileOut, "main")
	if fileCloseError != nil {
		return nil, nil, fileCloseError
	}
	return errorsNumbers, errorsStrs, nil
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
