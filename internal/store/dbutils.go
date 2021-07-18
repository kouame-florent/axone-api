package store

import (
	"log"

	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		zap.L().Error("failed to open db", zap.Any("error", err))
		return &gorm.DB{}, err
	}

	return db, nil
}

func CreateSchema(db *gorm.DB) error {
	log.Println("migrating db ...")
	err := db.AutoMigrate(&axone.User{}, &axone.Role{}, &axone.Requester{}, &axone.Agent{},
		&axone.Administrator{}, &axone.Ticket{}, &axone.Tag{}, &axone.Attachment{},
		&axone.Comment{}, &axone.Assignment{}, &axone.Knowledge{})
	return err
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}

func NewDB() (*gorm.DB, error) {
	dsn := config.DataSourceName()
	return OpenDB(dsn)
}

/*
func buildSqlSchema(db *gorm.DB) {
	err := store.CreateSchema(db)
	if err != nil {
		log.Fatal(err)
	}
}
*/
