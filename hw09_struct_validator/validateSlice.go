package hw09structvalidator

import (
	"fmt"
	"reflect"
)

func validateSlice(validator, name string, value reflect.Value) error {
	elem := value.Type().Elem()
	switch elem.Kind() {
	case reflect.Int:
		values, ok := value.Interface().([]int64)
		if !ok {
			return ErrNotSupportedType
		}

		return validateSliceInt(validator, name, values)
	case reflect.String:
		values, ok := value.Interface().([]string)
		if !ok {
			return ErrNotSupportedType
		}

		return validateSliceString(validator, name, values)
	}

	return ErrNotSupportedType
}

func validateSliceInt(validator, name string, values []int64) error {
	validationErrors := make(ValidationErrors, 0)
	for k, v := range values {
		if err := validateInt(validator, name, v); err != nil {
			validationErrors = append(validationErrors, ValidationError{
				Field: fmt.Sprintf("%s elem %d", name, k),
				Err:   err,
			})
		}
	}

	return validationErrors
}

func validateSliceString(validator, name string, values []string) error {
	validationErrors := make(ValidationErrors, 0)
	for k, v := range values {
		if err := validateString(validator, name, v); err != nil {
			validationErrors = append(validationErrors, ValidationError{
				Field: fmt.Sprintf("%s elem %d", name, k),
				Err:   err,
			})
		}
	}

	return validationErrors
}
