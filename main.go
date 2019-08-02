package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"

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
	interrupted := make(chan os.Signal, 1)
	signal.Notify(interrupted, os.Interrupt)

	conf := config.Init()
	s := newService(conf)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.ServeHTTP()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.websocket.server.ListenAndServe()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-interrupted:
				conf.Release()
				os.Exit(0)
			}
		}
	}()

	wg.Wait()
}
