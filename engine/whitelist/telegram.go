package whitelist

import (
	"errors"
	"strconv"

	"inverse.so/dbutils"
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
		return nil, errors.New("creator has not been onboarded to create a new collection")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if item.Criteria != nil {
		//Delete Existing criteria
		err := engine.DeleteCriteriaIfExists(item)
		if err != nil {
			return nil, err
		}
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

	groupTitle, err := InverseBot.GetTelegramGroupTitle(int64(groupIDToInt))
	if err == nil {
		criteria.GroupTitle = groupTitle
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

func ValidateTelegramClaimCriteria(itemID, authID string) (*model.ValidationRespoonse, error) {

	resp := &model.ValidationRespoonse{
		Valid: false,
	}

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return resp, errors.New("item not found")
	}

	if item.TelegramCriteria == nil {
		return resp, errors.New("item does not have a telegram criteria")
	}

	IdToInt, _ := strconv.Atoi(authID)
	member, err := InverseBot.GetTelegramGroupUser(item.TelegramCriteria.GroupID, int64(IdToInt))
	if err != nil {
		return resp, errors.New("telegram account not authorized by group admin")
	}

	if member.User.IsBot {
		return resp, errors.New("telegram account cannot be a bot")
	}

	if member.Status == "member" || member.Status == "creator" || member.Status == "administrator" {

		PassID, err := createMintPassForTelegramMint(item)
		if err != nil {
			return resp, errors.New("error creating mint pass")
		}

		resp.Valid = true
		resp.PassID = PassID
		return resp, nil
	}

	return nil, errors.New("telegram account not authorized by group admin")
}

func createMintPassForTelegramMint(item *models.Item) (*string, error) {
	collection, err := engine.GetCollectionByID(item.CollectionID.String())
	if err != nil {
		return nil, errors.New("collection not found")
	}

	if collection.ContractAddress == nil {
		return nil, errors.New("collection contract address not found")
	}

	if item.TokenID == nil {
		return nil, errors.New("The requested item is not ready to be claimed, please try agan n a few minutes")
	}

	newMint := models.MintPass{
		ItemId:                    item.ID.String(),
		ItemIdOnContract:          *item.TokenID,
		CollectionContractAddress: *collection.ContractAddress,
	}

	err = dbutils.DB.Create(&newMint).Error
	if err != nil {
		return nil, err
	}

	passId := newMint.ID.String()
	return &passId, nil
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
