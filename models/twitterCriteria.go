package models

import "inverse.so/graph/model"

type TwitterCriteria struct {
	Base
	ItemID           string                  `gorm:"type:uuid;index:idx_item_id,unique;not null;"`
	CreatorID        string                  `gorm:"primaryKey"`
	ProfileLink      string                  `gorm:"column:profile_link"`
	ProfileID        string                  `gorm:"column:profile_id"`
	TweetLink        string                  `gorm:"column:tweet_link"`
	TweetID          string                  `gorm:"column:tweet_id"`
	AuthID           string                  `gorm:"column:auth_id"`
	CriteriaType     model.ClaimCriteriaType `gorm:"column:criteria_type"`
	Interactions     string                  `gorm:"column:interaction"`
	CutOffDate       string                  `gorm:"column:cutoff_date"`
	IndexedFollowers string                  `gorm:"column:indexed_followers"`
	IndexedRetweets  string                  `gorm:"column:indexed_retweets"`
}
