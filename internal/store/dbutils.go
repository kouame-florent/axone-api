package store

import (
	"log"

	axo "github.com/kouame-florent/axone-api/internal/axone"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("failed to open db", zap.Any("error", err))
		//log.Fatal(err)
	}

	return db
}

func CreateSchema(db *gorm.DB) error {
	log.Println("migrating db ...")
	err := db.AutoMigrate(&axo.Organization{}, &axo.User{}, &axo.Role{}, &axo.EndUser{}, &axo.Agent{},
		&axo.Administrator{}, &axo.Ticket{}, &axo.Tag{}, &axo.Attachment{},
		&axo.Comment{}, &axo.Assignment{}, &axo.Knowledge{})
	return err
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
