package acronym

import (
	"strings"
)

const testVersion = 2

func Abbreviate(longName string) string {
	//if strings.Contains(longName, "-") {
		alteredString := strings.Replace(longName, "-", " ", -1)
	//}

	splicedName := strings.Split(alteredString, " ")

	abbreviation := ""

	for _, v := range splicedName {
		abbreviation += v[:1]
	}

	return strings.ToUpper(abbreviation)
}
