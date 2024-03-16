package models

import (
	"github.com/lucidconnect/inverse/graph/model"
	uuid "github.com/satori/go.uuid"
)

type Drop struct {
	Base
	CreatorID              uuid.UUID
	CreatorAddress         string
	Name                   string
	Image                  string `json:"image"`
	Thumbnail              string `json:"thumbnail"`
	Description            string `json:"description"`
	AAContractAddress      *string
	TransactionHash        *string
	AAWalletDeploymentHash *string
	BlockchainNetwork      *model.BlockchainNetwork
	Featured               bool `gorm:"default:false"`
	MintUrl                string
	MintPrice              *float64
	GasIsCreatorSponsored  bool
	Criteria               *model.ClaimCriteriaType
	FarcasterCriteria      *FarcasterCriteria `gorm:"foreignKey:DropId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserLimit              *int `gorm:"default:null"`
	EditionLimit           *int `gorm:"default:null"`
}

type DeplyomenResponse struct {
	Type          int    `json:"type"`
	ChainID       int    `json:"chainId"`
	Nonce         int    `json:"nonce"`
	To            string `json:"to"`
	Data          string `json:"data"`
	Hash          string `json:"hash"`
	From          string `json:"from"`
	Confirmations int    `json:"confirmations"`
}

func (c *Drop) ToGraphData(items []*model.Item) *model.Drop {
	mappedDrop := &model.Drop{
		ID:                    c.ID.String(),
		CreatorID:             c.CreatorID.String(),
		CreatedAt:             c.CreatedAt,
		Name:                  c.Name,
		Description:           c.Description,
		Image:                 c.Image,
		Thumbnail:             c.Thumbnail,
		ContractAddress:       c.AAContractAddress,
		Network:               c.BlockchainNetwork,
		MintURL:               c.MintUrl,
		GasIsCreatorSponsored: c.GasIsCreatorSponsored,
		ClaimCriteria:         c.Criteria,
	}

	if c.AAContractAddress != nil {
		mappedDrop.ContractAddress = c.AAContractAddress
	}

	if c.MintPrice != nil {
		mappedDrop.MintPrice = c.MintPrice
	}

	if c.FarcasterCriteria != nil {
		mappedDrop.FarcasterClaimCriteriaInteractions = InteractionsToArr(c.FarcasterCriteria.Interactions)
		mappedDrop.CastURL = &c.FarcasterCriteria.CastUrl
		mappedDrop.FarcasterProfileID = &c.FarcasterCriteria.FarcasterProfileID
		mappedDrop.FarcasterChannelID = &c.FarcasterCriteria.ChannelID
	}

	if items != nil {
		for _, item := range items {
			item.DropAddress = *c.AAContractAddress
		}
		mappedDrop.Items = items
	}

	return mappedDrop
}
