package items

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/jobs"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/utils"
	"github.com/rs/zerolog/log"
)

func TempCreateItem(input *model.ItemInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new collection")
	}

	if input.Name == nil || input.Image == nil || input.CollectionID == nil || input.Description == nil {
		return nil, errors.New("pass in all Fields inorder to create a new item")
	}

	collection, err := engine.GetCollectionByID(*input.CollectionID)
	if err != nil {
		return nil, errors.New("collection not found")
	}

	if collection.CreatorID != creator.ID {
		return nil, errors.New("collection doesn't belong to the creator if the item")
	}

	newItem := &models.Item{
		Name:         *input.Name,
		Image:        *input.Image,
		Description:  *input.Description,
		CollectionID: collection.ID,
		UserLimit:    input.UserLimit,
		EditionLimit: input.EditionLimit,
	}

	if input.ClaimFee != nil {
		newItem.ClaimFee = *input.ClaimFee
	}

	err = engine.CreateModel(newItem)
	if err != nil {
		return nil, errors.New("couldn't create new collection")
	}

	return newItem.ToGraphData(), nil
}

func CreateItem(input *model.ItemInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new collection")
	}

	if input.Name == nil || input.Image == nil || input.CollectionID == nil || input.Description == nil {
		return nil, errors.New("pass in all Fields inorder to create a new item")
	}

	collection, err := engine.GetCollectionByID(*input.CollectionID)
	if err != nil {
		return nil, errors.New("collection not found")
	}

	if collection.CreatorID != creator.ID {
		return nil, errors.New("collection doesn't belong to the creator if the item")
	}

	newItem := &models.Item{
		Name:         *input.Name,
		Image:        *input.Image,
		Description:  *input.Description,
		CollectionID: collection.ID,
		UserLimit:    input.UserLimit,
		EditionLimit: input.EditionLimit,
	}

	if input.ClaimFee != nil {
		newItem.ClaimFee = *input.ClaimFee
	}

	err = engine.CreateModel(newItem)
	if err != nil {
		return nil, errors.New("couldn't create new collection")
	}

	// This was removed because all deployments are now triggered by the FE
	go func() {
		inverseAAServerURL := utils.UseEnvOrDefault("AA_SERVER", "https://inverse-aa.onrender.com")
		inverseAPIBaseURL := utils.UseEnvOrDefault("API_BASE_URL", "https://inverse-backend.onrender.com")
		client := &http.Client{}
		if collection.AAContractAddress == nil {
			log.Info().Msg("🪼TODO ADD SUPPORT FOR QUEING")
			return
		}
		itemData, err := json.Marshal(map[string]interface{}{
			"image":           fmt.Sprintf("%s/metadata/%s/%s", inverseAPIBaseURL, *collection.AAContractAddress, newItem.ID.String()),
			"contractAddress": *collection.AAContractAddress,
			"Network":         collection.BlockchainNetwork,
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := http.NewRequest(http.MethodPost, inverseAAServerURL+"/additem", bytes.NewBuffer(itemData))
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		var isBase bool
		if collection.BlockchainNetwork != nil {
			isBase = *collection.BlockchainNetwork == model.BlockchainNetworkBase
		}

		if res.StatusCode == http.StatusOK {
			go func() {
				tokenID, err := jobs.FetchTokenUri(*collection.AAContractAddress, newItem.ID.String(), isBase)
				if err != nil {
					return
				}

				if tokenID == nil {
					log.Info().Msgf("🚨 Token ID not found for Item %s", newItem.ID)
					return
				}

				tokenIDint64 := int64(*tokenID)
				newItem.TokenID = &tokenIDint64
				err = engine.SaveModel(&newItem)
				if err != nil {
					log.Error().Msg(err.Error())
				}
			}()
		}
		defer res.Body.Close()
	}()

	return newItem.ToGraphData(), nil
}

func DeleteItem(itemID string, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new item")
	}

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("collection not found")
	}

	collection, err := engine.GetCollectionByID(item.CollectionID.String())
	if err != nil {
		return nil, errors.New("collection not found")
	}

	if creator.ID != collection.CreatorID {
		return nil, errors.New("the collection doesn't belong to this creator")
	}

	err = engine.SoftDeleteModel(item)
	if err != nil {
		return nil, errors.New("couldn't delete the item")
	}

	return item.ToGraphData(), nil
}

func UpdateItem(itemID string, input *model.ItemInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new item")
	}

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("collection not found")
	}

	if input.Name != nil {
		item.Name = *input.Name
	}

	if input.Image != nil {
		item.Image = *input.Image
	}

	if input.Description != nil {
		item.Description = *input.Description
	}

	if input.CollectionID != nil {
		collection, err := engine.GetCollectionByID(*input.CollectionID)
		if err != nil {
			return nil, errors.New("collection not found")
		}

		if creator.ID != collection.CreatorID {
			return nil, errors.New("the collection doesn't belong to this creator")
		}

		item.CollectionID = collection.ID
	}

	err = engine.SaveModel(item)
	if err != nil {
		return nil, errors.New("couldn't create new item")
	}

	return item.ToGraphData(), nil
}

func FetchAuthotizedSubdomainsForItem(itemID string) ([]string, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("items not found")
	}

	if !item.ShowEmailDomainHints || item.Criteria == nil || *item.Criteria != model.ClaimCriteriaTypeEmailDomain {
		return nil, nil
	}

	subdomains, err := engine.GetAuthorizedSubdomainsForItem(itemID)
	if err != nil {
		return nil, errors.New("authorized subdomains not found")
	}

	mappedDomains := make([]string, len(subdomains))
	for idx, subDomainEntry := range subdomains {
		mappedDomains[idx] = subDomainEntry.BaseDomain
	}

	return mappedDomains, nil
}

func FetchCollectionItems(collectionID string, includeDelelted bool, authDetails *internal.AuthDetails) ([]*model.Item, error) {
	// All Collection data will be public for now

	// creator, err := engine.GetCreatorByAddress(authDetails.Address)
	// if err != nil {
	// 	return nil, errors.New("creator has not been onboarded")
	// }

	var err error
	var items []*models.Item

	if includeDelelted {
		items, err = engine.GetCollectionItemsIncludeDeleted(collectionID)
		if err != nil {
			return nil, errors.New("items not found")
		}

	} else {
		items, err = engine.GetCollectionItems(collectionID)
		if err != nil {
			return nil, errors.New("items not found")
		}
	}

	mappedItems := make([]*model.Item, len(items))
	for idx, item := range items {
		mintPasses, _ := FetchMintPassesForItems(item.ID.String())
		item.MintPasses = mintPasses
		mappedItems[idx] = item.ToGraphData()
	}

	return mappedItems, nil
}

func FetchItemByID(itemID string) (*model.Item, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	mintPasses, _ := FetchMintPassesForItems(item.ID.String())
	item.MintPasses = mintPasses

	return item.ToGraphData(), nil
}

func FetchQuestionsByItemID(itemID string) ([]*model.QuestionnaireType, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	questions, err := engine.GetItemQuestionsByItem(item)
	if err != nil {
		return nil, errors.New("questions not found")
	}

	return questions, nil
}

func FetchFeaturedItems() ([]*model.Item, error) {
	items, err := engine.GetFeaturedItems()
	if err != nil {
		return nil, errors.New("items not found")
	}

	mappedItems := make([]*model.Item, len(items))

	for idx, item := range items {
		mappedItems[idx] = item.ToGraphData()
	}

	return mappedItems, nil
}

func SetItemClaimDeadline(itemID string, deadline string) (*model.Item, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	dateForrmatted, err := time.Parse(time.RFC3339Nano, deadline)
	if err != nil {
		return nil, err
	}

	item.ClaimDeadline = &dateForrmatted
	err = engine.SaveModel(item)
	if err != nil {
		return nil, errors.New("couldn't save item")
	}

	return item.ToGraphData(), nil
}

func FetchMintPassesForItems(itemID string) ([]models.MintPass, error) {

	var mintPasses []models.MintPass
	err := dbutils.DB.Model(&models.MintPass{}).Where("item_id = ?", itemID).Find(&mintPasses).Error
	if err != nil {
		return nil, err
	}

	return mintPasses, nil
}
