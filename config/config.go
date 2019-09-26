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
}

// Environment model
type Environment struct {
	HTTPPort      int
	GRPCPort      int
	CORSWhitelist string
	Username      string
	Password      string
	TokenAge      time.Duration
	Websocket     bool
}

// GlobalEnv global env
var GlobalEnv Environment

// Init global config
func Init() *Config {
	var err error

	conf := new(Config)
	conf.DB = database.LoadDatabaseConnection()

	conf.PrivateKey = key.LoadPrivateKey()
	conf.PublicKey = key.LoadPublicKey()

	GlobalEnv.HTTPPort, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
	GlobalEnv.GRPCPort, err = strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		panic(err)
	}

	GlobalEnv.Username = os.Getenv("USERNAME")
	GlobalEnv.Password = os.Getenv("PASSWORD")

	if v, ok := os.LookupEnv("CORS_WHITELIST"); ok {
		GlobalEnv.CORSWhitelist = v
	} else {
		GlobalEnv.CORSWhitelist = "*"
	}

	GlobalEnv.TokenAge, err = time.ParseDuration(os.Getenv("TOKEN_AGE"))
	if err != nil {
		panic(err)
	}

	GlobalEnv.Websocket, _ = strconv.ParseBool(os.Getenv("WEBSOCKET"))

	return conf
}

// Release all config variable in memory when application is exited
func (c *Config) Release() {
	c.DB.Close()
	log.Println("database is closed")
}
