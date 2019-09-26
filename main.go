package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	env "github.com/joho/godotenv"
)

func init() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("APP_PATH", appPath)

	if err := env.Load(fmt.Sprintf("%s/.env", appPath)); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conf := config.Init()
	s := newService(conf)

	// serve HTTP for graphql
	go s.ServeHTTP()

	// serve GRPC server
	go s.ServeGRPC()

	// serve websocket server
	go s.websocket.server.ListenAndServe()

	// wait os interupted or PID has been killed
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case <-quit:
		s.Shutdown()
		conf.Release()
		os.Exit(0)
	}
}
