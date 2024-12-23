package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readReports() [][]int {
	file, err := os.Open("files/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var reports [][]int
	reportIndex := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		reports = append(reports, []int{})
		for _, number := range strings.Fields(line) {
			num, err := strconv.Atoi(number)
			if err != nil {
				log.Fatalf("Expected %s to be a number", number)
			}
			reports[reportIndex] = append(reports[reportIndex], num)
		}
		reportIndex++
	}

	return reports
}

func isReportSafe(numbers []int) bool {
	if (numbers == nil) || (len(numbers) < 2) {
		log.Fatal("isReportSafe expected a list of at least 2 numbers")
	}
	asc := numbers[0] < numbers[1]
	for i := 0; i < len(numbers)-1; i++ {
		if asc {
			if numbers[i] > numbers[i+1] {
				return false
			}
		} else {
			if numbers[i] < numbers[i+1] {
				return false
			}
		}
		diff := int(math.Abs(float64(numbers[i] - numbers[i+1])))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func day2Part1() {
	reports := readReports()
	safeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}
	fmt.Printf("[DAY 2 - Part 1] - Safe reports count: %d\n", safeReports)
}

func day2Part2() {
	reports := readReports()
	safeReports := 0
	for _, report := range reports {
		for i := -1; i < len(report); i++ {
			var r []int
			if i == -1 {
				r = make([]int, len(report))
				copy(r, report)
			} else {
				r = make([]int, len(report)-1)
				copy(r, report[:i])
				copy(r[i:], report[i+1:])
			}
			if isReportSafe(r) {
				safeReports++
				break
			}
		}
	}
	fmt.Printf("[DAY 2 - Part 2] - Safe reports count: %d\n", safeReports)
}

func Day2() {
	day2Part1()
	day2Part2()
}
