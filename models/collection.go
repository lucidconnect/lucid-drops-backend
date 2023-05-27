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
	return &model.Collection{
		ID:          c.ID.String(),
		Name:        c.Name,
		Description: c.Description,
		Image:       c.Image,
		Thumbnail:   c.Thumbnail,
	}
}
