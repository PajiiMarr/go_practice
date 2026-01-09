package main

import (
	"fmt"
	"math/cmplx"
)

func cn() {
	problem_one_cn()
	problem_two_cn()
	problem_three_cn()
	problem_four_cn()
	problem_five_cn()
}

func problem_one_cn() {
	variableOne := complex64(complex(3, 4))
	variableTwo := complex(3, 4)

	fmt.Println("VariableOne: ", variableOne)
	fmt.Println("variableTwo: ", variableTwo)
}

func problem_two_cn() {
	variableOne := complex(27, 4)

	fmt.Println("Real Part: ", real(variableOne), " Imaginary Part: ", imag(variableOne))
}

func problem_three_cn() {
	variableOne := complex(4, 5)
	variableTwo := complex(6, 7)

	sum := variableOne + variableTwo
	difference := variableOne - variableTwo
	product := variableOne * variableTwo
	quotient := variableOne / variableTwo

	fmt.Printf("\nSum: %.10f", sum)
	fmt.Printf("\nDifference: %.10f", difference)
	fmt.Printf("\nProduct: %.10f", product)
	fmt.Printf("\nQuotient: %.10f", quotient)
}

func problem_four_cn() {
	variableOne := complex(3, 4)

	absoluteVariable := cmplx.Abs(variableOne)

	fmt.Println("Absolute Value: ", absoluteVariable)

}

func problem_five_cn() {
	original := complex(0.0, 1.0) // complex128
	rotated := original * complex(0, 1)

	fmt.Println("Original:", original)
	fmt.Println("Rotated:", rotated)
}
