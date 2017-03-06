package acronym

import (
	"strings"
	"unicode"
	"fmt"
)

const testVersion = 2

func Abbreviate(longName string) string {
	alteredString := strings.Replace(longName, "-", " ", -1)

	altertedStringToRunes := []rune(alteredString)

	fmt.Println(altertedStringToRunes)

	for i, v := range altertedStringToRunes[1:] {
		if unicode.IsUpper(v) {
			if !unicode.IsUpper(altertedStringToRunes[i]) && altertedStringToRunes[i] != 32 {
				altertedStringToRunes[i] = ' '
			}
		}
	}

	newString := string(altertedStringToRunes)

	splicedName := strings.Split(newString, " ")

	abbreviation := ""

	for _, v := range splicedName {
		abbreviation += v[:1]
	}

	return strings.ToUpper(abbreviation)
}
