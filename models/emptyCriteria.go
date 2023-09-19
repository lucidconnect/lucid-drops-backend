package models

type EmptyCriteria struct {
	Base
	ItemID       string `gorm:"index"`
	CreatorID    string `gorm:"index"`
}
