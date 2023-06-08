package models

type TwitterCriteria struct {
	Base
	ItemID       string `gorm:"type:uuid;index:idx_item_id,unique;not null;"`
	CreatorID    string `gorm:"primaryKey"`
	ProfileLink  string `gorm:"column:profile_link"`
	TweetLink    string `gorm:"column:tweet_link"`
	TweetID      string `gorm:"column:tweet_id"`
	Interactions string `gorm:"column:interaction"`
	CutOffDate   string `gorm:"column:cutoff_date"`
}
