package main

import (
	"fmt"
)

func fr () {
	problemOneFr()
	problemTwoFr()
	problemThreeFr()
	problemFourFr()
	problemFiveFr()
}

// 1️⃣ Slice Sum (Easy)
// Iterate slice
// Add all numbers
// Print total
func problemOneFr(){
	variableOne := []int{1,2,3,4,5,}

	sum := int(0)

	for n := range variableOne {
		sum += variableOne[n]
	}

	fmt.Println("Sum is: ", sum)
}


// 2️⃣ Map Keys Collector (Medium-Easy)
// Iterate map
// Collect only keys
// Store in slice
func problemTwoFr () {
	variableOne := map[string]int {
		"apple": 1,
		"orange": 2,
		"carrots": 3,
	}

	variableSlice := []string{}

	for i := range variableOne {
		variableSlice = append(variableSlice, i)
	}

	fmt.Println("Your variable slice: ", variableSlice)
}


// 3️⃣ Unicode Counter (Medium)
// Iterate string
// Count all unique runes
// Handle emojis & accents

func problemThreeFr () {
	sampleText := "Go ❤️ gophers 🐹 café naïve 🚀❤️"
	seen := make(map[rune]bool)
	count := int(0)
	for i, r := range sampleText {
		fmt.Printf("Index %d: %c\n", i, r)
		
		if !seen[r] {
			seen[r] = true
			count++
		}
	}	

	fmt.Println("All runes: ", count)
}


// 4️⃣ Channel Aggregator (Medium-Hard)
// Iterate channel until closed
// Compute sum or average
// Handle dynamic input
func problemFourFr() {
	variableLength := int(0)

	fmt.Print("Length: ")
	fmt.Scan(&variableLength)

	variableOne := make(chan int, variableLength)
	for i := 0; i < variableLength; i++ {
		var input int 
		fmt.Print("Channel Element ", i, ": ")
		fmt.Scan(&input)
		variableOne <- input
	}

	close(variableOne)

	sum := 0
	count := 0

	for v := range variableOne {
		sum += v
		count++
	}

	fmt.Println("Sum: ", sum)
	fmt.Println("Average: ", float32(sum)/float32(count))

}


// 5️⃣ Map Filter & Delete (Hard)
// Iterate map
// Delete entries meeting condition
// Keep remaining entries safe

func problemFiveFr () {
	variableMap := map[string]int{
		"orange": 5,
		"apple": 5,
		"banana": 3,
		"red": 10,
	}

	for fruit, qty := range variableMap{
		if qty > 5{
			delete(variableMap, fruit)
			continue
		}

		fmt.Println("Safe values: ", fruit , ", ", qty)
	}
}