package database

import (
	"assignment-2/databaseconfig"
	"assignment-2/entity"
	"log"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(databaseconfig.GetDBConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := db.AutoMigrate(&entity.Order{}, &entity.Item{}); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Database Connected")
}

func GetPostgresInstance() *gorm.DB {
	return db
}
