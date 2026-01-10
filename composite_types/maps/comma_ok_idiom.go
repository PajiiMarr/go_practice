package main

import "fmt"

func coi () {
	problemOneCoi()
	problemTwoCoi()
	problemThreeCoi()
	problemFourCoi()
	problemFiveCoi()
}

func problemOneCoi() {
	variableOne := map[string]int{
		"apple": 3,
		"banana": 2,
		"orange": 5,
	}

	if v, ok := variableOne["apple"]; ok {
		fmt.Println("Apple exist, quantity: ", v, ok)
	} else {
		fmt.Println("Apple don't exist")
	}

	if v, ok := variableOne["banana"]; ok {
		fmt.Println("Banana exist, quantity: ", v, ok)
	} else {
		fmt.Println("Banana don't exist")
	}

	if v, ok := variableOne["berry"]; ok {
		fmt.Println("Berry exist, quantity: ", v, ok)
	} else {
		fmt.Println("Berry don't exist", v, ok)
	}
}


func problemTwoCoi() {
	variableOne := map[int]string {
		0: "hello",
		1: "hii",
		2: "hihihih",
	}

	v, ok := variableOne[5]
	fmt.Println("Result: ", v, " ", ok)
}

func problemThreeCoi() {
	variableOne := map[string]int {
		"zero": 0,
		"one": 1,
		"two": 2,
	}

	for _, key := range []string{"zero", "three"} {
		v, ok := variableOne[key]

		if ok {
			fmt.Printf("Key %q exist, key = %d\n", key, v)
		} else {
			fmt.Printf("Key %q does not exist\n", key)
		}
	}
}

func problemFourCoi() {
	var variableOne interface{} = "Hello"

	if s, ok := variableOne.(int); ok {
		fmt.Println("Result: ", s, ok)
	} else {
		fmt.Println("Invalid type")
	}
}

func problemFiveCoi(){
	variableOne := map[string]interface {}{
		"zero": 0,
		"one": true,
		"two": "dos",
	}

	for key, val := range variableOne {
		switch v := val.(type) {
			case int:
				fmt.Printf("Key %q is int: %d\n", key, v)
			case string:
				fmt.Printf("Key %q is string: %q\n", key, v)
			case bool:
				fmt.Printf("Key %q is bool: %t\n", key, v)
			default:
				fmt.Printf("Key %q has unknown type\n", key)
		}
	}

}