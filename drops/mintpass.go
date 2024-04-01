package drops

import (
	"time"

	"github.com/lucidconnect/inverse/graph/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type MintPass struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	ItemId        string
	DropID        string
	MinterAddress string
	TokenID       string
	UsedAt        *time.Time `gorm:"default:null"`
	// ItemIdOnContract    int64
	DropContractAddress string
	BlockchainNetwork   *model.BlockchainNetwork `gorm:"default:base"`
}

func (m *MintPass) BeforeCreate(scope *gorm.DB) error {
	m.ID = uuid.NewV4()
	return nil
}

func (m *MintPass) ToGraphData() *model.ClaimDetails {

	//TODO: find a way arount import cycle error
	// var username *string = nil
	// creator, err := graph.FetchCreatorByAddress(common.HexToAddress(m.MinterAddress))
	// if err == nil && creator.InverseUsername != nil {
	// 	username = creator.InverseUsername
	// }
	return &model.ClaimDetails{
		ClaimerAddress: &m.MinterAddress,
		ClaimTime:      &m.CreatedAt,
		// ClaimerUsername: username,
	}
}
