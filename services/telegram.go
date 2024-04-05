package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotImplementation struct {
	localBot *telegrambot.BotAPI
}

func InitTelegramBot() *BotImplementation {
	bot, err := telegrambot.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Print(err)
		return nil
	}

	bot.Debug = false
	loc := &BotImplementation{
		localBot: bot,
	}

	go loc.getTelegramUpdates()
	return loc
}

func (bot *BotImplementation) getTelegramUpdates() {
	u := telegrambot.NewUpdate(0)
	u.Timeout = 60
	// u.AllowedUpdates = []string{"chat_member"}

	updates := bot.localBot.GetUpdatesChan(u)
	for update := range updates {
		bot.processTelegramUpdates(&update)
	}
}

func (bot *BotImplementation) processTelegramUpdates(update *telegrambot.Update) {

	if isGetID(update, bot.localBot) {
		bot.sendTelegramMessage(update.Message.Chat.ID, fmt.Sprintf("ID is %d", update.Message.Chat.ID))
	}
}

// func (bot *BotImplementation) processNewMemberEvent(update *telegrambot.Update) {

// 	config := telegrambot.ChatInviteLinkConfig{
// 		ChatConfig: telegrambot.ChatConfig{
// 			ChatID:             update.Message.Chat.ID,
// 			SuperGroupUsername: "",
// 		},
// 	}

// 	link, err := bot.localBot.GetInviteLink(config)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	criteria, err := engine.FetchTelegramCriteriaByLink(link)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	criteria.GroupID = getChatID(update.Message)
// 	// criteria.BotAdded = true

// 	err = engine.SaveModel(criteria)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// }

func isNewMember(update *telegrambot.Update) bool {
	botID, _ := strconv.Atoi(os.Getenv("TELEGRAM_BOT_ID"))

	if update.ChatMember != nil {
		return update.ChatMember.NewChatMember.User.ID == int64(botID)
	}

	return false
}

func isGetID(update *telegrambot.Update, localBot *telegrambot.BotAPI) bool {
	if update.Message != nil && update.Message.Chat != nil {
		return strings.EqualFold(update.Message.Text, fmt.Sprintf("@%s getid", localBot.Self.UserName))
	}

	return false
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

func (bot *BotImplementation) sendTelegramMessage(chatID int64, message string) error {

	msg := telegrambot.NewMessage(chatID, message)
	_, err := bot.localBot.Send(msg)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (bot *BotImplementation) GetTelegramGroupUser(chatID, userID int64) (*telegrambot.ChatMember, error) {

	memberConfig := telegrambot.GetChatMemberConfig{
		ChatConfigWithUser: telegrambot.ChatConfigWithUser{
			ChatID: chatID,
			UserID: userID,
		},
	}

	member, err := bot.localBot.GetChatMember(memberConfig)
	if err != nil {
		return nil, err
	}

	return &member, nil
}

func (bot *BotImplementation) GetTelegramGroupTitle(chatID int64) (string, error) {

	chatConfig := telegrambot.ChatInfoConfig{
		ChatConfig: telegrambot.ChatConfig{
			ChatID: chatID,
		},
	}

	chat, err := bot.localBot.GetChat(chatConfig)
	if err != nil {
		return "", err
	}

	return chat.Title, nil
}

func getChatID(m *telegrambot.Message) int64 {
	return int64(m.Chat.ID)
}
