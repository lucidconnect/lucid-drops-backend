package claimers

import (
	"inverse.so/engine"
	"inverse.so/graph/model"
)

func FetchClaimedItems(address string) ([]*model.Item, error) {
	items, err := engine.GetClaimedItemByAddress(address)
	if err != nil {
		return nil, err
	}

	mappedItems := make([]*model.Item, len(items))

	for idx, item := range items {
		mappedItems[idx] = item.ToGraphData()
	}

	return mappedItems, nil
}
