package items

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
	"inverse.so/utils"
)

func CreateItem(input *model.ItemInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection")
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
	}

	err = engine.CreateModel(newItem)
	if err != nil {
		return nil, errors.New("couldn't create new collection")
	}

	go func() {
		inverseAAServerURL := utils.UseEnvOrDefault("INVERSE_AA_SERVER", "https://inverse-aa.onrender.com")
		inverseAPIBaseURL := utils.UseEnvOrDefault("INVERSE_API_BASEURL", "https://inverse-backend.onrender.com")

		client := &http.Client{}

		if collection.ContractAddress == nil {
			log.Println("🪼TODO ADD SUPPORT FOR QUEING")
			return
		}

		itemData, err := json.Marshal(map[string]string{
			"image":           fmt.Sprintf("%s/metadata/%s/%s", inverseAPIBaseURL, *collection.ContractAddress, newItem.ID.String()),
			"contractAddress": *collection.ContractAddress,
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

		defer res.Body.Close()
	}()

	return newItem.ToGraphData(), nil
}

func UpdateItem(itemID string, input *model.ItemInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new item")
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

func FetchCollectionItems(collectionID string, authDetails *internal.AuthDetails) ([]*model.Item, error) {
	// All Collection data will be public for now
	// creator, err := engine.GetCreatorByAddress(authDetails.Address)
	// if err != nil {
	// 	return nil, errors.New("creator has not been onboarded")
	// }

	items, err := engine.GetCollectionItems(collectionID)
	if err != nil {
		return nil, errors.New("items not found")
	}

	mappedItems := make([]*model.Item, len(items))

	for idx, item := range items {
		mappedItems[idx] = item.ToGraphData()
	}

	return mappedItems, nil
}

func FetchItemByID(itemID string) (*model.Item, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

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