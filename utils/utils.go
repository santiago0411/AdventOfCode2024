package utils

import (
	"io"
	"log"
	"os"
	"strconv"
)

func ParseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Expected %s to be a number", str)
	}
	return num
}

func ReadFileToString(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}
