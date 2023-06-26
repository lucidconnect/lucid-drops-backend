package models

type TelegramCriteria struct {
	Base
	ItemID      string `gorm:"type:uuid;index:idx_item_id,unique;not null;"`
	CreatorID   string `gorm:"primaryKey"`
	GroupID     int64  `gorm:"column:group_id"`
}
