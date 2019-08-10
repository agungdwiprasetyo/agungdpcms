package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// LoadDatabaseConnection open database connection
func LoadDatabaseConnection() *gorm.DB {
	// conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DB"))
	conn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		"disable")

	log.Println("Connecting to database...")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	isDebugMode, _ := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	db.LogMode(isDebugMode)

	log.Println("Success connect to database")

	return db
}
