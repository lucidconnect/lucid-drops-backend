package models

import "time"

type MintPass struct {
	Base
	ItemId        string
	MinterAddress string
	UsedAt        *time.Time `gorm:"default:null"`

	ItemIdOnContract          int64
	CollectionContractAddress string
}
