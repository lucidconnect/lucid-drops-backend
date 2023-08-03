package collections

import (
	"errors"

	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
)

func CreateCollection(input *model.CollectionInput, authDetails *internal.AuthDetails) (*model.Collection, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection")
	}

	if input.Name == nil || input.Image == nil || input.Thumbnail == nil || input.Description == nil {
		return nil, errors.New("pass in all Fields inorder to create a new collection")
	}

	aaSigerInfo, err := engine.GetAltSignerByAddress(creator.ID.String())
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection ( They lack an AA wallet )")
	}

	newCollection := &models.Collection{
		CreatorID:      creator.ID,
		CreatorAddress: aaSigerInfo.WalletAddress,
		Name:           *input.Name,
		Image:          *input.Image,
		Thumbnail:      *input.Thumbnail,
		Description:    *input.Description,
	}

	err = engine.CreateModel(newCollection)
	if err != nil {
		return nil, errors.New("couldn't create new collection")
	}

	return newCollection.ToGraphData(), nil
}

func UpdateCollection(collectionID string, input *model.CollectionInput, authDetails *internal.AuthDetails) (*model.Collection, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection")
	}

	collection, err := engine.GetCollectionByID(collectionID)
	if err != nil {
		return nil, errors.New("collection not found")
	}

	if creator.ID != collection.CreatorID {
		return nil, errors.New("the collection doesn't belong to this creator")
	}

	if input.Name != nil {
		collection.Name = *input.Name
	}

	if input.Image != nil {
		collection.Image = *input.Image
	}

	if input.Thumbnail != nil {
		collection.Image = *input.Thumbnail
	}

	if input.Description != nil {
		collection.Description = *input.Description
	}

	err = engine.SaveModel(collection)
	if err != nil {
		return nil, errors.New("couldn't create new collection")
	}

	return collection.ToGraphData(), nil
}

func FetchCreatorCollections(authDetails *internal.AuthDetails) ([]*model.Collection, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded")
	}

	collections, err := engine.GetCreatorCollections(creator.ID.String())
	if err != nil {
		return nil, errors.New("collections not found")
	}

	mappedCollections := make([]*model.Collection, len(collections))

	for idx, collection := range collections {
		mappedCollections[idx] = collection.ToGraphData()
	}

	return mappedCollections, nil
}

func FetchCollectionByID(collectionID string) (*model.Collection, error) {
	collection, err := engine.GetCollectionByID(collectionID)
	if err != nil {
		return nil, errors.New("collection not found")
	}

	return collection.ToGraphData(), nil
}
