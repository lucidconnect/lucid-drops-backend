package drops

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type MetaData struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	DropId      uuid.UUID `gorm:"index;"`
	ItemId      uuid.UUID
	TokenId     string
	Name        string
	Description string
	Properties  any
}

func (md *MetaData) BeforeCreate(scope *gorm.DB) error {
	md.ID = uuid.NewV4()
	return nil
}

// name,description,properties[Mouth],properties[Clothes],properties[Head],properties[Eyes]
