package drops

import (
	"time"

	"github.com/lucidconnect/inverse/graph/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SignerInfo struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	CreatorID     string         `gorm:"type:uuid;index:idx_creator_id,unique;not null;"`
	WalletAddress string         `gorm:"type:varchar(255);index:idx_wallet_address,unique;not null;"`
	Signature     *string
	AltPublicKey  string
	AltPrivateKey string
	Provider      model.SignerProvider
}

func (s *SignerInfo) BeforeCreate(scope *gorm.DB) error {
	s.ID = uuid.NewV4()
	return nil
}
