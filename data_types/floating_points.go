package main

import "fmt"

func fp() {
	problem_one_fp()
	problem_two_fp()
	problem_three_fp()
	problem_four_fp()
	problem_five_fp()
}

func problem_one_fp() {
	var float32Var float32 = 3.14
	var float64Var float64 = 3.14159265359

	fmt.Printf("float32Var: %.10f", float32Var)
	fmt.Printf("\nfloat64Var: %.10f", float64Var)
}

func problem_two_fp() {
	float32Var := float32(3.14)
	float64Var := float64(3.14159265359)

	sum := float64(float32Var) + float64Var
	difference := float64(float32Var) - float64Var
	product := float64(float32Var) * float64Var
	quotient := float64(float32Var) / float64Var

	fmt.Printf("\nSum: %.10f", sum)
	fmt.Printf("\nDifference: %.10f", difference)
	fmt.Printf("\nProduct: %.10f", product)
	fmt.Printf("\nQuotient: %.10f", quotient)
}

func problem_three_fp() {
	float32Var := float32(0.0)
	float64Var := float64(0.0)

	for i := 0; i < 10; i++ {
		float32Var += 0.1
		float64Var += 0.1
	}

	fmt.Printf("\nfloat32Var: %.10f", float32Var)
	fmt.Printf("\nfloat64Var: %.10f", float64Var)
}

func problem_four_fp() {
	float32Var := float32(3.14)
	float64Var := float32(14.32512)

	product := float32Var * float32(float64Var)

	fmt.Printf("%.10f", product)
}

func problem_five_fp() {
	const pie float64 = 3.14
	var radius float64

	fmt.Println("Input radius: ")
	fmt.Scan(&radius)

	area := pie * radius * radius

	fmt.Printf("Area: %.15f", area)
}
