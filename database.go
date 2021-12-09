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

type Metrics struct {
	Name      string  `gorm:"size:64" json:"__name__"`
	Instance  string  `gorm:"size:64" json:"instance"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

func initDB() error {
	if database, err := gorm.Open(sqlite.Open("db/mesh.db"), &gorm.Config{}); err != nil {
		return err
	} else {
		DB = database
	}

	if err := DB.AutoMigrate(&Settings{}, &Metrics{}); err != nil {
		return err
	}

	log.Info("DB Initialized")
	return nil
}
