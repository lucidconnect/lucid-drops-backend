package models

import (
	"github.com/lucidconnect/inverse/graph/model"
	uuid "github.com/satori/go.uuid"
)

type FarcasterCriteria struct {
	Base
	DropId             uuid.UUID `gorm:"unique;not null"`
	CreatorID          uuid.UUID
	CastUrl            string
	CriteriaType       model.ClaimCriteriaType
	Interactions       string
	ChannelID          string
	FarcasterProfileID string // fid of account to do verifications against
}
