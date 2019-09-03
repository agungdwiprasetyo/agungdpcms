package jsonschema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/agungdwiprasetyo/go-utils"
	"github.com/xeipuuv/gojsonschema"
)

var jsonSchemaList = map[string]*gojsonschema.Schema{}

// Load all schema
func Load(path string) error {
	gojsonschema.Locale = &customMessage{}
	return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		if strings.HasSuffix(fileName, ".json") {
			s, err := ioutil.ReadFile(p)
			if err != nil {
				return err
			}

			var data map[string]interface{}
			err = json.Unmarshal(s, &data)
			if err != nil {
				return err
			}
			id, ok := data["id"].(string)
			if !ok {
				id = strings.TrimSuffix(strings.TrimPrefix(p, path), ".json")
			}
			jsonSchemaList[id], err = gojsonschema.NewSchema(gojsonschema.NewBytesLoader(s))
		}
		return nil
	})
}

// Get json schema by ID
func Get(schemaID string) (*gojsonschema.Schema, error) {
	schema, ok := jsonSchemaList[schemaID]
	if !ok {
		return nil, fmt.Errorf("schema '%s' not found", schemaID)
	}

	return schema, nil
}

// Validate from Go data type
func Validate(schemaID string, input interface{}) error {
	multiError := utils.NewMultiError()

	schema, err := Get(schemaID)
	if err != nil {
		multiError.Append("getSchema", err)
		return multiError
	}

	document := gojsonschema.NewGoLoader(input)
	return validate(schema, document)
}

// ValidateDocument document
func ValidateDocument(schemaID string, jsonByte []byte) error {
	multiError := utils.NewMultiError()

	schema, err := Get(schemaID)
	if err != nil {
		multiError.Append("getSchema", err)
		return multiError
	}

	document := gojsonschema.NewBytesLoader(jsonByte)
	return validate(schema, document)
}

func validate(schema *gojsonschema.Schema, document gojsonschema.JSONLoader) error {
	multiError := utils.NewMultiError()

	result, err := schema.Validate(document)
	if err != nil {
		multiError.Append("validateInput", err)
		return multiError
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
