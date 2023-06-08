package utils

import (
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"inverse.so/models"
)

var DB *gorm.DB

func initialiseDB(dsn string) {
	var err error

	log.Debug().Msgf("Initialising Database...")

	isProd, _ := IsProduction()
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

	DB, err = gorm.Open(postgres.Open(dsn), ormConfig)

	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, err := DB.DB()
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
}

func SetupDB(dsn string) *gorm.DB {
	initialiseDB(dsn)

	//User managment
	DB.AutoMigrate(
		&models.Creator{},
		&models.Collection{},
		&models.Item{},
		&models.SingleEmailClaim{},
		&models.EmailDomainWhiteList{},
		&models.EmailOTP{},
		&models.TwitterCriteria{},
		&models.TelegramCriteria{},
	)

	return DB
}

func GetTableName(model interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	if err := stmt.Parse(&model); err != nil {
		log.Debug().Err(err).Msg("Error parsing model to get table name ")
		return ""
	}

	return stmt.Schema.Table
}

func GetRelationalName(model interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	if err := stmt.Parse(&model); err != nil {
		log.Debug().Err(err).Msg("Error parsing model to get table name")
		return ""
	}

	return stmt.Schema.Name
}
