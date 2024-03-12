package models

import (
	"strings"
	"time"

	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/utils"
	uuid "github.com/satori/go.uuid"
)

type Item struct {
	Base
	Name                 string
	DropID               uuid.UUID `gorm:"index"`
	DropAddress          string
	TokenID              *int64 `gorm:"index;default:null"`
	Image                string `json:"image"`
	Description          string `json:"description"`
	ClaimFee             int    `gorm:"default:0"`
	Criteria             *model.ClaimCriteriaType
	ClaimDeadline        *time.Time        `gorm:"default:null"`
	TwitterCriteria      *TwitterCriteria  `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TelegramCriteria     *TelegramCriteria `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PatreonCriteria      *PatreonCriteria  `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ShowEmailDomainHints bool              `gorm:"default:false"`
	Featured             bool              `gorm:"default:false"`
	UserLimit            *int              `gorm:"default:null"`
	EditionLimit         *int              `gorm:"default:null"`
	MintPasses           []MintPass        `gorm:"foreignKey:ItemId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Ref : https://docs.opensea.io/docs/metadata-standards
type OpenSeaMetaDataFormat struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	ExternalUrl string `json:"external_url"`
}

func (i *Item) ToOpenSeaMetadata() *OpenSeaMetaDataFormat {
	openSeaMetadata := &OpenSeaMetaDataFormat{
		Name:        i.Name,
		Image:       i.Image,
		Description: i.Description,
		ExternalUrl: utils.UseEnvOrDefault("API_BASE_URL", "https://localhost:8090"),
	}

	return openSeaMetadata
}

func (i *Item) ToGraphData() *model.Item {

	var tokenID *int
	if i.TokenID != nil {
		tokenID = utils.GetIntPtr(int(*i.TokenID))
	}

	item := &model.Item{
		ID:            i.ID.String(),
		Name:          i.Name,
		Image:         i.Image,
		Description:   i.Description,
		DropID:        i.DropID.String(),
		ClaimCriteria: i.Criteria,
		ClaimFee:      i.ClaimFee,
		CreatedAt:     i.CreatedAt,
		Deadline:      i.ClaimDeadline,
		EditionLimit:  i.EditionLimit,
		TokenID:       tokenID,
	}

	if i.TwitterCriteria != nil {
		item.TwitterClaimCriteriaInteractions = InteractionsToArr(i.TwitterCriteria.Interactions)
		item.TweetLink = &i.TwitterCriteria.TweetLink
		item.ProfileLink = &i.TwitterCriteria.ProfileLink
	}

	if i.PatreonCriteria != nil {
		item.CampaignName = &i.PatreonCriteria.CampaignName
	}

	if i.TelegramCriteria != nil {
		item.TelegramGroupTitle = &i.TelegramCriteria.GroupTitle
	}

	var mintPasses []*model.ClaimDetails
	for j := range i.MintPasses {
		mintPasses = append(mintPasses, i.MintPasses[j].ToGraphData())
	}

	item.ClaimDetails = mintPasses
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
