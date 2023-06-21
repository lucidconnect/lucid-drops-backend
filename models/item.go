package models

import (
	"strings"

	uuid "github.com/satori/go.uuid"
	"inverse.so/graph/model"
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
	ShowEmailDomainHints bool              `gorm:"default:false"`
}

func (i *Item) ToGraphData() *model.Item {

	item := &model.Item{
		ID:            i.ID.String(),
		Name:          i.Name,
		Image:         i.Image,
		Description:   i.Description,
		CollectionID:  i.CollectionID.String(),
		ClaimCriteria: i.Criteria,
	}

	if i.TwitterCriteria != nil {
		item.TwitterClainCriteriaInteractions = interactionsToArr(i.TwitterCriteria.Interactions)
	}

	return item
}

func interactionsToArr(interaction string) []*model.InteractionType {
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
