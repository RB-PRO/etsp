package app

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot() {
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

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages

			// Непосредственная обработка

			str := strings.Split(update.Message.Text, "\n")

			errorSearch := Run(str)
			if errorSearch != nil {
				//log.Fatal(errorSearch)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка. Повторите запрос."))
			} else {

				log.Println("Done", update.Message.Chat.UserName)

				file := tgbotapi.FilePath("etsp.xlsx")

				bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, file))
				//bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Вот Ваш файл"))
			}
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
