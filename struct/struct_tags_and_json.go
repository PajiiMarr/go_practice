package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func staj() {
	// problemOneStaj()
	// problemTwoStaj()
	// problemThreeStaj()
	// problemFourStaj()
	problemFiveStaj()
}

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"`
}

func problemOneStaj() {
	variablePerson := Person{Name: "Marlo"}

	b, _ := json.Marshal(variablePerson)

	fmt.Println(string(b))
}

func problemTwoStaj() {
	variableOne := Person{Name: "Mar", Age: 0}

	b, _ := json.Marshal(variableOne)

	fmt.Println(string(b))
}

func problemThreeStaj() {
	variablePerson := Person{Name: "Mar", Age: 23, Password: "zamboanga"}

	b, _ := json.Marshal(variablePerson)

	fmt.Println(string(b))
}

func problemFourStaj() {
	variableJson := []byte(`{"name": "Mar", "age": 23, "password": "zamboanga"}`)

	var variablePerson map[string]interface{}

	json.Unmarshal(variableJson, &variablePerson)

	fmt.Println(variablePerson)
}

func problemFiveStaj () {
	variableJson := []byte(`{"name": "Mar", "age": 23, "password": "zamboanga"}`)
	
	var variablePerson Person 
	json.Unmarshal(variableJson, &variablePerson)

	val := reflect.ValueOf(variablePerson)
	typ := reflect.TypeOf(variablePerson)	

	for i := 0; i < val.NumField(); i ++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		fmt.Printf("Field %q: type=%s, value=%v\n", fieldType.Name, field.Type(), field.Interface())
	}
}