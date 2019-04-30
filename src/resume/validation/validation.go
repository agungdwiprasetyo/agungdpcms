package validation

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/agungdwiprasetyo/go-utils"
	"github.com/xeipuuv/gojsonschema"
)

// Validator for create resume json schema
type Validator struct {
	resume *gojsonschema.Schema
	filter *gojsonschema.Schema
}

// New validator constructor, only once to read file *.json when init application
func New() *Validator {
	s, err := ioutil.ReadFile(fmt.Sprintf("%s/src/resume/validation/resume.json", os.Getenv("APP_PATH")))
	if err != nil {
		panic(err)
	}
	resumeSchema, err := gojsonschema.NewSchema(gojsonschema.NewStringLoader(string(s)))
	if err != nil {
		panic(err)
	}

	filterLoader := gojsonschema.NewSchemaLoader()
	filterSchema, err := filterLoader.Compile(filter.LoadSchema([]string{"slug"}))
	if err != nil {
		panic(err)
	}

	return &Validator{
		resume: resumeSchema,
		filter: filterSchema,
	}
}

// Validate create resume input
func (v *Validator) Validate(input interface{}) (err error) {
	document := gojsonschema.NewGoLoader(input)

	// take value input if type is pointer
	refValue := reflect.ValueOf(input)
	if refValue.Kind() == reflect.Ptr {
		refValue = refValue.Elem()
	}
	input = refValue.Interface()

	var result *gojsonschema.Result
	switch input.(type) {
	case domain.Resume:
		result, err = v.resume.Validate(document)
	case filter.Filter:
		result, err = v.filter.Validate(document)
	default:
		return fmt.Errorf("unknown input type")
	}
	if err != nil {
		return err
	}

	return parseError(result)
}

func parseError(result *gojsonschema.Result) error {
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
