package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	host: locahost
}

var storage *gorm.DB

func Storage() *gorm.DB {
	return storage
}

func InitDB() {
	dsn := "host=localhost user=memories_user password=memories_password dbname=memories_dev port=5432 sslmode=disable"
	var err error
	storage, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	}
}
