package filter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

// LoadSchema filter with custom enum sort by
func LoadSchema(enumSortBy []string) gojsonschema.JSONLoader {
	s, err := ioutil.ReadFile(fmt.Sprintf("%s/shared/filter/schema.json", os.Getenv("APP_PATH")))
	if err != nil {
		panic(err)
	}

	var tmp map[string]interface{}
	json.Unmarshal(s, &tmp)
	var dst = make(map[string]interface{}, len(tmp))

	for key, val := range tmp {
		if key == "properties" {
			vm := make(map[string]interface{})
			for k, v := range val.(map[string]interface{}) {
				if k == "sortBy" {
					v = map[string]interface{}{
						"type": "string",
						"enum": enumSortBy,
					}
				}
				vm[k] = v
			}
			val = vm
		}
		dst[key] = val
	}

	return gojsonschema.NewGoLoader(dst)
}
