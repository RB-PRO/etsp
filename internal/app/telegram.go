package app

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/RB-PRO/etsp/pkg/etsp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot() {
	// **********************************
	token, ErrorFile := dataFile("Token")
	if ErrorFile != nil {
		log.Fatal(ErrorFile)
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	// *************************************
	// Получение логина и пароля из файлов
	login, ErrorFile := dataFile("Login")
	if ErrorFile != nil {
		log.Fatalln(ErrorFile)
	}
	password, ErrorFile := dataFile("Password")
	if ErrorFile != nil {
		log.Fatalln(ErrorFile)
	}

	// Объявление пользователя
	user := etsp.User{
		Login:    login,
		Password: password,
	}

	// Авторизация
	_, errorAuf := user.Logon()
	if errorAuf != nil {
		log.Fatalln(errorAuf)
	}
	time.Sleep(100 * time.Microsecond)
	// *************************************

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages

			// Непосредственная обработка

			str := strings.Split(update.Message.Text, "\n")

			errNumb, errStr, errorSearch := Run(user, str)
			if len(errNumb) != 0 {
				var outStr string
				for i := 0; i < len(errNumb); i++ {
					outStr += errNumb[i] + " - " + errStr[i] + "\n"
				}
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Некоторые номера не смогли отработать из-за следующих ошибок:\n"+outStr))
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Попробуйте повторить запрос для номеров:"))
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(errNumb, "\n")))
			}
			fmt.Println(errorSearch, update.Message.Chat.UserName)
			if errorSearch != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errorSearch.Error()))
			}

			log.Println("Done", update.Message.Chat.UserName)

			file := tgbotapi.FilePath("etsp.xlsx")

			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, file))
			//bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Вот Ваш файл"))
			continue

		}

		switch update.Message.Command() {
		case "start":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот, который поможет тебе собрать информацию с сайта www.etsp.ru\nОтправь мне данные для поиска, где один запрос - одна строка. Пример:\n\n1261-2919010\n1261-2919010"))
			continue
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду, дружище...\nПопробуй /start"))
			continue
		}
	}
}
