package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ParseAndValidate(data io.Reader, v any) error {
	decoder := json.NewDecoder(data)
	_ = decoder.Decode(v)
	validate := validator.New()
	err := validate.Struct(v)
	if err != nil {
		return NewValidationError(v, err.(validator.ValidationErrors))
	}
	return nil
}

func NewValidationError(v interface{}, err validator.ValidationErrors) error {
	msg := ""
	for _, err := range err {
		field := getJSONTag(v, err.Field())
		msg += fmt.Sprintf("'%s' %s \n", field, err.Tag())
	}
	return errors.New(msg)
}

func getJSONTag(v interface{}, fieldName string) string {
	val := reflect.ValueOf(v)

	// Check if the value is a struct
	if val.Kind() != reflect.Struct {
		return ""
	}

	typ := val.Type()

	// Loop through the struct fields to find the matching field
	for i := 0; i < val.NumField(); i++ {
		if typ.Field(i).Name == fieldName {
			// Get the JSON tag for the matching field
			return typ.Field(i).Tag.Get("json")
		}
	}

	return ""
}
