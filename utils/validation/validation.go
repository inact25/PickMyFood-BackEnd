package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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
