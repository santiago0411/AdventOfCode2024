package main

import "fmt"

func Day2() {
	list1, list2 := ReadLists()
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

	fmt.Printf("[DAY 2] - Similarity is %d\n", similarity)
}
