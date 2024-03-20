package models

import (
	"time"

	"github.com/lucidconnect/inverse/graph/model"
)

type MintPass struct {
	Base
	ItemId        string
	DropID        string
	MinterAddress string
	TokenID       string
	UsedAt        *time.Time `gorm:"default:null"`

	// ItemIdOnContract    int64
	DropContractAddress string
	BlockchainNetwork   *model.BlockchainNetwork `gorm:"default:base"`
}

func (m *MintPass) ToGraphData() *model.ClaimDetails {

	//TODO: find a way arount import cycle error
	// var username *string = nil
	// creator, err := graph.FetchCreatorByAddress(common.HexToAddress(m.MinterAddress))
	// if err == nil && creator.InverseUsername != nil {
	// 	username = creator.InverseUsername
	// }
	return &model.ClaimDetails{
		ClaimerAddress: &m.MinterAddress,
		ClaimTime:      &m.CreatedAt,
		// ClaimerUsername: username,
	}
}
