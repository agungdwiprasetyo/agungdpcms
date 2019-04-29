package config

import (
	"crypto/rsa"
	"log"
	"os"
	"strconv"

	"github.com/agungdwiprasetyo/agungdpcms/config/database"
	"github.com/agungdwiprasetyo/agungdpcms/config/key"
	"github.com/jinzhu/gorm"
)

// Config abstraction
type Config struct {
	DB *gorm.DB

	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey

	Env struct {
		HTTPPort      int
		CORSWhitelist string
		Username      string
		Password      string
	}
}

// Init global config
func Init() *Config {
	var err error

	conf := new(Config)
	conf.DB = database.LoadDatabaseConnection()

	conf.PrivateKey = key.LoadPrivateKey()
	conf.PublicKey = key.LoadPublicKey()

	conf.Env.HTTPPort, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
	conf.Env.Username = os.Getenv("USERNAME")
	conf.Env.Password = os.Getenv("PASSWORD")

	if v, ok := os.LookupEnv("CORS_WHITELIST"); ok {
		conf.Env.CORSWhitelist = v
	} else {
		conf.Env.CORSWhitelist = "*"
	}

	return conf
}

// Release all config variable in memory when application is exited
func (c *Config) Release() {
	c.DB.Close()
	log.Println("database is closed")
}
