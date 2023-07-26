package models

import (
	"strings"

	uuid "github.com/satori/go.uuid"
	"inverse.so/graph/model"
	"inverse.so/utils"
)

type Item struct {
	Base
	Name                 string
	CollectionID         uuid.UUID `gorm:"index"`
	Image                string    `json:"image"`
	Description          string    `json:"description"`
	Criteria             *model.ClaimCriteriaType
	TwitterCriteria      *TwitterCriteria  `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TelegramCriteria     *TelegramCriteria `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PatreonCriteria      *PatreonCriteria  `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ShowEmailDomainHints bool              `gorm:"default:false"`
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
		ExternalUrl: utils.UseEnvOrDefault("INVERSE_API_BASEURL", "https://localhost:8090"),
	}

	return openSeaMetadata
}

func (i *Item) ToGraphData() *model.Item {

	item := &model.Item{
		ID:            i.ID.String(),
		Name:          i.Name,
		Image:         i.Image,
		Description:   i.Description,
		CollectionID:  i.CollectionID.String(),
		ClaimCriteria: i.Criteria,
		CreatedAt:     i.CreatedAt,
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
