package main

import "fmt"

func cag () {
	// problemOneCag()
	// problemTwoCag()
	// problemThreeCag()
	problemFourCag()
	// problemFiveCag()
}

func problemOneCag() {
	variableOne := make([]int, 4, 10)

	fmt.Println("Before Append: ")
	fmt.Println("Length: ", len(variableOne), " Capacity: ", cap(variableOne))
	
	variableOne = append(variableOne, 1)
	
	fmt.Println("After Append: ")
	fmt.Println("Length: ", len(variableOne), " Capacity: ", cap(variableOne))
}

func problemTwoCag() {
	variableOne := []int{}

	fmt.Println("Length: ", len(variableOne), " Capacity: ", cap(variableOne))
	for i := 0; i < 10; i++ {
		variableOne = append(variableOne, 1)
		fmt.Println("Length: ", len(variableOne), " Capacity: ", cap(variableOne))
	}
}

func problemThreeCag() {
	variableOne := make([]int, 3, 50)

	fmt.Println("Before Append:")
	fmt.Println("Length: ", len(variableOne), " Capacity: ", cap(variableOne))

	for i := 0; i < 10; i++ {
		variableOne = append(variableOne, 1)
	}
	
	fmt.Println("After Append:")
	fmt.Println("Length: ", len(variableOne), " Capacity: ", cap(variableOne))
}

func problemFourCag(){
	variableOne := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Slice Length and Capacity")
	fmt.Println("Length: ", len(variableOne), " Capacity: ", cap(variableOne))

	variableSubSlice := variableOne[2:5]

	fmt.Println("Sub Slice Length and Capacity")
	fmt.Println("Length: ", len(variableSubSlice), " Capacity: ", cap(variableSubSlice))
	fmt.Println(variableSubSlice)
}


func problemFiveCag(){
	var nilVariable []int
	preAllocated := make([]int, 3, 20)

	nilVariable = append(nilVariable, 1230123)
	preAllocated = append(preAllocated, 1230123)

	fmt.Println("Nil Variable:")
	fmt.Println("Length: ", len(nilVariable), " Capacity: ", cap(nilVariable))
	fmt.Println("Pre Allocated:")
	fmt.Println("Length: ", len(preAllocated), " Capacity: ", cap(preAllocated))
}