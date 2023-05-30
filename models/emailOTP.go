package models

import uuid "github.com/satori/go.uuid"

type EmailOTP struct {
	Base
	IssuedAt        int64     `gorm:"not null"`
	ExpiresAt       int64     `gorm:"not null"`
	ItemID          uuid.UUID `gorm:"index"`
	UserEmail       string
	ClaimingAddress string
	ExpectedOTP     string
	Attempts        int
}
