package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"reflect"
)

var (
	ErrValueIsNotStruct = errors.New("value is not a struct")
	ErrValueIsInterface = errors.New("value is interface")
	ErrNotSupportedType = errors.New("type is not supported")
	ErrInvalidRule      = errors.New("rule is invalid")
	ErrInvalidValue     = errors.New("value is invalid")
)

type ValidationError struct {
	Field string `json:"field"`
	Err   error  `json:"error"`
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	if len(v) == 0 {
		return ""
	}

	bb, _ := json.Marshal(v)
	return string(bb)
}

func Validate(v interface{}) error {
	reflectValue := reflect.ValueOf(v)
	if reflectValue.Kind() != reflect.Struct {
		return ErrValueIsNotStruct
	}

	if reflectValue.Kind() == reflect.Interface {
		return ErrValueIsInterface
	}

	validationErrors := make(ValidationErrors, 0)
	t := reflectValue.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		validator := field.Tag.Get("validate")
		if validator == "" {
			continue
		}
		fieldValue := reflectValue.Field(i)

		switch fieldValue.Kind() {
		case reflect.Int:
			if err := validateInt(validator, field.Name, fieldValue.Int()); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: field.Name,
					Err:   err,
				})
			}
		case reflect.String:
			if err := validateString(validator, field.Name, fieldValue.String()); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: field.Name,
					Err:   err,
				})
			}
		case reflect.Slice:
			if err := validateSlice(validator, field.Name, fieldValue); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: field.Name,
					Err:   err,
				})
			}
		default:
			return ErrNotSupportedType
		}
	}

	if len(validationErrors) != 0 {
		return validationErrors
	}

	return nil
}
