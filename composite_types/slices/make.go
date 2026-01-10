package main

import "fmt"

func makeSLice(){
	// problemOneMake()
	// problemTwoMake()
	// problemThreeMake()
	// problemFourMake()
	problemFiveMake()
}

func problemOneMake(){
	variableOne := make([]int, 10, 50)

	variableOne = append(variableOne, 10)
}

func problemTwoMake() {
	variableOne := make(map[string]int, 10)

	variableOne["apple"] = 23
	variableOne["orange"] = 24

	fmt.Println("Variable One: ", variableOne["apple"])
	fmt.Println("Missing key: ", variableOne["pineapple"])

}

func problemThreeMake () {
	variableOne := make([]int, 10, 50)
	variableTwo := variableOne

	variableTwo[0] = 99

	fmt.Println("Variable One: ", variableOne)
	fmt.Println("Variable Two: ", variableTwo)
}

func problemFourMake () {
	variableOne := make(chan []int, 4)

	variableOne <- []int {2,2,4,1}
	value := <-variableOne

	fmt.Println("Received value from channel: ", value)
}


func problemFiveMake() {
	sliceVariable := []string{"task1", "task2", "task3", "task4"}
	mapVariable := make(map[string]string, len(sliceVariable))

	for _, task := range sliceVariable {
		mapVariable[task] = "pending "
	}

	channelVariable := make(chan string, len(sliceVariable))

	for _, task := range sliceVariable {
		go func(t string) {
			fmt.Println("Processing:", t)
			
			mapVariable[t] = "done"
			
			channelVariable <- t
		}(task)
	}

	for i := 0; i < len(sliceVariable); i++ {
		completedTask := <-channelVariable
		fmt.Println("Completed:", completedTask)
	}

	fmt.Println("Final task statuses:")
	for task, status := range mapVariable {
		fmt.Println(task, "=>", status)
	}
}