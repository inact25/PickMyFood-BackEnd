package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ValidateInputNotNil(data ...interface{}) error {
	for _, value := range data {
		switch result := value.(type) {
		case string:
			if len(result) == 0 {
				return errors.New("Data Input Cannot Empty")
			}
		case int:
			if result == 0 {
				return errors.New("Data Input Cannot 0")
			}
		}
	}
	return nil
}

func ValidateInputNumber(data interface{}) error {
	switch result := data.(type) {
	case string:
		if _, err := strconv.Atoi(result); err != nil {
			return errors.New("ID Cannot Contain Characters")
		}
	default:
		return errors.New("ID input not an INT data type")
	}
	return nil
}

func IsStatusValid(status string) bool {
	validation := strings.ToLower(status)
	if validation != "i" && validation != "a" {
		return false
	}
	return true
}

func CheckEmpty(data ...interface{}) error {
	for _, value := range data {
		switch value {
		case "":
			return errors.New("make sure input not empty")
		case 0:
			return errors.New("make sure input not a zero")
		case nil:
			return errors.New("make sure input not a nil")
		}
	}
	return nil
}

func CheckInt(data ...string) error {
	for _, value := range data {
		_, err := strconv.Atoi(value)
		if err != nil {
			return errors.New(fmt.Sprintf("%v Is not valid", value))
		}
	}
	return nil
}
