package models

import (
	"github.com/ethereum/go-ethereum/common"
	"inverse.so/graph/model"
)

type Creator struct {
	Base
	WalletAddress   common.Address
	InverseUsername *string
	SignerInfo      []SignerInfo `gorm:"foreignKey:CreatorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (c *Creator) ToGraphData() *model.CreatorDetails {
	return &model.CreatorDetails{
		Address:         c.WalletAddress.String(),
		CreatorID:       c.ID.String(),
		InverseUsername: c.InverseUsername,
	}
}
