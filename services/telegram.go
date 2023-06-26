package services

import (
	"log"
	"os"
	"strconv"
	"strings"

	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"inverse.so/engine"
)

func initTelegramBot() *telegrambot.BotAPI {
	bot, err := telegrambot.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Print(err)
	}

	bot.Debug = true
	return bot
}

func getTelegramUpdates(bot *telegrambot.BotAPI) {
	u := telegrambot.NewUpdate(0)
	u.Timeout = 60
	u.AllowedUpdates = []string{"chat_member"}

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		go processTelegramUpdates(bot, &update)
	}
}

func processTelegramUpdates(bot *telegrambot.BotAPI, update *telegrambot.Update) {

	var updateBool bool
	switch updateBool {
	case isNewMember(update):
		processNewMemberEvent(bot, update)
	}
}

func processNewMemberEvent(bot *telegrambot.BotAPI, update *telegrambot.Update) {

	config := telegrambot.ChatInviteLinkConfig{
		ChatConfig: telegrambot.ChatConfig{
			ChatID:             update.Message.Chat.ID,
			SuperGroupUsername: "",
		},
	}

	link, err := bot.GetInviteLink(config)
	if err != nil {
		log.Print(err)
		return
	}

	criteria, err := engine.FetchTelegramCriteriaByLink(link)
	if err != nil {
		log.Print(err)
		return
	}

	criteria.ChannelID = getChatID(update.Message)
	criteria.BotAdded = true

	err = engine.SaveModel(criteria)
	if err != nil {
		log.Print(err)
		return
	}
}

func isNewMember(update *telegrambot.Update) bool {
	botID, _ := strconv.Atoi(os.Getenv("TELEGRAM_BOT_ID"))
	return update.ChatMember.NewChatMember.User.ID == int64(botID)
}

func isValidVerificationMessage(update *telegrambot.Update) bool {

	if update.Message == nil {
		return false
	}

	if strings.Contains(update.Message.Text, "@inverseverifybot verify") {
		return true
	}

	return false
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

func getChatID(m *telegrambot.Message) int64 {
	return int64(m.Chat.ID)
}
