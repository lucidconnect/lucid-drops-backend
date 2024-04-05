package drops

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type FarcasterCriteria struct {
	ID                 uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt          time.Time      `gorm:"not null"`
	UpdatedAt          time.Time      `gorm:"not null"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	DropId             uuid.UUID      `gorm:"not null"`
	CreatorID          uuid.UUID
	CastUrl            string
	CriteriaType       string
	Interactions       string
	ChannelID          string
	FarcasterProfileID string // fid of account to do verifications against
	FarcasterUsername  string
}

func (fc *FarcasterCriteria) BeforeCreate(scope *gorm.DB) error {
	fc.ID = uuid.NewV4()
	return nil
}
