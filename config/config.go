package config

import (
	"log"
	"os"
	"strconv"

	"github.com/agungdwiprasetyo/agungdpcms/config/database"
	"github.com/jinzhu/gorm"
)

// Config abstraction
type Config struct {
	DB  *gorm.DB
	Env struct {
		HTTPPort int
		Username string
		Password string
	}
}

// Init global config
func Init() *Config {
	var err error

	conf := new(Config)
	conf.DB = database.LoadDatabaseConnection()

	conf.Env.HTTPPort, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
	conf.Env.Username = os.Getenv("USERNAME")
	conf.Env.Password = os.Getenv("PASSWORD")

	return conf
}

// Release all config variable in memory when application is exited
func (c *Config) Release() {
	c.DB.Close()
	log.Println("database is closed")
}
