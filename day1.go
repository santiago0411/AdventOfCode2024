package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLists() ([]int, []int) {
	file, err := os.Open("files/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) != 2 {
			log.Fatal("Expected 2 numbers per line")
		}

		for i := 0; i < len(numbers); i++ {
			num, err := strconv.Atoi(numbers[i])
			if err != nil {
				log.Fatalf("Expected %s to be a number", numbers[i])
			}
			if i == 0 {
				list1 = append(list1, num)
			} else {
				list2 = append(list2, num)
			}
		}
	}

	return list1, list2
}

func day1Part1() {
	var list1, list2 = readLists()

	sort.Ints(list1)
	sort.Ints(list2)

	distance := 0
	for i := 0; i < len(list1); i++ {
		distance += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Printf("[DAY 1 - Part 1] - Distance is %d\n", distance)
}

func day1Part2() {
	list1, list2 := readLists()
	frequencies := make(map[int]int)

	for _, num1 := range list1 {
		frequencies[num1] = 0
		for _, num2 := range list2 {
			if num1 == num2 {
				frequencies[num1] += 1
			}
		}
	}

	similarity := 0
	for num, freq := range frequencies {
		similarity += num * freq
	}

	fmt.Printf("[DAY 1 - Part 2] - Similarity is %d\n", similarity)
}

func Day1() {
	day1Part1()
	day1Part2()
}
