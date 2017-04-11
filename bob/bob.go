package bob // package name must match the package name in bob_test.go

import (
	"strings"
	"regexp"
	"fmt"
)

const testVersion = 2 // same as targetTestVersion
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func Hey(input string) string {
	matched, _ := regexp.MatchString("([a-zA-Z!]+)", input)

	if allSpacey(input) {
		return "Fine. Be that way!"
	} else if input == strings.ToUpper(input) && matched {
		return "Whoa, chill out!"
	} else if isQuestionable(input) {
		return "Sure."
	}
	
	return "Whatever."
}

func allSpacey(input string) bool {
	chars := strings.Split(input, "")
	fmt.Println(chars)

	if len(chars) == 0 {
		return true
	}

	for i := 0; i < len(input); i++ {
		if (chars[i] != " ") && (chars[i] != "\t") {
			return false
		}
	}

	return true
}

func isQuestionable(input string) bool {
	noSpaces := strings.Replace(input, " ", "", -1)

	if noSpaces[len(noSpaces)-1:] == "?" {
		return true
	}

	return false
}