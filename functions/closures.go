package main

import "fmt"

func c() {
	// counter := problemOneC()

	// fmt.Println(counter())
	// fmt.Println(counter())
	// fmt.Println(counter())


	add, remove, get := problemTwoC()

	add("apple", 2)
	add("apple", 3)
	add("banana", 1)

	fmt.Println(get("apple"))  // 5
	fmt.Println(get("banana")) // 1

	remove("banana")
	fmt.Println(get("banana")) // 0
}

// 4️⃣ Closures
// Challenge 1 (Medium)
// Create a counter function
// Returns a function that increments and prints counter each call

func problemOneC() func() int {
	counter := 0

	return func() int {
		counter += 1
		return counter	
	}
}

// Challenge 2 (Hard)

// Create a closure that maintains a map of string → int
// Returns functions to add, remove, and query counts
// All state must stay within closure

func problemTwoC() (add func(string, int), remove func(string), get func(string) int) {
	store := make(map[string]int)

	add = func(key string, value int) {
		store[key] += value
	}

	remove = func (key string)  {
		delete(store, key)
	}

	get = func(key string) int {
		return store[key]
	}


	return
}