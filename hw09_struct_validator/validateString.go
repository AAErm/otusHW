package hw09structvalidator

import (
	"regexp"
	"strconv"
	"strings"
)

func validateString(validator, name string, value string) error {
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
		case "in":
			if err := validateInStr(valRule, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: name,
					Err:   err,
				})
			}
		case "len":
			if err := validateLen(valRule, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: name,
					Err:   err,
				})
			}
		case "regexp":
			if err := validateRegexp(valRule, value); err != nil {
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

func validateLen(valRule, value string) error {
	valLen, err := strconv.Atoi(valRule)
	if err != nil {
		return ErrInvalidRule
	}
	if len(value) != valLen {
		return ErrInvalidValue
	}
	return nil
}

func validateRegexp(valRule, value string) error {
	reg, err := regexp.Compile(valRule)
	if err != nil {
		return ErrInvalidRule
	}
	if !reg.MatchString(value) {
		return ErrInvalidValue
	}
	return nil
}

func validateInStr(valRule, value string) error {
	valIn := strings.Split(valRule, ",")
	if !inStr(valIn, value) {
		return ErrInvalidValue
	}
	return nil
}

func inStr(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func keyValueRule(rule string) (string, string, error) {
	ruleParts := strings.Split(rule, ":")
	if len(ruleParts) <= 1 {
		return "", "", ErrInvalidRule
	}
	keyRule := strings.Split(rule, ":")[0]
	valRule := strings.Split(rule, ":")[1]
	if len(ruleParts) > 2 {
		valRule = strings.Join(strings.Split(rule, ":")[1:], "")
	}
	return keyRule, valRule, nil
}
