package main

import (
	"AdventOfCode/utils"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func day3Part1() {
	re, err := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	if err != nil {
		log.Fatal(err)
		return
	}

	matches := re.FindAllStringSubmatch(utils.ReadFileToString("files/day3.txt"), -1)
	result := 0
	for _, m := range matches {
		num1 := utils.ParseInt(m[1])
		num2 := utils.ParseInt(m[2])
		result += num1 * num2
	}

	fmt.Printf("[DAY 3 - Part 1] - Result: %d\n", result)
}

func day3Part2() {
	re, err := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)|(do\\(\\)|don't\\(\\))")
	if err != nil {
		log.Fatal(err)
		return
	}

	matches := re.FindAllString(utils.ReadFileToString("files/day3.txt"), -1)
	result := 0
	do := true

	for _, m := range matches {
		if do && strings.HasPrefix(m, "mul") {
			comma := strings.Index(m, ",")
			num1 := utils.ParseInt(m[4:comma])
			num2 := utils.ParseInt(m[comma+1 : len(m)-1])
			result += num1 * num2
			continue
		}
		if m == "do()" {
			do = true
		} else if m == "don't()" {
			do = false
		}
	}

	fmt.Printf("[DAY 3 - Part 2] - Result: %d\n", result)
}

func Day3() {
	day3Part1()
	day3Part2()
}
