package models

import (
	uuid "github.com/satori/go.uuid"
	"inverse.so/graph/model"
)

type Collection struct {
	Base
	CreatorID  uuid.UUID
	Name       string
	ContentURI string // TODO add support for multiple content types
}

func (c *Collection) ToGraphData() *model.Collection {
	return &model.Collection{
		ID:         c.ID.String(),
		Name:       c.Name,
		ContentURI: c.ContentURI,
	}
}
