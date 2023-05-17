package collections

import (
	"errors"

	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
)

func CreateCollection(input *model.NewCollection, authDetails *internal.AuthDetails) (*model.Collection, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator is has not been onboarded to create a new collection")
	}

	newCollection := &models.Collection{
		CreatorID:  creator.ID,
		Name:       input.Name,
		ContentURI: input.ContentURI,
	}

	err = engine.CreateCollection(newCollection)
	if err != nil {
		return nil, errors.New("couldn't create new collection")
	}

	return newCollection.ToGraphData(), nil
}
