package whitelist

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm/clause"
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
	"inverse.so/utils"
)

func CreateWalletAddressWhitelistForItem(input *model.NewWalletAddressWhitelistInput, authDetails *internal.AuthDetails) (*model.Item, error) {
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

	if len(input.AuthorizedWalletAddresses) == 0 {
		return nil, errors.New("please passing in a wallet address list")
	}

	dbWallets := make([]*models.WalletAddressClaim, len(input.AuthorizedWalletAddresses))
	for idx, address := range input.AuthorizedWalletAddresses {

		var ENS *string
		if strings.Contains(address, ".eth") {
			origialEns := address
			resolvedAddress, err := utils.ResolveENSName(address)
			if err != nil {
				continue
			}

			ENS = &origialEns
			address = *resolvedAddress
		}

		dbWallets[idx] = &models.WalletAddressClaim{
			CreatorID:     creator.ID,
			ItemID:        item.ID,
			ENS:           ENS,
			WalletAddress: address,
		}
	}

	insertionErr := dbutils.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dbWallets, 100).Error
	if insertionErr != nil {
		return nil, insertionErr
	}

	for _, dbWallet := range dbWallets {

		insertionErr := dbutils.DB.Save(dbWallet).Error
		if insertionErr != nil {
			continue
		}
	}

	addressCriteria := model.ClaimCriteriaTypeWalletAddress
	item.Criteria = &addressCriteria
	itemUpdateErr := engine.SaveModel(item)
	if itemUpdateErr != nil {
		return nil, itemUpdateErr
	}

	return item.ToGraphData(), nil
}

func ValidateAddressCriteria(itemID, walletAddress string, authDetails *internal.AuthDetails) (*model.ValidationRespoonse, error) {
	resp := &model.ValidationRespoonse{
		Valid: false,
	}

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if item.ClaimDeadline != nil {
		if time.Now().After(*item.ClaimDeadline) {
			return nil, errors.New("the item is no longer available to be claimed")
		}
	}

	if item.Criteria == nil {
		return nil, errors.New("item does not have any criteria")
	}

	if *item.Criteria != model.ClaimCriteriaTypeWalletAddress {
		return nil, errors.New("item does not have wallet address criteria")
	}

	claimVal := &models.WalletAddressClaim{}
	err = dbutils.DB.Where("item_id = ? AND wallet_address = ?", item.ID, walletAddress).First(&claimVal).Error
	if err != nil {
		return resp, errors.New("this wallet address is not allow-listed for this item")
	}

	if claimVal.SentOutAt != nil {
		return resp, errors.New("wallet address has claimed the item already")
	}

	passResp, err := CreateMintPassForValidatedCriteriaItem(item.ID.String())
	if err != nil {
		return passResp, err
	}

	return passResp, nil
}

func resolveENSToAddress(AddressList []string) ([]string, error) {

	for idx, address := range AddressList {
		if strings.Contains(address, ".eth") {
			resolvedAddress, err := utils.ResolveENSName(address)
			if err != nil {
				continue
			}

			AddressList[idx] = *resolvedAddress
		}
	}

	return AddressList, nil
}
