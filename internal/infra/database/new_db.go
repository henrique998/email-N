package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=henrique password=123456 dbname=emailn_dev port=5432 sslmode=disabled"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to database")
	}

	return db
}
