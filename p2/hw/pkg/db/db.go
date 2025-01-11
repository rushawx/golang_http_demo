// pkg/db/db.go
package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hw/configs"
	"log"
)

type Db struct {
	*gorm.DB
}

// Db - база данных с хранением в Postgres

func NewDb(config *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return &Db{db}
}

// NewDb создает новый экземпляр базы данных Postgres
