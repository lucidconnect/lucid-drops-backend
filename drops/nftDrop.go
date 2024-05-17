package drops

import (
	"strings"
	"time"

	"github.com/lucidconnect/inverse/graph/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type NFTRepository interface {
	CreateDrop(drop *Drop, items *Item) error
	FindDropById(dropId string) (*Drop, error)
	UpdateDropDetails(drop *Drop) error
	FindDropByCreatorId(creatorId string) ([]Drop, error)
	FindItemById(itemId string) (*Item, error)
	UpdateItemDetails(item *Item) error
	DeleteDrop(drop *Drop) error
	AddFarcasterCriteriaToDrop(drop *Drop, criteria *FarcasterCriteria) error
	UpdateFarcasterCriteria(dropId string, criteriaUpdate *FarcasterCriteria) error
	RemoveFarcasterCriteria(dropId string) error
	FetchDropItems(dropId string, includeDeleted bool) ([]Item, error)
	FindFeaturedDrops() ([]Drop, error)
	FindClaimedDropsByAddress(addresss string) ([]Item, error)
	CreateMintPass(mintPass *MintPass) error
	UpdateMintPass(mintPass *MintPass) error
	GetMintPassById(passId string) (*MintPass, error)
	GetMintPassForWallet(dropId, walletAddress string) (*MintPass, error)
	GetMintPassesForWallet(dropId, walletAddress string) (int64, error)
	CountMintPassesForAddress(dropId, address string) (int64, error)
	CountMintPassesForDrop(dropId string) (int64, error)
	FetchMintPassesForItems(itemID string) ([]MintPass, error)
	FindItemsWithUnresolvesTokenIDs() ([]Item, error)
	CreateMetadata(metadata *MetaData) error
	ReadMetadata(id string) (*MetaData, error)
	ReadMetadataByDropId(dropId, tokenId string) (*MetaData, error)
}

type Drop struct {
	ID                     uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt              time.Time      `gorm:"not null"`
	UpdatedAt              time.Time      `gorm:"not null"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
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
	DropUri                string
}

func (d *Drop) BeforeCreate(scope *gorm.DB) error {
	d.ID = uuid.NewV4()
	return nil
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
		EditionLimit:          d.EditionLimit,
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
		mappedDrop.FarcasterProfileID = &d.FarcasterCriteria.FarcasterUsername
		channelIds := strings.Split(d.FarcasterCriteria.ChannelID, ",")
		mappedDrop.FarcasterChannelID = channelIds
	}

	if d.DropUri != "" {
		mappedDrop.URI = &d.DropUri
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
