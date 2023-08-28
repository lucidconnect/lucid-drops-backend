package collections

import (
	"errors"

	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
)

func StoreHashForDeployment(authDetails *internal.AuthDetails, input *model.DeploymentInfo) (*bool, error) {
	collection, err := engine.GetCollectionByID(input.CollectionID)
	if err != nil {
		return nil, errors.New("collection not found")
	}

	collection.AAWalletDeploymentHash = &input.DeploymentHash
	collection.AAContractAddress = input.ContractAddress

	err = engine.SaveModel(collection)
	if err != nil {
		return nil, err
	}

	t := true
	return &t, nil
}
