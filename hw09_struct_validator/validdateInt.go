package hw09structvalidator

import (
	"strconv"
	"strings"
)

func validateInt(validator, name string, value int64) error {
	rules := strings.Split(validator, "|")
	validationErrors := make(ValidationErrors, 0)
	for _, rule := range rules {
		keyRule, valRule, err := keyValueRule(rule)
		if err != nil {
			validationErrors = append(validationErrors, ValidationError{
				Field: name,
				Err:   err,
			})
			continue
		}
		switch keyRule {
		case "min":
			if err := validateMin(valRule, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: name,
					Err:   err,
				})
			}
		case "max":
			if err := validateMax(valRule, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: name,
					Err:   err,
				})
			}
		case "in":
			if err := validateInInt(valRule, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: name,
					Err:   err,
				})
			}
		default:
			validationErrors = append(validationErrors, ValidationError{
				Field: name,
				Err:   ErrInvalidRule,
			})

		}
	}

	if len(validationErrors) != 0 {
		return validationErrors
	}

	return nil
}

func validateMin(valRule string, value int64) error {
	intValRul, err := strconv.ParseInt(valRule, 10, 0)
	if err != nil {
		return ErrInvalidRule
	}
	if intValRul >= value {
		return ErrInvalidValue
	}
	return nil
}

func validateMax(valRule string, value int64) error {
	intValRul, err := strconv.ParseInt(valRule, 10, 0)
	if err != nil {
		return ErrInvalidRule
	}
	if intValRul <= value {
		return ErrInvalidValue
	}
	return nil
}

func validateInInt(valRule string, value int64) error {
	valIn := strings.Split(valRule, ",")
	intValIn, err := intValIn(valIn)
	if err != nil {
		return ErrInvalidRule
	}
	if !inInt(intValIn, value) {
		return ErrInvalidValue
	}
	return nil
}

func intValIn(strs []string) ([]int64, error) {
	res := make([]int64, 0)
	for _, v := range strs {
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return nil, ErrInvalidRule
		}
		res = append(res, val)
	}
	return res, nil
}

func inInt(arr []int64, str int64) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
