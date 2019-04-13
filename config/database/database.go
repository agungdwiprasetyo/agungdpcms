package database

import (
	"fmt"
	"log"
	"os"

	rd "github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbWrite, dbRead *gorm.DB
)

// LoadDatabaseConnection open database connection
func LoadDatabaseConnection() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DB"))

	log.Println("Connecting to database...")
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.AutoMigrate(&rd.Resume{})
	db.AutoMigrate(&rd.Achievement{})
	db.AutoMigrate(&rd.Education{})
	db.AutoMigrate(&rd.Experience{})
	db.AutoMigrate(&rd.Skill{})

	log.Println("Success connect to database")

	return db
}
