package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	params := os.Args[1:]
	numbers := make([]int, len(params))

	for i, n := range params {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s isn't a valid number!\n", n)
			os.Exit(1)
		}
		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	less, great := partition(n, pivot)

	return append(
		append(quicksort(less), pivot),
		quicksort(great)...)
}

func partition(
	numbers []int,
	pivot int) (less []int, great []int) {
	for _, n := range numbers {
		if n <= pivot {
			less = append(less, n)
		} else {
			great = append(great, n)
		}
	}

	return less, great
}
