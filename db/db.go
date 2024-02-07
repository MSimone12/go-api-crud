package db

import (
	"go-api-crud/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable Timezone=America/Sao_Paulo"

	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&structs.User{})

}
