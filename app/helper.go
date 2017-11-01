package app

import "strconv"

func StringToInt(value string) (number int) {
	newNumber, _ := strconv.Atoi(value)

	return newNumber
}
