package validation

import (
	"fmt"
	"regexp"
)

func EmailAdress(input string) (bool, error) {
	match, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, input)
	if err != nil || match != true {
		return false, fmt.Errorf("The email address does not match the required format.")
	}
	return match, nil
}

func PhoneNumber(input string) (bool, error) {
	match, err := regexp.MatchString(`^\+[1-9][0-9]{1,14}$`, input)
	if err != nil || match != true {
		return false, fmt.Errorf("The phone number does not match the required format.")
	}
	return match, nil
}
