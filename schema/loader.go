package schema

import (
	"io/ioutil"
	"log"
)

func LoadSchema() string {
	s, err := ioutil.ReadFile("./schema/schema.graphql")
	if err != nil {
		log.Fatal(err)
	}

	return string(s)
}
