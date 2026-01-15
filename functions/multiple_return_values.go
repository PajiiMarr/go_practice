package main

import (
	"errors"
	"fmt"
	"strconv"
) 

func mrv() {
	q, r, err := problemOneMrv(9,7)
	if err != nil {
		fmt.Print(err)
		return 
	}
	fmt.Printf("Quotient and remainder: %d, %d\n", q, r)


	valid, failed := problemTwoMrv("10", "42", "hello", "7", "3.14", "100")

	fmt.Println("Valid: ", valid)
	fmt.Println("Failed: ", failed)

}

// 2️⃣ Multiple Return Values
// Challenge 1 (Easy)
// Divide two integers
// Return quotient and remainder
// Handle division by zero with error

func problemOneMrv(num1 int, num2 int) (int, int, error) {
	if num2 == 0 {
		return 0, 0, errors.New("Division by zero")
	}

	quotient := num1 / num2
	remainder := num1 % num2

	return quotient, remainder, nil
}

// Challenge 2 (Medium-Hard)
// Parse list of strings into integers
// Return two slices: valid numbers and failed parses
// Caller must handle both

func problemTwoMrv (word ...string)([]int, []string) {
	valid := []int{}
	failed := []string{}

	for _, v := range word {
		n, err := strconv.Atoi(v)

		if err != nil {
			failed = append(failed, v)
		} else {
			valid = append(valid, n)
		}
	}

	return valid, failed
}