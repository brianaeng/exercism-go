package acronym

import (
	"strings"
	//"unicode"
	"fmt"
)

const testVersion = 2

func Abbreviate(longName string) string {
	//if strings.Contains(longName, "-") {
	alteredString := strings.Replace(longName, "-", " ", -1)
	//}

	altertedStringToRunes := []rune(alteredString)

	fmt.Println(altertedStringToRunes)

	// for i, v := range altertedStringToRunes[1:] {
	// 	if unicode.IsUpper(v) {
	// 		if !unicode.IsUpper(altertedStringToRunes[i-1]) {
	// 			altertedStringToRunes = append(altertedStringToRunes, 0)
	// 			copy(altertedStringToRunes[i-1:], altertedStringToRunes[i:])
	// 			altertedStringToRunes[i] = ' '
	// 		}
	// 	}
	// }

	splicedName := strings.Split(alteredString, " ")

	abbreviation := ""

	for _, v := range splicedName {
		abbreviation += v[:1]
	}

	return strings.ToUpper(abbreviation)
}
