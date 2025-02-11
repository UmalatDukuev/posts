package database

import (
	"fmt"
	"log"
	"posts/dbmodel"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBInstance *gorm.DB
var err error

var CONNECTION_STRING string = "host=localhost user=postgres password=03795 dbname=posts port=5432 sslmode=disable"

func ConnectDB() {
	DBInstance, err = gorm.Open(postgres.Open(CONNECTION_STRING), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	} else {
		fmt.Println("Database Connected successfully...")
	}
}

func CreateDB() {
	DBInstance.Exec("CREATE SCHEMA IF NOT EXISTS public")
}

func MigrateDB() {
	DBInstance.AutoMigrate(&dbmodel.Post{})
	fmt.Println("Database migration completed....")
}
