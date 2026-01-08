package main

import "fmt"

func vvc () {
	var integerVariable int
	var float64Variable float64
	var stringVariable string
	var booleanVariable bool

	fmt.Println("integerVariable: ", integerVariable)
	fmt.Println("float64Variable: ", float64Variable)
	fmt.Println("stringVariable: ", stringVariable)
	fmt.Println("booleanVariable: ", booleanVariable)

	problemTwo()
	problemThree()
	problemFour()
	problemFive()
}

func problemTwo () {
	var variableOne int8 = 12
	variableTwo := int8(12)

	fmt.Println("\nvariableOne: ", variableOne)
	fmt.Println("variableTwo: ", variableTwo)
}

func problemThree () {
	var variableOne int8 = 24
	
	fmt.Println("\nBefore redeclaration: ", variableOne)
	{
		variableOne := int8(25)
		fmt.Println("During redeclaration: ", variableOne)
	}

	fmt.Println("After block: ", variableOne)
}

func problemFour() {
	var variableOne,  variableTwo ,  variableThree int8 = 1,2,3
	variableFour, variableFive, variableSix := int8(4), int8(5), int8(6)
	
	fmt.Println("\nvar: ", variableOne, variableTwo, variableThree)
	fmt.Println(":= ", variableFour, variableFive, variableSix)

	variableOne = 8
	variableFour = int8(23)

	fmt.Println("Reassignment variableOne: ", variableOne)
	fmt.Println("Reassignment variableFour: ", variableFour)
}

func problemFive() {
	// var int8Variable int8 = 10
	// var uint8Variable uint8 = 20
	// var int16Variable int16 = -300
	// var uint16Variable uint16 = 60000

	// var name string = "Mar"
	// var isActive bool = true
	// var score float64 = 98.6

	// counter := 0
	// message := "Hello, Go!"
	// flag := false
	// pi := 3.1415
}