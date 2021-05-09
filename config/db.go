package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	DBUser     = "postgres"
	DBPassword = ""
	DBName     = "payment"
	DBHost     = "127.0.0.1"
	DBPort     = "5432"
)

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := GetPostgresConnectionString()

	log.Print(conString)

	DB, err = gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
