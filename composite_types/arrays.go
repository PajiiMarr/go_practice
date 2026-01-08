package main 

import "fmt"

func main () {
	// problemOneA()
	problemTwoA()
	// problemThreeA()
	// problemFourA()
	// problemFiveA()
}

func problemOneA() {
	var variableOne [5]int

	fmt.Println("Array declaration: ", variableOne)
}

func problemTwoA(){
	var variableOne = [4]float64{0.2, 0.4, 0.6, 0.8}

	for i, val := range variableOne {
		fmt.Println("Array:", i, "Value:", val)
	}
}

func problemThreeA() {
	var original = [4]string{"hello", "Hwllo", "HwllOO", "H343ello"}
	var copy = original

	copy[3] = "Heeeellllooooo"

	fmt.Println("Original Array: ", original)
	fmt.Println("Copy Array: ", copy)
}

func problemFourA(){
	var variableOne = [6]int{1,4,5,8,3,2}

	var sum int = 0

	for i := 0; i < len(variableOne); i++ {
		sum += variableOne[i]
	}


	fmt.Println("Sum: ", sum)
	fmt.Println("Average: ", float32(sum)/6)
}

func problemFiveA () {
	var variableOne = [2][3]int {
		{1,2,3},
		{4,5,6},
	}

	var sum int = 0

	for i := 0; i < len(variableOne); i++ {
		fmt.Println("\nRow ", i+1)
		for j := 0; j < len(variableOne[i]); j++ {
			fmt.Print(variableOne[i][j], " ")
			sum += variableOne[i][j]
		}
	}

	fmt.Println("\nSum: ", sum)
}