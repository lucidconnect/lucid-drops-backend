package models

import (
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
	ShowEmailDomainHints bool `gorm:"default:false"`
}

func (i *Item) ToGraphData() *model.Item {
	return &model.Item{
		ID:            i.ID.String(),
		Name:          i.Name,
		Image:         i.Image,
		Description:   i.Description,
		CollectionID:  i.CollectionID.String(),
		ClaimCriteria: i.Criteria,
	}
}
