package main

import "fmt"

func af() {
	problemOneAf()
	sliceVariable := []int {
		1,2,3,4,
	}
	result := problemTwoAF(sliceVariable, func(n int) int {
		return n * 2
	})

	fmt.Println(result)
}

// Challenge 1 (Easy)
// Assign function to variable that doubles an integer
// Call and print result

func problemOneAf() {
	num := func(a int) int {
		return a + a
	}

	double_num := num(2)

	fmt.Println("Double number", double_num)
}

// Challenge 2 (Medium)
// Pass anonymous function as argument to another function
// Apply function to slice of integers and print results



func problemTwoAF(nums []int, op func(int) int ) []int {
	result := []int{}
	for _, n := range nums {
		result = append(result, op(n))
	}
	return result
}

