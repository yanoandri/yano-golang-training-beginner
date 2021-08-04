package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB interface {
	SetupDB()
}

var DB *gorm.DB

const (
	DBUser     = "test"
	DBPassword = "test"
	DBName     = "payment"
	DBHost     = "postgres"
	DBPort     = "5432"
)

func GetPostgresConnectionString(user string, password string, name string, host string, port string) string {
	database := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		name,
		password)

	return database
}

func SetupDB() {
	var err error
	conString := GetPostgresConnectionString(DBUser, DBPassword, DBName, DBHost, DBPort)

	log.Print(conString)

	DB, err = gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
}

func GetDBInstance() *gorm.DB {
	return DB
}
