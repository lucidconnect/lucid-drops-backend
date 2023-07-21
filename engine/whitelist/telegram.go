package whitelist

import (
	"errors"
	"strconv"

	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
	"inverse.so/services"
)

var InverseBot *services.BotImplementation

func CreateTelegramCriteria(input model.NewTelegramCriteriaInput, authDetails *internal.AuthDetails) (*model.Item, error) {

	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if input.GroupID[0] != '-' {
		input.GroupID = "-" + input.GroupID
	}

	groupIDToInt, _ := strconv.Atoi(input.GroupID)
	criteria := &models.TelegramCriteria{
		ItemID:    item.ID.String(),
		CreatorID: creator.ID.String(),
		GroupID:   int64(groupIDToInt),
	}

	criteriaUpdateErr := engine.SaveModel(criteria)
	if criteriaUpdateErr != nil {
		return nil, criteriaUpdateErr
	}

	telegramCriteria := model.ClaimCriteriaTypeTelegram
	item.Criteria = &telegramCriteria
	err = engine.SaveModel(item)
	if err != nil {
		return nil, err
	}

	return item.ToGraphData(), nil
}

func ValidateTelegramClaimCriteria(itemID, authID string) (bool, error) {

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return false, errors.New("item not found")
	}

	if item.TelegramCriteria == nil {
		return false, errors.New("item does not have a telegram criteria")
	}

	auth, err := engine.FetchTelegramAuthByID(authID)
	if err != nil {
		return false, errors.New("telegram account not authorized")
	}

	IdToInt, _ := strconv.Atoi(auth.UserID)
	member, err := InverseBot.GetTelegramGroupUser(item.TelegramCriteria.GroupID, int64(IdToInt))
	if err != nil {
		return false, errors.New("telegram account not authorized")
	}

	if member.User.IsBot {
		return false, errors.New("telegram account cannot be a bot")
	}

	return member.IsMember, nil
}

func ProcessTelegramCallBack(id, username, hash, photoURL string) (*string, error) {

	telegramInfo := &models.TelegramAuthDetails{
		UserID:   id,
		Username: username,
		Hash:     hash,
		PhotoURL: photoURL,
	}

	telegramInfoUpdateErr := engine.SaveModel(telegramInfo)
	if telegramInfoUpdateErr != nil {
		return nil, telegramInfoUpdateErr
	}

	ID := telegramInfo.ID.String()
	return &ID, nil
}
