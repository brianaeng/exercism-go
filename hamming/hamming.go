package hamming

import (
	"strings"
	"errors"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
	counter := 0

	if len(a) == 0 || len(b) == 0 {
		return 0, nil
	} else if len(a) != len(b) {
		err := errors.New("String too short")
		return -1, err
	}

	stringA := strings.Split(a, "")
	stringB := strings.Split(b, "")

	for i := 0; i < len(a); i++ {
		if stringA[i] != stringB[i] {
			counter++
		}
	}

	return counter, nil
}
