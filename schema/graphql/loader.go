package graphql

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// LoadSchema graphql from file
func LoadSchema() string {
	var schema strings.Builder
	here := fmt.Sprintf("%s/schema/graphql/", os.Getenv("APP_PATH"))

	// load main schema
	s, err := ioutil.ReadFile(here + "schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	schema.Write(s)

	// load resume schema
	s, err = ioutil.ReadFile(here + "resume.graphql")
	if err != nil {
		log.Fatal(err)
	}
	schema.Write(s)

	// load user schema
	s, err = ioutil.ReadFile(here + "user.graphql")
	if err != nil {
		log.Fatal(err)
	}
	schema.Write(s)

	// load chat schema
	s, err = ioutil.ReadFile(here + "chat.graphql")
	if err != nil {
		log.Fatal(err)
	}
	schema.Write(s)

	return schema.String()
}
