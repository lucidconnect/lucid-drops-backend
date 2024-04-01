package models

import (
	"time"

	"github.com/lucidconnect/inverse/graph/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type TwitterCriteria struct {
	ID               uuid.UUID               `gorm:"type:uuid;primary_key;"`
	CreatedAt        time.Time               `gorm:"not null"`
	UpdatedAt        time.Time               `gorm:"not null"`
	DeletedAt        gorm.DeletedAt          `gorm:"index"`
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
