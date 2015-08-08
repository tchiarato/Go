package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]

	statistics := processStatistics(words)

	printResult(statistics)
}

func processStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		letter := strings.ToUpper(string(word[0]))
		count, isFound := statistics[letter]

		if isFound {
			statistics[letter] = count + 1
		} else {
			statistics[letter] = 1
		}
	}

	return statistics
}

func printResult(statistics map[string]int) {
	fmt.Println("Counter of words that starts with the same letter!")

	for letter, count := range statistics {
		fmt.Printf("%s = %d\n", letter, count)
	}
}
