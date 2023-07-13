package models

type PatreonCriteria struct {
	Base
	AuthID       string `gorm:"index"`
	ItemID       string `gorm:"index"`
	CreatorID    string `gorm:"index"`
	CampaignName string `gorm:"index"`
}
