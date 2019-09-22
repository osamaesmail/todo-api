package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"todo-api/models"
	"log"
	"os"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Init() {
	user := getEnv("PG_USER", "admin")
	password := getEnv("PG_PASSWORD", "admin")
	host := getEnv("PG_HOST", "127.0.0.53")
	port := getEnv("PG_PORT", "5433")
	database := getEnv("PG_DB", "todos")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	if !db.HasTable(&models.Todo{}) {
		err := db.CreateTable(&models.Todo{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&models.Todo{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}