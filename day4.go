package main

import (
	"AdventOfCode/utils"
	"fmt"
	"strings"
)

func day4Part1() {
	str := utils.ReadFileToString("files/day4.txt")

	var separator string
	if strings.Contains(str, "\r\n") {
		separator = "\r\n"
	} else {
		separator = "\n"
	}

	width := strings.Index(str, separator)
	height := strings.Count(str, separator)
	str = strings.Replace(str, separator, "", -1)

	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := x + y*width
			// Horizontal forwards
			if x < width-3 && str[idx:idx+4] == "XMAS" {
				count++
			}
			// Horizontal backwards
			if x >= 3 && str[idx-3:idx+1] == "SAMX" {
				count++
			}
			// Vertical downwards
			if y < height-3 &&
				str[idx] == 'X' &&
				str[idx+height] == 'M' &&
				str[idx+height*2] == 'A' &&
				str[idx+height*3] == 'S' {
				count++
			}
			// Vertical upwards
			if y >= 3 &&
				str[idx] == 'X' &&
				str[idx-height] == 'M' &&
				str[idx-height*2] == 'A' &&
				str[idx-height*3] == 'S' {
				count++
			}
			// Diagonal down right
			if x < width-3 && y < height-3 &&
				str[idx] == 'X' &&
				str[idx+height+1] == 'M' &&
				str[idx+height*2+2] == 'A' &&
				str[idx+height*3+3] == 'S' {
				count++
			}
			// Diagonal down left
			if x >= 3 && y < height-3 &&
				str[idx] == 'X' &&
				str[idx+height-1] == 'M' &&
				str[idx+height*2-2] == 'A' &&
				str[idx+height*3-3] == 'S' {
				count++
			}
			// Diagonal up right
			if x < width-3 && y >= 3 &&
				str[idx] == 'X' &&
				str[idx-height+1] == 'M' &&
				str[idx-height*2+2] == 'A' &&
				str[idx-height*3+3] == 'S' {
				count++
			}
			// Diagonal up left
			if x >= 3 && y >= 3 &&
				str[idx] == 'X' &&
				str[idx-height-1] == 'M' &&
				str[idx-height*2-2] == 'A' &&
				str[idx-height*3-3] == 'S' {
				count++
			}
		}
	}

	fmt.Printf("[DAY 4 - Part 1] - XMAS Count: %d\n", count)
}

func day4Part2() {
	str := utils.ReadFileToString("files/day4.txt")
	
	var separator string
	if strings.Contains(str, "\r\n") {
		separator = "\r\n"
	} else {
		separator = "\n"
	}

	width := strings.Index(str, separator)
	height := strings.Count(str, separator) + 1 // FILE SHOULDN'T END WITH A NEW LINE
	str = strings.Replace(str, separator, "", -1)

	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := x + y*width
			diagonalOne := false
			diagonalTwo := false

			// Diagonal one, down right and up left
			if x < width-2 && y < height-2 {
				if str[idx+height*0+0] == 'M' &&
					str[idx+height*1+1] == 'A' &&
					str[idx+height*2+2] == 'S' {
					diagonalOne = true
				} else if str[idx+height*2+2] == 'M' &&
					str[idx+height*1+1] == 'A' &&
					str[idx+height*0+0] == 'S' {
					diagonalOne = true
				}
			}

			// Diagonal two, up right and down left
			if x < width-2 && y < height-2 {
				if str[idx+height*2+0] == 'M' &&
					str[idx+height*1+1] == 'A' &&
					str[idx+height*0+2] == 'S' {
					diagonalTwo = true
				} else if str[idx+height*0+2] == 'M' &&
					str[idx+height*1+1] == 'A' &&
					str[idx+height*2+0] == 'S' {
					diagonalTwo = true
				}
			}

			if diagonalOne && diagonalTwo {
				count++
			}
		}
	}

	fmt.Printf("[DAY 4 - Part 2] - X-MAS Count: %d\n", count)
}

func Day4() {
	day4Part1()
	day4Part2()
}
