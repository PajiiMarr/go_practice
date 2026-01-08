package main

import (
	"fmt"
	"encoding/json"
)

func isu () {
	problem_one()
	problem_two()
	problem_three()
	problem_four()
	problem_five()

}

func problem_one () {
	var int8_variable int8 = -128     // min value for int8
	var uint8_variable uint8 = 255    // max value for uint8
	var int16_variable int16 = 32767  // max value for int16
	var uint16_variable uint16 = 65535 // max value for uint16

	fmt.Println("Int8 Variable:", int8_variable)
	fmt.Println("UInt8 Variable:", uint8_variable)
	fmt.Println("Int16 Variable:", int16_variable)
	fmt.Println("UInt16 Variable:", uint16_variable)
}

func problem_two () {
	var variableOne int8

	fmt.Print("\nInput a number (-127 to 127): ")
	fmt.Scan(&variableOne)

	fmt.Println("You inputted: ", variableOne)

	const maxInt int8 = 127

	if variableOne == maxInt {
		fmt.Println("Adding 1 would cause int8 overflow!")
	}
}

func problem_three () {
	var age uint8
	var daysAlive int16

	fmt.Print("Age: ")
	fmt.Scan(&age)
	fmt.Print("Days Alive: ")
	fmt.Scan(&daysAlive)

	fmt.Println("Age: ", age, "Days Alive: ", daysAlive)
}

func problem_four() {
	var tempChange int16 = 150
	var totalSensorReadings int16 = 6

	average := float64(tempChange) / float64(totalSensorReadings)

	fmt.Println("Average Change per Reading: ", average)
}


type User struct {
	UserId uint64
	AccountBalance float64
	LoginCount uint16
	ErrorCode int8
}
 
func problem_five() {
	u1 := User{
		UserId: 12345,
		AccountBalance: 43571.23,
		LoginCount: 12,
		ErrorCode: -123,
	}

	u2 := User{
		UserId: 12346,
		AccountBalance: 23523.23,
		LoginCount: 7,
		ErrorCode: 127,
	}
	
	b, _ := json.MarshalIndent(u1, "", "  ")
	fmt.Println(string(b))
	a, _ := json.MarshalIndent(u2, "", "  ")
	fmt.Println(string(a))
}
