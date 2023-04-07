package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ms-product/model"
)

var DB *gorm.DB

func RunMigrations() {
	if err := DB.AutoMigrate(&model.Product{}); err != nil {
		panic(err)
	}
}

func DbConnection() {
	dsn := "host=localhost user=admin password=admin dbname=product port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = conn
}
