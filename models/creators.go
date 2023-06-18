package models

import "inverse.so/graph/model"

type Creator struct {
	Base
	WalletAddress   string
	InverseUsername *string
	SignerInfo      []SignerInfo `gorm:"foreignKey:CreatorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (c *Creator) ToGraphData() *model.CreatorDetails {
	return &model.CreatorDetails{
		Address:         c.WalletAddress,
		CreatorID:       c.ID.String(),
		InverseUsername: c.InverseUsername,
	}
}
