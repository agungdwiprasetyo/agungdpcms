package schema

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func LoadSchema() string {
	s, err := ioutil.ReadFile(fmt.Sprintf("%s/schema/schema.graphql", os.Getenv("APP_PATH")))
	if err != nil {
		log.Fatal(err)
	}

	return string(s)
}
