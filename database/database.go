package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/lucidconnect/inverse/drops"
	"github.com/lucidconnect/inverse/ledger"
	"github.com/lucidconnect/inverse/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	database *gorm.DB
}

func SetupDB(dsn string) *DB {
	var err error

	log.Debug().Msgf("Initialising Database...")

	isProd, _ := utils.IsProduction()
	ormConfig := &gorm.Config{
		PrepareStmt: false,
	}

	//switch to silent mode in production
	if isProd {
		log.Print("SQL Log level set to silent")
		ormConfig.Logger = logger.Default.LogMode(logger.Silent)
	} else {
		log.Print("SQL Log level set to Error")
		ormConfig.Logger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(postgres.Open(dsn), ormConfig)

	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(90)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Print("Successfully connected!")
	db.AutoMigrate(
		&drops.Creator{},
		&ledger.Wallet{},
		&drops.SignerInfo{},
	)
	return &DB{database: db}
}

func (db *DB) CreateProfile(creator *drops.Creator, signer *drops.SignerInfo) error {
	tx := db.database.Begin()
	if err := tx.Create(creator).Error; err != nil {
		tx.Rollback()
		return err
	}

	signer.CreatorID = creator.ID.String()
	if err := tx.Create(signer).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (db *DB) FindCreatorById(creatorId string) (*drops.Creator, error) {
	var creator drops.Creator
	err := db.database.Where("id = ?", creatorId).First(&creator).Error
	if err != nil {
		return nil, fmt.Errorf("creator with id (%s) not found", creatorId)
	}
	return &creator, nil
}

func (db *DB) FindCreatorByUsername(creatorUsername string) (*drops.Creator, error) {
	var creator drops.Creator
	err := db.database.Where("inverse_username = ?", creatorUsername).First(&creator).Error
	if err != nil {
		return nil, fmt.Errorf("username (%s) not in use", creatorUsername)
	}
	return &creator, nil
}

func (db *DB) FindCreatorByEthereumAddress(address string) (*drops.Creator, error) {
	var creator drops.Creator
	query := fmt.Sprintf("SELECT * FROM creators WHERE LOWER(wallet_address)='%s'", strings.ToLower(address))
	err := db.database.Raw(query).First(&creator).Error
	if err != nil {
		return nil, err
	}
	return &creator, nil
}

func (db *DB) FindSignerByCreatorId(creatorId string) (*drops.SignerInfo, error) {
	var altSigner drops.SignerInfo

	err := db.database.Where("creator_id=?", creatorId).First(&altSigner).Error
	if err != nil {
		return nil, fmt.Errorf("signer for creator (%s) not found", creatorId)
	}

	return &altSigner, nil
}

func (db *DB) FindSignerByEthereumAddress(address string) (*drops.SignerInfo, error) {
	var altSigner drops.SignerInfo

	err := db.database.Where("wallet_address=?", address).First(&altSigner).Error
	if err != nil {
		return nil, fmt.Errorf("signer for address (%s) not found", address)
	}

	return &altSigner, nil
}

func (db *DB) UpdateCreatorProfile(creator *drops.Creator, signer *drops.SignerInfo) error {
	tx := db.database.Begin()
	if err := tx.Save(creator).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(signer).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (db *DB) CreateDrop(drop *drops.Drop, item *drops.Item) error {
	tx := db.database.Begin()
	if err := tx.Create(drop).Error; err != nil {
		tx.Rollback()
		return err
	}

	item.DropID = drop.ID
	if err := tx.Create(item).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
func (db *DB) FindDropById(dropId string) (*drops.Drop, error) {
	var drop drops.Drop

	if err := db.database.Where("id = ?", dropId).Preload("FarcasterCriteria").Preload("MintPasses").First(&drop).Error; err != nil {
		return nil, err
	}

	return &drop, nil
}

func (db *DB) UpdateDropDetails(drop *drops.Drop) error {
	return db.database.Save(drop).Error
}

func (db *DB) DeleteDrop(drop *drops.Drop) error {
	return db.database.Delete(drop).Error
}

func (db *DB) AddFarcasterCriteriaToDrop(drop *drops.Drop, criteria *drops.FarcasterCriteria) error {
	tx := db.database.Begin()
	if err := tx.Save(drop).Error; err != nil {
		return err
	}
	if err := tx.Create(criteria).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (db *DB) UpdateFarcasterCriteria(dropId string, criteriaUpdate *drops.FarcasterCriteria) error {
	return nil
}

func (db *DB) RemoveFarcasterCriteria(dropId string) error {
	if err := db.database.Unscoped().Where("drop_id = ?", dropId).Delete(&drops.FarcasterCriteria{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) FetchDropItems(dropId string, includeDeleted bool) ([]drops.Item, error) {
	var items []drops.Item

	if includeDeleted {
		if err := db.database.Unscoped().Where("drop_id = ?", dropId).Find(&items).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.database.Where("drop_id=?", dropId).Find(&items).Error; err != nil {
			return nil, err
		}
	}
	return items, nil
}

func (db *DB) FindDropByCreatorId(creatorId string) ([]drops.Drop, error) {
	var drops []drops.Drop

	if err := db.database.Where("creator_id=?", creatorId).Find(&drops).Error; err != nil {
		return nil, err
	}
	return drops, nil
}

func (db *DB) FindFeaturedDrops() ([]drops.Drop, error) {
	var drops []drops.Drop

	if err := db.database.Where("featured=?", true).Find(&drops).Error; err != nil {
		return nil, err
	}

	return drops, nil
}

func (db *DB) FindClaimedDropsByAddress(addresss string) ([]drops.Item, error) {
	var mintPasses []drops.MintPass
	var claimedItems []drops.Item

	err := db.database.Model(&drops.MintPass{}).Where("minter_address=?", addresss).Find(&mintPasses).Error
	if err != nil {
		return nil, err
	}

	itemIds := make([]string, len(mintPasses))

	for i, pass := range mintPasses {
		itemIds[i] = pass.ItemId
	}

	err = db.database.Where("id IN (?)", itemIds).Find(&claimedItems).Error
	if err != nil {
		return nil, err
	}

	return claimedItems, nil
}

func (db *DB) CreateMintPass(mintPass *drops.MintPass) error {
	return db.database.Create(mintPass).Error
}

func (db *DB) GetMintPassById(passId string) (*drops.MintPass, error) {
	var pass drops.MintPass

	err := db.database.Where("id=?", passId).First(&pass).Error
	if err != nil {
		return nil, err
	}

	return &pass, nil
}

func (db *DB) GetMintPassForWallet(dropId, walletAddress string) (*drops.MintPass, error) {
	var pass drops.MintPass

	err := db.database.Where("drop_id = ?", dropId).Where("minter_address = ?", walletAddress).First(&pass).Error
	if err != nil {
		return nil, err
	}

	return &pass, nil
}

func (db *DB) CountMintPassesForAddress(dropId, address string) (int64, error) {
	var passes int64
	err := db.database.Model(&drops.MintPass{}).Where("drop_id = ? AND minter_address = ? AND used_at <> NULL", dropId, address).Count(&passes).Error
	if err != nil {
		return 0, err
	}
	return passes, nil
}

func (db *DB) UpdateMintPass(mintPass *drops.MintPass) error {
	return db.database.Save(mintPass).Error
}

func (db *DB) FetchMintPassesForItems(itemID string) ([]drops.MintPass, error) {

	var mintPasses []drops.MintPass
	err := db.database.Model(&drops.MintPass{}).Where("item_id = ?", itemID).Find(&mintPasses).Error
	if err != nil {
		return nil, err
	}

	return mintPasses, nil
}

func (db *DB) FindItemsWithUnresolvesTokenIDs() ([]drops.Item, error) {
	var items []drops.Item
	oneHourAgo := time.Now().Add(-1 * time.Hour)
	err := db.database.Where("token_id is NULL and created_at BETWEEN ? AND ?", oneHourAgo, time.Now()).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (db *DB) FindItemById(itemId string) (*drops.Item, error) {
	var item drops.Item

	if err := db.database.Where("id = ?", itemId).First(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (db *DB) CountMintPassesForDrop(dropId string) (int64, error) {
	var editionCount int64
	err := db.database.Model(&drops.MintPass{}).Where("drop_id = ? AND used_at <> NULL", dropId).Count(&editionCount).Error
	if err != nil {
		return 0, err
	}
	return editionCount, nil
}

func (db *DB) GetMintPassesForWallet(dropId, walletAddress string) (int64, error) {
	var walletCount int64
	err := db.database.Model(&drops.MintPass{}).Where("drop_id = ? AND used_at <> NULL", dropId).Count(&walletCount).Error
	if err != nil {
		return 0, err
	}
	return walletCount, nil
}
