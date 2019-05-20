package validation

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/agungdwiprasetyo/agungdpcms/shared/validator"
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
func New() validator.Validator {
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
func (v *Validator) Validate(input interface{}) (multiError *utils.MultiError) {
	document := gojsonschema.NewGoLoader(input)

	// take value input if type is pointer
	refValue := reflect.ValueOf(input)
	if refValue.Kind() == reflect.Ptr {
		refValue = refValue.Elem()
	}
	input = refValue.Interface()

	var result *gojsonschema.Result
	var err error
	multiError = utils.NewMultiError()

	switch input.(type) {
	case domain.Resume:
		result, err = v.resume.Validate(document)
	case filter.Filter:
		result, err = v.filter.Validate(document)
	default:
		err = fmt.Errorf("unknown input type of %T", input)
	}
	if err != nil {
		multiError.Append("validateInput", err)
		return
	}

	if !result.Valid() {
		for _, desc := range result.Errors() {
			multiError.Append(desc.Field(), fmt.Errorf("value '%v' %v", desc.Value(), desc.Description()))
		}
	}

	if multiError.HasError() {
		return multiError
	}

	return nil
}
