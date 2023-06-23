package models

import (
	uuid "github.com/satori/go.uuid"
	"inverse.so/graph/model"
)

type Collection struct {
	Base
	CreatorID   uuid.UUID
	Name        string
	Image       string `json:"image"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

func (c *Collection) ToGraphData() *model.Collection {
	testAddress := "0xdEFdf3C58e416aD22b4b47613A41ce1e6050B13B"
	return &model.Collection{
		ID:              c.ID.String(),
		Name:            c.Name,
		Description:     c.Description,
		Image:           c.Image,
		Thumbnail:       c.Thumbnail,
		ContractAddress: &testAddress,
	}
}
