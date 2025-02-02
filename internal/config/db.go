package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	Dbconfig := LoadDBConfig()

	dsn := "user=" + Dbconfig.DBUser +
		" password=" + Dbconfig.DBPassword +
		" host=" + Dbconfig.DBHost +
		" dbname=" + Dbconfig.DBName +
		" port=" + Dbconfig.DBPort +
		" sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
}
