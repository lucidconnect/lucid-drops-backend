package notifier

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"inverse.so/structure"

	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NotifyTelegram(messageBody string, team structure.Team) {

	isProd, _ := IsProduction()
	if !isProd {
		log.Print(messageBody)
		return
	}

	chatID, _ := strconv.Atoi(string(team))
	err := sendTelegramMessage(int64(chatID), messageBody)
	if err != nil {
		log.Print("An error occurred while sending telegram message: ", err)
		return
	}

}

func getOpsRecipients() []string {
	list := os.Getenv("OPS_EMAIL_LIST")
	return strings.Split(list, ",")
}

func initTelegramBot() *telegrambot.BotAPI {
	bot, err := telegrambot.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Print(err)
	}

	bot.Debug = true
	return bot
}

func sendTelegramMessage(chatID int64, message string) error {

	bot := initTelegramBot()
	msg := telegrambot.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func IsProduction() (bool, error) {
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
	case "production":
		return true, nil
	case "staging":
		return false, nil
	case "development":
		return false, nil
	case "test":
		return false, nil
	default:
		return false, fmt.Errorf("unknown environment %s", appEnv)
	}
}
