package models

import (
	"strings"

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
	Criteria               string
	FarcasterCriteria      *FarcasterCriteria `gorm:"foreignKey:DropId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserLimit              *int               `gorm:"default:null"`
	EditionLimit           *int               `gorm:"default:null"`
	MintPasses             []MintPass         `gorm:"foreignKey:DropID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
		// ClaimCriteria:         c.Criteria,
	}

	if c.Criteria != "" {
		var claimCriterias []*model.ClaimCriteriaType
		criterias := strings.Split(c.Criteria, ",")
		for _, criteria := range criterias {
			cr := model.ClaimCriteriaType(criteria)
			claimCriterias = append(claimCriterias, &cr)
		}
		mappedDrop.ClaimCriteria = claimCriterias
	}

	if c.FarcasterCriteria != nil {
		var claimCriteriaInteractions []*model.InteractionType
		interactions := strings.Split(c.FarcasterCriteria.Interactions, ",")
		for _, interaction := range interactions {
			i := model.InteractionType(interaction)
			claimCriteriaInteractions = append(claimCriteriaInteractions, &i)
		}
		mappedDrop.FarcasterClaimCriteriaInteractions = claimCriteriaInteractions
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
		mappedDrop.FarcasterProfileID = &c.FarcasterCriteria.FarcasterUsername
		mappedDrop.FarcasterChannelID = &c.FarcasterCriteria.ChannelID
	}

	var mintPasses []*model.ClaimDetails
	for j := range c.MintPasses {
		mintPasses = append(mintPasses, c.MintPasses[j].ToGraphData())
	}
	mappedDrop.ClaimDetails = mintPasses
	if items != nil {
		for i := 0; i < len(items); i++ {
			items[i].DropAddress = *c.AAContractAddress
			// items[i].ClaimDetails = mintPasses
		}
		mappedDrop.Items = items
	}
	// fmt.Println(mappedDrop.Items[0].ClaimDetails[0].ClaimerAddress)
	return mappedDrop
}
