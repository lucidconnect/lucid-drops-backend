package models

import (
	"time"

	"inverse.so/graph/model"
)

type MintPass struct {
	Base
	ItemId        string
	MinterAddress string
	TokenID       string
	UsedAt        *time.Time `gorm:"default:null"`

	ItemIdOnContract          int64
	CollectionContractAddress string
}

func (m *MintPass) ToGraphData() *model.ClaimDetails {
	return &model.ClaimDetails{
		ClaimerAddress: &m.MinterAddress,
		ClaimTime:      &m.CreatedAt,
	}
}
