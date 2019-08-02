package config

import (
	"crypto/rsa"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/agungdwiprasetyo/agungdpcms/config/database"
	"github.com/agungdwiprasetyo/agungdpcms/config/key"
	"github.com/jinzhu/gorm"
)

// Config abstraction
type Config struct {
	DB *gorm.DB

	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey

	Env Environment
}

// Environment model
type Environment struct {
	HTTPPort      int
	CORSWhitelist string
	Username      string
	Password      string
	TokenAge      time.Duration
	Websocket     bool
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

	conf.Env.TokenAge, err = time.ParseDuration(os.Getenv("TOKEN_AGE"))
	if err != nil {
		panic(err)
	}

	conf.Env.Websocket, _ = strconv.ParseBool(os.Getenv("WEBSOCKET"))

	return conf
}

// Release all config variable in memory when application is exited
func (c *Config) Release() {
	c.DB.Close()
	log.Println("database is closed")
}
