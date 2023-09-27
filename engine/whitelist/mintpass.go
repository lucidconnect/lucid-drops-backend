package whitelist

import (
	"errors"
	"regexp"
	"time"

	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/models"
	"inverse.so/utils"
)

func IsThisAValidEthAddress(maybeAddress string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	if len(maybeAddress) != 43 {
		return false
	}

	return re.MatchString(maybeAddress)
}

func CreateMintPassForNoCriteriaItem(itemID string) (*model.ValidationRespoonse, error) {

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, err
	}

	if time.Now().After(*item.ClaimDeadline) {
		return nil, errors.New("the item is no longer available to be claimed")
	}

	if item.Criteria == nil || *item.Criteria != model.ClaimCriteriaTypeEmptyCriteria {
		return nil, errors.New("unable to generate mintpass for this item")
	}

	collection, err := engine.GetCollectionByID(item.CollectionID.String())
	if err != nil {
		return nil, errors.New("collection not found")
	}

	if collection.AAContractAddress == nil {
		return nil, errors.New("collection contract address not found")
	}

	if item.TokenID == nil {
		return nil, errors.New("The requested item is not ready to be claimed, please try again in a few minutes")
	}

	newMint := models.MintPass{
		ItemId:                    item.ID.String(),
		ItemIdOnContract:          *item.TokenID,
		CollectionContractAddress: *collection.AAContractAddress,
	}

	err = dbutils.DB.Create(&newMint).Error
	if err != nil {
		return nil, err
	}

	return &model.ValidationRespoonse{
		Valid:  true,
		PassID: utils.GetStrPtr(newMint.ID.String()),
	}, nil
}
