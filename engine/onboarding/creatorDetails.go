package onboarding

import (
	"inverse.so/engine"
	"inverse.so/graph/model"
)

func FetchItemCreatorByCollectionId(collectionID string) (*model.CreatorDetails, error) {
	// add cache
	collection, err := engine.GetCollectionByID(collectionID)
	if err != nil {
		return nil, err
	}

	creator, err := engine.GetCreatorByID(collection.CreatorID.String())
	if err != nil {
		return nil, err
	}
	return creator.ToGraphData(), nil
}
