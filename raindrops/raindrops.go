package raindrops

import (
	"strconv"
)

const testVersion = 2

func Convert(number int) string {
	returnValue := ""

	if number % 3 == 0 {
		returnValue += "Pling"
	} 
	
	if number % 5 == 0 {
		returnValue += "Plang"
	}
	
	if number % 7 == 0 {
		returnValue += "Plong"
	} 

	if returnValue == "" {
		return strconv.Itoa(number)
	}

	return returnValue
}

// The test program has a benchmark too.  How fast does your Convert convert?
