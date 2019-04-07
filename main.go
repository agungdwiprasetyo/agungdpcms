package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("APP_PATH", appPath)

	s := NewService()
	s.ServeHTTP()
}
