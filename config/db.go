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

func GetPostgresConnectionString(user string, password string, name string, host string, port string, urlMode bool) string {
	var database string
	if urlMode {
		database = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			user,
			password,
			host,
			port,
			name)
	} else {
		database = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			host,
			port,
			user,
			name,
			password)
	}

	return database
}

func SetupDB() {
	var err error
	conString := GetPostgresConnectionString(DBUser, DBPassword, DBName, DBHost, DBPort, false)

	log.Print(conString)

	DB, err = gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
}

func GetDBInstance() *gorm.DB {
	return DB
}
