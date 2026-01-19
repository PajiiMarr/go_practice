package main

import "fmt"

// import (
// 	"fmt"
// )

func nr () {
	// product := problemOneNr(1, 7)

	// fmt.Println(product)

	sliceVar := []int {
		1,2,3,4,5,
	}
	fmt.Println(problemTwoNr(sliceVar))
}




// 5️⃣ Named Return Values
// Challenge 1 (Easy)
// Function returns two named integers
// Perform calculation inside function and return without explicit return values

func problemOneNr (name1 int, name2 int)  (product int) {
	product = name1 * name2
	return
}



// Challenge 2 (Medium)

// Function computes statistics of a slice: sum, average, max, min
// Use named return values
// Return all values in single return

func problemTwoNr(nums []int) (sum int, avg float32, max int, min int ){
	if len(nums) == 0 {
		return
	}

	min, max = nums[0], nums[0]

	for _, n := range nums {
		sum  += n

		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}

	avg = float32(sum) / float32(len(nums))
	return	
}