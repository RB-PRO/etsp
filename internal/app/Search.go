package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/RB-PRO/etsp/pkg/etsp"
)

func Run() {
	// Получение логина и пароля из файлов
	login, ErrorFile := dataFile("Login")
	if ErrorFile != nil {
		log.Fatal(ErrorFile)
	}
	password, ErrorFile := dataFile("Password")
	if ErrorFile != nil {
		log.Fatal(ErrorFile)
	}

	// Объявление пользователя
	user := etsp.User{
		Login:    login,
		Password: password,
	}

	// Авторизация
	_, errorAuf := user.Logon()
	if errorAuf != nil {
		log.Fatal(errorAuf)
	}
	time.Sleep(100 * time.Microsecond)

	// ************************************************

	// Простой поиск
	SearchBasicRes, SearchBasicError := user.SearchBasic("1261-2919010")
	if SearchBasicError != nil {
		log.Fatal(SearchBasicError)
	}
	time.Sleep(100 * time.Microsecond)

	fmt.Println("code:", SearchBasicRes.Data.Items[0].Code)

	// Поиск по коду товара
	PartAttendantByCodeRes, PartAttendantByCodeError := user.PartAttendantByCode(SearchBasicRes.Data.Items[0].Code)
	if PartAttendantByCodeError != nil {
		log.Fatal(PartAttendantByCodeError)
	}
	fmt.Println(PartAttendantByCodeRes)

	// ************************************************

	// Деавторизация
	_, errorLogout := user.Logout()
	if errorLogout != nil {
		log.Fatal(errorLogout)
	}
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
