package validation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/agungdwiprasetyo/go-utils"
	"github.com/xeipuuv/gojsonschema"
)

// Validator for create resume json schema
type Validator struct {
	schema string
}

// New validator constructor
func New() *Validator {
	s, err := ioutil.ReadFile(fmt.Sprintf("%s/src/resume/validation/resume.json", os.Getenv("APP_PATH")))
	if err != nil {
		panic(err)
	}

	return &Validator{
		schema: string(s),
	}
}

// Validate create resume input
func (v *Validator) Validate(input interface{}) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	document := gojsonschema.NewStringLoader(string(b))
	schema := gojsonschema.NewStringLoader(v.schema)
	result, err := gojsonschema.Validate(schema, document)
	if err != nil {
		return err
	}

	multiError := utils.NewMultiError()
	if !result.Valid() {
		for i, desc := range result.Errors() {
			multiError.Append(fmt.Sprint(i), fmt.Errorf("%v", desc))
		}
	}

	if multiError.HasError() {
		return multiError
	}

	return nil
}
