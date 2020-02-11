package validations

import (
	"os"
	"fmt"
	"unicode"
	"regexp"
)


func CheckError(err error)  {
	if err !=  nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func IsLetter(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func HasSymbol(str string) bool {
	for _, letter := range str {
		if unicode.IsSymbol(letter) {
			return true
		}
	}
	return false
}

func StringMatch(str string) bool {
	matched, _ := regexp.MatchString("[^a-zA-Z0-9_' p{L}]+", str)
	return matched
}

func ValidateString(str string) bool{
	matched := StringMatch(str)
	if !matched {
		return false
	}
	return true
}
