package validations

import (
	"os"
	"fmt"
	"net/url"
	"strconv"
	"math"
	"unicode/utf8"
)


func CheckError(err error)  {
	if err !=  nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func IsUTF8(str string) bool {
	valid := utf8.Valid([]byte(str))
	return valid
}

func ValidateURL(stringURL string) bool {
	_, err := url.ParseRequestURI(stringURL)
	if err != nil {
		return false
	}
	return true
}

func InBetween(num float64, min float64, max float64) bool {
	if (num >= min) && (num <= max) {
		return true
	}
	return false
}

func IsNumeric(num string) bool {
	_, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return false
	}
	return true
}

func IsPositive(num float64) bool {
	inValid := math.Signbit(num)
	return inValid
}

func ValidateStars(stringStars string) bool {
	if !IsNumeric(stringStars) {
		return false
	}

	numberStars, _ := strconv.ParseFloat(stringStars, 64)
	if !IsPositive(numberStars) && InBetween(numberStars, 0, 5) {
		return true
	}

	return false
}

func Validate(stringName string, stringURL string, numberStars string) bool{
	valid := IsUTF8(stringName)
	validURL := ValidateURL(stringURL)
	validStars := ValidateStars(numberStars)

	if !valid {
		return false
	} else if !validURL {
		return false
	} else if !validStars {
		return false
	}
	return true
}
