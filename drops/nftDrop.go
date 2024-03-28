package drops

import (
	"strings"
	"time"

	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/utils"
	uuid "github.com/satori/go.uuid"
)

type NFTRepository interface {
	CreateDrop(drop *Drop, items *Item) error
	FindDropById(dropId string) (*Drop, error)
	UpdateDropDetails(drop *Drop) error
	FindDropByCreatorId(creatorId string) ([]Drop, error)
	DeleteDrop(drop *Drop) error
	AddFarcasterCriteriaToDrop(drop *Drop, criteria *FarcasterCriteria) error
	UpdateFarcasterCriteria(dropId string, criteriaUpdate *FarcasterCriteria) error
	RemoveFarcasterCriteria(drop *Drop) error
	FetchDropItems(dropId string, includeDeleted bool) ([]Item, error)
	FindFeaturedDrops() ([]Drop, error)
}

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

type Item struct {
	Base
	Name          string
	DropID        uuid.UUID `gorm:"index"`
	DropAddress   string
	TokenID       int64  `gorm:"index"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	ClaimFee      int    `gorm:"default:0"`
	Criteria      *model.ClaimCriteriaType
	ClaimDeadline *time.Time `gorm:"default:null"`
	// TwitterCriteria      *TwitterCriteria  `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// TelegramCriteria     *TelegramCriteria `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// PatreonCriteria      *PatreonCriteria  `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ShowEmailDomainHints bool `gorm:"default:false"`
	Featured             bool `gorm:"default:false"`
	UserLimit            *int `gorm:"default:null"`
	EditionLimit         *int `gorm:"default:null"`
	// MintPasses           []MintPass        `gorm:"foreignKey:ItemId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Ref : https://docs.opensea.io/docs/metadata-standards
type OpenSeaMetaDataFormat struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	ExternalUrl string `json:"external_url"`
}

func (d *Drop) ToGraphData(items []*model.Item) *model.Drop {
	mappedDrop := &model.Drop{
		ID:                    d.ID.String(),
		CreatorID:             d.CreatorID.String(),
		CreatedAt:             d.CreatedAt,
		Name:                  d.Name,
		Description:           d.Description,
		Image:                 d.Image,
		Thumbnail:             d.Thumbnail,
		ContractAddress:       d.AAContractAddress,
		Network:               d.BlockchainNetwork,
		MintURL:               d.MintUrl,
		GasIsCreatorSponsored: d.GasIsCreatorSponsored,
		// ClaimCriteria:         d.Criteria,
	}

	var claimCriterias []*model.ClaimCriteriaType
	if d.Criteria != "" {
		criterias := strings.Split(d.Criteria, ",")
		for _, criteria := range criterias {
			cr := model.ClaimCriteriaType(criteria)
			claimCriterias = append(claimCriterias, &cr)
		}
		mappedDrop.ClaimCriteria = claimCriterias
	}

	if d.AAContractAddress != nil {
		mappedDrop.ContractAddress = d.AAContractAddress
	}

	if d.MintPrice != nil {
		mappedDrop.MintPrice = d.MintPrice
	}

	if d.FarcasterCriteria != nil {
		mappedDrop.FarcasterClaimCriteriaInteractions = InteractionsToArr(d.FarcasterCriteria.Interactions)
		mappedDrop.CastURL = &d.FarcasterCriteria.CastUrl
		mappedDrop.FarcasterProfileID = &d.FarcasterCriteria.FarcasterProfileID
		mappedDrop.FarcasterChannelID = &d.FarcasterCriteria.ChannelID
	}

	var mintPasses []*model.ClaimDetails
	for j := range d.MintPasses {
		mintPasses = append(mintPasses, d.MintPasses[j].ToGraphData())
	}
	mappedDrop.ClaimDetails = mintPasses
	if items != nil {
		for i := 0; i < len(items); i++ {
			items[i].DropAddress = *d.AAContractAddress
			// items[i].ClaimDetails = mintPasses
		}
		mappedDrop.Items = items
	}
	// fmt.Println(mappedDrop.Items[0].ClaimDetails[0].ClaimerAddress)
	return mappedDrop
}

func (i *Item) ToGraphData() *model.Item {
	tokenID := utils.GetIntPtr(int(i.TokenID))

	item := &model.Item{
		ID:           i.ID.String(),
		Name:         i.Name,
		Image:        i.Image,
		Description:  i.Description,
		DropID:       i.DropID.String(),
		DropAddress:  i.DropAddress,
		ClaimFee:     i.ClaimFee,
		CreatedAt:    i.CreatedAt,
		Deadline:     i.ClaimDeadline,
		EditionLimit: i.EditionLimit,
		TokenID:      tokenID,
	}

	// var mintPasses []*model.ClaimDetails
	// for j := range i.MintPasses {
	// 	mintPasses = append(mintPasses, i.MintPasses[j].ToGraphData())
	// }

	// item.ClaimDetails = mintPasses
	return item
}

func InteractionsToArr(interaction string) []*model.InteractionType {
	interactions := strings.Split(interaction, ",")
	if len(interactions) == 0 {
		return nil
	}

	var interactionArr []*model.InteractionType
	for i := range interactions {
		if interactions[i] == "" {
			continue
		}

		interactionArr = append(interactionArr, (*model.InteractionType)(&interactions[i]))
	}

	return interactionArr
}
