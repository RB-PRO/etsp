package app

import (
	"io"
	"log"
	"os"

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
