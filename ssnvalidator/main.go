package main

import (
	"errors"
	"log"
	"strings"
	"strconv"
	"fmt"
)

var (
	errInvalidSSNLength = errors.New("invalid SSn Length")
	errInvalidSSNNumbers = errors.New("SSN with Non numeric digit")
	errInvalidSSNPrefix = errors.New("SSn with three zeros as prefix")
	errInvalidSSNDigitPlace = errors.New("SSN that starts with 9 requires 7 or 9 in the fourth place")
)
func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	validateSSN:= []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321","123-45-zzzz"}

	log.Printf("Checking Data : %#v", validateSSN)

	for idx, ssn := range validateSSN {
		log.Printf("validating Data.... %#v %d of %d",ssn, idx+1, len(validateSSN))

		ssn = strings.Replace(ssn, "-", "", -1)
		// calling each function and logging the errors from each
		err := validateLength(ssn)
		if err != nil {
			log.Print(err)
		}
		err = isValidNumbers(ssn)
		if err != nil {
			log.Print(err)
		}
		err = isValidPrefix(ssn)
		if err != nil {
			log.Print(err)
		}
		err = isValidDigitPlace(ssn)
		if err != nil {
			log.Print(err)
		}

	}
}

//check if the length of the SSN is not 9 and return error statement
func validateLength(ssn string)error {
	ssn = strings.TrimSpace(ssn)
	if len(ssn) != 9 {
		return fmt.Errorf("The value of %s caused an error: %v\n", ssn, errInvalidSSNLength)
	}
	return nil
}

// check if the SSN all contains numbers
func isValidNumbers(ssn string)error {
	_, err := strconv.Atoi(ssn)
	if err != nil {
		return fmt.Errorf("The value of %s caused an error: %v\n", ssn, errInvalidSSNNumbers)
	}
	return nil
}

//Check if each of the SSN starts with 000
func isValidPrefix(ssn string)error {
	if strings.HasPrefix(ssn, "000"){
		return fmt.Errorf("The value of %s caused an error: %v\n", ssn, errInvalidSSNPrefix)
	}
	return nil
}

// Check if SSN starts with a 9, then it should have a 7 or 9 in the fourth place
func isValidDigitPlace(ssn string) error {
	if string(ssn[0]) == "9" && (string(ssn[3]) != "7" && string(ssn[3]) != "9") {
		return fmt.Errorf("The value of %s caused an error: %v\n", ssn, errInvalidSSNDigitPlace)
	}
	return nil
}

