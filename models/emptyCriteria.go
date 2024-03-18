package models

type EmptyCriteria struct {
	Base
	DropID       string `gorm:"index"`
	CreatorID    string `gorm:"index"`
}
