package utils

import (
	"log"
	"strconv"
)

func ParseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Expected %s to be a number", str)
	}
	return num
}
