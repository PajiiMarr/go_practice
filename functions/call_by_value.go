package main

import "fmt"

func cbv () {
	// problemOneCbv()
	problemTwoCbv()
}

// 6️⃣ Call by Value / Pointers

// Challenge 1 (Easy)
// Pass integer to function
// Modify inside function using pointer
// Show difference between value vs pointer

func modifyOne (x *int) {
	*x = 100
}

func problemOneCbv() {
	x := 50
	fmt.Println(x)

	modifyOne(&x)
	fmt.Println(x)
}

// Challenge 2 (Medium-Hard)

// Pass slice or struct to function
// Modify elements inside function
// Return original vs modified values to demonstrate memory behavior

type User struct {
	Name string
	Age int
}

func modifyTwo(x *User){
	x.Name= "Marlo"
}

func problemTwoCbv () {
	name := User{
		Name: "Maria",
	}

	fmt.Println(name)
	modifyTwo(&name)

	fmt.Println(name)
}