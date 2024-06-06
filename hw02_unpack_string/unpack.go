package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	result := []rune{}
	for k, r := range str {
		i, err := strconv.Atoi(string(r))
		if err != nil {
			result = append(result, r)
			continue
		}

		if k == 0 || nextIsInt(str, k) {
			return "", ErrInvalidString
		}

		if prevIsRepeated(str, k) {
			strToAppend := strings.Repeat(string(str[k-1]), i)
			result = append(result[0:len(result)-1], []rune(strToAppend)...)
			continue
		}

		result = append(result, r)
	}
	return string(result), nil
}

func nextIsInt(s string, i int) bool {
	nextKey := i + 1
	if len(s)-1 < nextKey {
		return false
	}

	if _, err := strconv.Atoi(string(s[nextKey])); err != nil {
		return false
	}

	return true
}

func prevIsRepeated(s string, i int) bool {
	if i == 0 {
		return false
	}

	if _, err := strconv.Atoi(string(s[i-1])); err != nil {
		return true
	}

	return false
}
