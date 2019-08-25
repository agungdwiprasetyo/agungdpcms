package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	mysqlConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DBMS_USER"), os.Getenv("DBMS_PASS"), os.Getenv("DBMS_HOST"), os.Getenv("DBMS_PORT"), os.Getenv("DBMS_DBNAME"))
	postgresConn = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DBMS_USER"), os.Getenv("DBMS_PASS"), os.Getenv("DBMS_HOST"), os.Getenv("DBMS_PORT"), os.Getenv("DBMS_DBNAME"))
)

// LoadDatabaseConnection open database connection
func LoadDatabaseConnection() *gorm.DB {
	log.Println("Connecting to database...")
	db, err := gorm.Open("postgres", postgresConn)
	if err != nil {
		panic(err)
	}

	isDebugMode, _ := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	db.LogMode(isDebugMode)

	log.Println("Success connect to database")
	return db
}
