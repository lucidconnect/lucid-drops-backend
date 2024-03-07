package drops

import (
	"errors"
	"time"

	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/models"
	"github.com/rs/zerolog/log"
)

func CreateDrop(input *model.DropInput, authDetails *internal.AuthDetails) (*model.Drop, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop")
	}

	if input.Name == nil || input.Image == nil || input.Thumbnail == nil || input.Description == nil {
		return nil, errors.New("pass in all Fields inorder to create a new drop")
	}

	aaSigerInfo, err := engine.GetAltSignerByCreatorID(creator.ID.String())
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop ( They lack an AA wallet )")
	}

	var contractAdddress string
	if input.ContractAddress == nil {
		// Introduce an artificial delay for before fethcing the actual contract address
		time.Sleep(time.Second * 3)

		contractAdddress, err = GetOnchainContractAddressFromDeploymentHash(input.DeploymentHash)
		if err != nil {
			log.Err(err)
		}

	} else {
		contractAdddress = *input.ContractAddress
	}

	newDrop := &models.Drop{
		CreatorID:              creator.ID,
		CreatorAddress:         aaSigerInfo.WalletAddress,
		Name:                   *input.Name,
		Image:                  *input.Image,
		Thumbnail:              *input.Thumbnail,
		Description:            *input.Description,
		BlockchainNetwork:      input.Network,
		AAWalletDeploymentHash: &input.DeploymentHash,
		AAContractAddress:      &contractAdddress,
	}

	err = engine.CreateModel(newDrop)
	if err != nil {
		return nil, errors.New("couldn't create new drop")
	}

	// create item
	dropId := newDrop.ID.String()
	itemInput := &model.ItemInput{
		Name:         input.Name,
		Image:        input.Image,
		Description:  input.Description,
		DropID:       &dropId,
		EditionLimit: input.EditionLimit,
		ClaimFee:     input.ClaimFee,
	}

	item, err := CreateItem(itemInput, authDetails)
	if err != nil {
		log.Err(err).Caller().Send()
		return newDrop.ToGraphData(nil), errors.New("failed to create item in drop")
	}

	var items []*model.Item
	items = append(items, item)
	return newDrop.ToGraphData(items), nil
}

func DeleteDrop(dropID string, authDetails *internal.AuthDetails) (*model.Drop, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop")
	}

	drop, err := engine.GetDropByID(dropID)
	if err != nil {
		return nil, errors.New("drop not found")
	}

	if creator.ID != drop.CreatorID {
		return nil, errors.New("the drop doesn't belong to this creator")
	}

	err = engine.SoftDeleteModel(drop)
	if err != nil {
		return nil, errors.New("couldn't delete drop")
	}

	return drop.ToGraphData(nil), nil
}

func UpdateDrop(dropID string, input *model.DropInput, authDetails *internal.AuthDetails) (*model.Drop, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop")
	}

	drop, err := engine.GetDropByID(dropID)
	if err != nil {
		return nil, errors.New("drop not found")
	}

	if creator.ID != drop.CreatorID {
		return nil, errors.New("the drop doesn't belong to this creator")
	}

	if input.Name != nil {
		drop.Name = *input.Name
	}

	if input.Image != nil {
		drop.Image = *input.Image
	}

	if input.Thumbnail != nil {
		drop.Image = *input.Thumbnail
	}

	if input.Description != nil {
		drop.Description = *input.Description
	}

	err = engine.SaveModel(drop)
	if err != nil {
		return nil, errors.New("couldn't create new drop")
	}

	return drop.ToGraphData(nil), nil
}

func FetchCreatorDrops(authDetails *internal.AuthDetails) ([]*model.Drop, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded")
	}

	drops, err := engine.GetCreatorDrops(creator.ID.String())
	if err != nil {
		return nil, errors.New("drops not found")
	}

	mappedDrops := make([]*model.Drop, len(drops))

	for idx, drop := range drops {
		items, _ := FetchDropItems(drop.ID.String(), false, nil)
		mappedDrops[idx] = drop.ToGraphData(items)
	}

	return mappedDrops, nil
}

func FetchDropByID(dropID string) (*model.Drop, error) {
	drop, err := engine.GetDropByID(dropID)
	if err != nil {
		return nil, errors.New("drop not found")
	}

	items, _ := FetchDropItems(dropID, false, nil)

	return drop.ToGraphData(items), nil
}

func FetchFeaturedDrops() ([]*model.Drop, error) {
	drops, err := engine.GetFeaturedDrops()
	if err != nil {
		return nil, errors.New("drops not found")
	}

	mappedDrops := make([]*model.Drop, len(drops))

	for idx, drop := range drops {
		mappedDrops[idx] = drop.ToGraphData(nil)
	}

	return mappedDrops, nil
}
