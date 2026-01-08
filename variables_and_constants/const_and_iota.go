package main

import "fmt"

func cai() {
	problem_one()
	problem_two()
	problem_three()
	problem_four()
	problem_five()
}

func problem_one() {
	const pi float64 = 3.14
	const language string = "Go"
	const year int = 2026

	fmt.Println("pi: ", pi)
	fmt.Println("language: ", language)
	fmt.Println("year: ", year)
}

func problem_two () {
	const typedInt int = 10
	const untypedInt = 20

	fmt.Println("\ntypedInt: ", typedInt)
	fmt.Println("untypedInt: ", untypedInt)
	fmt.Println("typedInt + untypedInt: ", typedInt + untypedInt)
}

func problem_three() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	days := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	values := []int {
		Sunday,
		Monday,
		Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday,
	}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i] + ": ", values[i])
	}
}

func problem_four() {
	const (
		_ = iota 
		Read
		Write
		Execute
	)

	variables := []string{
		"Read",
		"Write",
		"Execute",
	}

	values := []int{
		Read,
		Write,
		Execute,
	}

	fmt.Println()
	for i := 0; i < len(variables); i++ {
		fmt.Println(variables[i]+":", values[i])
	}
}


func problem_five() {
	const (
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)

	filesizes := []string {
		"KB",
		"MB",
		"GB",
		"TB",
	}

	values := []float64 {
		KB,
		MB,
		GB,
		TB,
	}

	for i := 0; i < len(filesizes); i++ {
		fmt.Println(filesizes[i]+": ", values[i])
	}
}