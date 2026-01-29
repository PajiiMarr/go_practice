package main

import "fmt"

func pf() {
	// problemOnePf()
	// problemTwoPf()
	problemThreePf()
}

// 🟢 EASY — Pointer Fundamentals
// 1️⃣ Value vs Pointer Modification

// Goal:
// Demonstrate the difference between passing a value and passing a pointer to a function.

// Hints:

// Pass an integer two ways
// Modify inside function
// Compare results

func modifyByValue(x int) {
	x = 765
}

func modifyByPointer(x *int) {
	*x = 765
}

func problemOnePf() {

	x := 896
	fmt.Println(x)

	modifyByValue(x)
	fmt.Println(x)

	modifyByPointer(&x)
	fmt.Println(x)
}

// 2️⃣ Swap Two Values

// Goal:
// Write a function that swaps two integers.

// Hints:

// Use pointers
// No return values
// Caller variables must change
func problemTwoPointer(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func problemTwoPf() {
	x := 1
	y := 5

	fmt.Println("Before swap", x, ' ', y )
	
	problemTwoPointer(&x, &y)
	fmt.Println("After swap", x, ' ' ,y)
}


// 3️⃣ Reset Struct Fields

// Goal:
// Write a function that resets all fields of a struct to default values.

// Hints:

// Pass struct by pointer

// Modify fields directly

// No struct copying

type User struct {
	Name string 
	Age int
}

func DefaultUser() User {
	return User{
		Name: "Marlo",
		Age: 19,
	}
}

func resetStructPointer(u *User){
	*u = DefaultUser()
}

func problemThreePf() {
	user := User{
		Name: "Maria",
		Age: 19,
	}

	fmt.Println("Before reset: ")
	fmt.Println(user)

	fmt.Println("After Reset: ")
	resetStructPointer(&user)
	fmt.Println(user)
}