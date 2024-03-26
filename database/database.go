package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/lucidconnect/inverse/drops"
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
		&drops.Wallet{},
		&drops.SignerInfo{},
	)
	return &DB{database: db}
}

func (db *DB) CreateProfile(creator *drops.Creator, signer *drops.SignerInfo) error {
	tx := db.database.Begin()
	if err := tx.Create(creator).Error; err != nil {
		return err
	}

	signer.CreatorID = creator.ID.String()
	if err := tx.Create(signer).Error; err != nil {
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

func (db *DB) UpdateCreatorProfile(creator *drops.Creator) error {
	return db.database.Save(creator).Error
}
