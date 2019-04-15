package database

import (
	"fmt"
	"log"
	"os"

	cd "github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	rd "github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	ud "github.com/agungdwiprasetyo/agungdpcms/src/user/domain"
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

	db.LogMode(true)

	db.AutoMigrate(&rd.Resume{}, &rd.Achievement{}, &rd.Education{}, &rd.Experience{}, &rd.Skill{})
	db.AutoMigrate(&cd.Group{}, &cd.Message{})
	db.AutoMigrate(&ud.User{}, &ud.Role{})

	log.Println("Success connect to database")

	return db
}
