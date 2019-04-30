package validation

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/agungdwiprasetyo/go-utils"
	"github.com/xeipuuv/gojsonschema"
)

// Validator for create resume json schema
type Validator struct {
	resume gojsonschema.JSONLoader
}

// New validator constructor, only once to read file *.json
func New() *Validator {
	s, err := ioutil.ReadFile(fmt.Sprintf("%s/src/resume/validation/resume.json", os.Getenv("APP_PATH")))
	if err != nil {
		panic(err)
	}

	return &Validator{
		resume: gojsonschema.NewStringLoader(string(s)),
	}
}

// Validate create resume input
func (v *Validator) Validate(input interface{}) error {
	document := gojsonschema.NewGoLoader(input)
	result, err := gojsonschema.Validate(v.resume, document)
	if err != nil {
		return err
	}

	multiError := utils.NewMultiError()
	if !result.Valid() {
		for _, desc := range result.Errors() {
			multiError.Append(desc.Field(), fmt.Errorf("value '%v' %v;", desc.Value(), desc.Description()))
		}
	}

	if multiError.HasError() {
		return multiError
	}

	return nil
}
