package database

import (
	"log"
	"os"

	"github.com/sixfwa/fiber-gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database \n", err.Error())
		os.Exit(2)

	}
	log.Println("connected to the database sucessfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("runnng migratiion")
	//todo
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}); err != nil {
		log.Fatalf("failed to run migration: %v", err)
	}

	Database = DbInstance{
		Db: db,
	}
}