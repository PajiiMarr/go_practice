package main

import "fmt"

func af() {
	problemOneAf()
}

// 3️⃣ Anonymous Functions
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