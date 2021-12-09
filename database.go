package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Settings struct {
	Key   string `gorm:"size:20"`
	Value string `gorm:"size: 254"`
}

func initDB() error {
	if database, err := gorm.Open(sqlite.Open("db/mesh.db"), &gorm.Config{}); err != nil {
		return err
	} else {
		DB = database
	}

	if err := DB.AutoMigrate(&Settings{}); err != nil {
		return err
	}

	log.Info("DB Initialized")
	return nil
}
