package main

import "fmt"

func vf () {
	// fmt.Println(problemOneVf(1,7,2,74,12))
	fmt.Println("Formatted Word: ", problemTwoVf("i", "love", "maria"))
}

// Challenge 1 (Easy)
// Accept any number of integers
// Return sum
// Print result

func problemOneVf (nums ...int)int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

// Challenge 2 (Medium)
// Accept strings of any length
// Concatenate with separator
// Return formatted string
func problemTwoVf (words ...string) string {
	formattedWord := ""
	for i, v := range words {
		if i > 0 {
			formattedWord += " "
		}

		formattedWord += v
	}

	return formattedWord
}