package models

import (
	uuid "github.com/satori/go.uuid"
)

type FarcasterCriteria struct {
	Base
	DropId             uuid.UUID `gorm:"not null"`
	CreatorID          uuid.UUID
	CastUrl            string
	CriteriaType       string
	Interactions       string
	ChannelID          string
	FarcasterProfileID string // fid of account to do verifications against
}
