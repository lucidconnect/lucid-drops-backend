package ledger

import (
	"time"

	"github.com/rs/zerolog/log"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"inverse.so/models"
	"inverse.so/utils"
)

type Ledger struct {
	DB         *gorm.DB
	SysAccount *models.Wallet
}

func New(db *gorm.DB) *Ledger {

	sysAccount, err := confirmOrSeedNewSysAccount(db)
	if err != nil {
		log.Print("Error confirming or seeding system account", err)
		panic(err)
	}

	return &Ledger{
		DB:         db,
		SysAccount: sysAccount,
	}
}

func confirmOrSeedNewSysAccount(db *gorm.DB) (*models.Wallet, error) {
	var sysAccount models.Wallet

	err := db.Where("can_be_negative = ?", "true").First(&sysAccount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			sysUserID := uuid.FromStringOrNil("8288925B-9AD8-431D-AAF0-1A6655727CDC")
			sysUser := models.Creator{
				Base: models.Base{
					ID:        sysUserID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				InverseUsername: utils.GetStrPtr("SystemAccount"),
				WalletAddress:   "sambankmanfried@tippr.com",
			}

			err = db.Create(&sysUser).Error
			if err != nil {
				return nil, err
			}

			sysAccount = models.Wallet{
				Base: models.Base{
					ID:        uuid.NewV4(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				CreatorID:     sysUser.ID.String(),
				BalanceBase:   0,
				CanBeNegative: true,
				Currency:      models.USD,
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
