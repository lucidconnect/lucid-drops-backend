package models

import (
	"github.com/lucidconnect/inverse/graph/model"
	uuid "github.com/satori/go.uuid"
)

type FarcasterCriteria struct {
	Base
	DropID             uuid.UUID `gorm:"unique;index:idx_drop_id;not null"`
	CreatorID          uuid.UUID
	CastUrl            string
	CriteriaType       model.ClaimCriteriaType
	Interactions       string
	ChannelID          string
	FarcasterProfileID string // fid of account to do verifications against
}
