package models

import "inverse.so/graph/model"

type Creator struct {
	Base
	WalletAddress   string
	InverseUsername *string
}

func (c *Creator) ToGraphData() *model.CreatorDetails {
	return &model.CreatorDetails{
		Address:         c.WalletAddress,
		CreatorID:       c.ID.String(),
		InverseUsername: c.InverseUsername,
	}
}
