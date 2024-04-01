package ledger

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/lucidconnect/inverse/drops"
	"github.com/lucidconnect/inverse/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseWithoutPrimaryKey struct {
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DoubleEntryLedger struct {
	ID                   uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt            time.Time      `gorm:"not null"`
	UpdatedAt            time.Time      `gorm:"not null"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	TransactionReference string         `gorm:"not null" json:"transaction_reference"`
	SourceAccoountID     uuid.UUID
	DestinationAccountID uuid.UUID
	Amount               int64  `gorm:"not null" json:"amount"`
	TransactionType      string `gorm:"not null" json:"transaction_type"`
	//ID of the corresponding double entry row
	PartnerID *string `json:"partner_id"`
	LedgerID  *string `json:"ledger_id"`
}

func (del *DoubleEntryLedger) BeforeCreate(scope *gorm.DB) error {
	del.ID = uuid.NewV4()
	return nil
}

type Wallet struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	CreatorID     string         `gorm:"type:uuid;index:idx_wallet_creatorId,unique;not null" json:"creator_id"`
	BalanceBase   int64          `gorm:"type:bigint;default:0" json:"balance_base"`
	CanBeNegative bool           `gorm:"default:false"`
	// Currency      CurrencyType `gorm:"default:USD" json:"currency"`
}

func (w *Wallet) BeforeCreate(scope *gorm.DB) error {
	w.ID = uuid.NewV4()
	return nil
}

type Ledger struct {
	DB             *gorm.DB
	SysAccount     *Wallet
	CollectAccount *Wallet
}

func New(db *gorm.DB) *Ledger {

	sysAccount, err := confirmOrSeedNewSysAccount(db)
	if err != nil {
		log.Print("Error confirming or seeding system account", err)
		panic(err)
	}

	dropsAccount, err := confirmOrSeedNewDropsAccount(db)
	if err != nil {
		log.Print("Error confirming or seeding drops account", err)
		panic(err)
	}

	return &Ledger{
		DB:             db,
		SysAccount:     sysAccount,
		CollectAccount: dropsAccount,
	}
}

func confirmOrSeedNewSysAccount(db *gorm.DB) (*Wallet, error) {
	var sysAccount Wallet

	err := db.Where("can_be_negative = ?", "true").First(&sysAccount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			sysUserID := uuid.FromStringOrNil("8288925B-9AD8-431D-AAF0-1A6655727CDC")
			sysUser := drops.Creator{
				ID:              sysUserID,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
				InverseUsername: utils.GetStrPtr("SystemAccount"),
				WalletAddress:   "sambankmanfried@tippr.com",
			}

			err = db.Create(&sysUser).Error
			if err != nil {
				return nil, err
			}

			sysAccount = Wallet{
				ID:            uuid.NewV4(),
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				CreatorID:     sysUser.ID.String(),
				BalanceBase:   0,
				CanBeNegative: true,
			}

			err = db.Create(&sysAccount).Error
			if err != nil {
				return nil, err
			}

		} else {
			return nil, err
		}
	}
	return &sysAccount, nil
}

func confirmOrSeedNewDropsAccount(db *gorm.DB) (*Wallet, error) {
	var dropsAccount Wallet
	var dropsAccountUser drops.Creator
	collectUserID := uuid.FromStringOrNil("97B91EAF-3EE1-4F9A-836B-6B49E7B0AC9E")
	collectUser := drops.Creator{
		ID:              collectUserID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		InverseUsername: utils.GetStrPtr("DropAccount"),
		WalletAddress:   "warrenbuffetcollects@inverse.wtf",
	}

	err := db.Model(&drops.Creator{}).Where("wallet_address = ?", "warrenbuffetcollects@inverse.wtf").First(&dropsAccountUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			err = db.Create(&collectUser).Error
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	err = db.Model(&Wallet{}).Where("creator_id = ?", collectUser.ID.String()).First(&dropsAccount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			dropsAccount = Wallet{
				ID:            uuid.NewV4(),
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				CreatorID:     collectUser.ID.String(),
				BalanceBase:   0,
				CanBeNegative: false,
			}

			err = db.Create(&dropsAccount).Error
			if err != nil {
				return nil, err
			}

		} else {
			return nil, err
		}
	}

	return &dropsAccount, nil
}
