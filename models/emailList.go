package models

import uuid "github.com/satori/go.uuid"

type SingleEmailClaim struct {
	BaseWithoutPrimaryKey
	CreatorID    uuid.UUID `gorm:"primaryKey"`
	ItemID       uuid.UUID `gorm:"primaryKey"`
	EmailAddress string    `gorm:"primaryKey"`
}

// incase we want to enforce single domain whitelists remote the BaseDomain from the compositeKey
type EmailDomainWhiteList struct {
	BaseWithoutPrimaryKey
	ItemID     uuid.UUID `gorm:"primaryKey"`
	CreatorID  uuid.UUID `gorm:"primaryKey"`
	BaseDomain string    `gorm:"primaryKey"`
}
