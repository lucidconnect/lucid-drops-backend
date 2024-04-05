package drops

import (
	"time"

	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Name          string
	DropID        uuid.UUID `gorm:"index"`
	DropAddress   string
	TokenID       int64  `gorm:"index"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	ClaimFee      int    `gorm:"default:0"`
	Criteria      *model.ClaimCriteriaType
	ClaimDeadline *time.Time `gorm:"default:null"`
	// TwitterCriteria      *TwitterCriteria  `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// TelegramCriteria     *TelegramCriteria `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// PatreonCriteria      *PatreonCriteria  `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ShowEmailDomainHints bool `gorm:"default:false"`
	Featured             bool `gorm:"default:false"`
	UserLimit            *int `gorm:"default:null"`
	EditionLimit         *int `gorm:"default:null"`
	// MintPasses           []MintPass        `gorm:"foreignKey:ItemId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (i *Item) BeforeCreate(scope *gorm.DB) error {
	i.ID = uuid.NewV4()
	return nil
}

func (i *Item) ToGraphData() *model.Item {
	tokenID := utils.GetIntPtr(int(i.TokenID))

	item := &model.Item{
		ID:           i.ID.String(),
		Name:         i.Name,
		Image:        i.Image,
		Description:  i.Description,
		DropID:       i.DropID.String(),
		DropAddress:  i.DropAddress,
		ClaimFee:     i.ClaimFee,
		CreatedAt:    i.CreatedAt,
		Deadline:     i.ClaimDeadline,
		EditionLimit: i.EditionLimit,
		TokenID:      tokenID,
	}

	// var mintPasses []*model.ClaimDetails
	// for j := range i.MintPasses {
	// 	mintPasses = append(mintPasses, i.MintPasses[j].ToGraphData())
	// }

	// item.ClaimDetails = mintPasses
	return item
}
