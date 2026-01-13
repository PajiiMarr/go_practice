package main

import "fmt"

func es () {
	problemOneEs()
	problemTwoEs()
	problemThreeEs()
	problemFourEs()
	problemFiveEs()
}

// 1️⃣ Basic Field Promotion (Easy)

// Embed a simple struct with a few fields into another struct and access those fields directly from the outer struct.

// Hints:
// Field promotion · anonymous field · direct access

type Address struct {
	Barangay string
	City string
	ZipCode int32
}

type Information struct {
	Name string
	Age int16
}

type Data struct {
	Name string
	Age int16
}

type Citizen struct {
	Address
}

func problemOneEs (){
	variableCitizen := Citizen{
		Address: Address{
			Barangay: "Pasonanca",
			City: "Zamboanga City",
			ZipCode: 7000,
		},
	}

	fmt.Println(variableCitizen.Barangay)
	fmt.Println(variableCitizen.City)
	fmt.Println(variableCitizen.ZipCode)
}

// 2️⃣ Method Promotion (Easy–Medium)

// Embed a struct that has methods and call those methods from the outer struct without referencing the inner struct explicitly.

// Hints:
// Method forwarding · promoted methods · receiver behavior

func (a Address) Input () {
	fmt.Println("Address inputted: ", a.Barangay, a.City, a.ZipCode )
}

type CitizenMethod struct {
	Address
}

func problemTwoEs () {
	variableCitizen := CitizenMethod{
		Address: Address{
			Barangay: "Pasonanca", 
			City: "Zamboanga City",
			ZipCode: 7000,
		},
	}

	variableCitizen.Input()
}

// 3️⃣ Method Override (Medium)

// Embed a struct that has a method, then define a method with the same name on the outer struct and control which one gets called.

// Hints:
// Shadowing · explicit qualification · behavior control

type CitizenOverride struct {
	Address
}

func (c CitizenOverride) Input () {
	fmt.Println("Method overrider for address: ", c.Barangay, c.City, c.ZipCode)
}

func problemThreeEs() {
	variableCitizen := CitizenOverride{
		Address: Address{
			Barangay: "Pasonanca",
			City: "Zamboanga City",
			ZipCode: 7000,
		},
	}

	variableCitizen.Input()
	variableCitizen.Address.Input()
}

// 4️⃣ Multiple Embedded Structs (Medium–Hard)

// Embed two or more structs that contain fields or methods, and resolve access when names collide.

// Hints:
// Name conflicts · ambiguity · qualified access

type CitizenMultiStruct  struct{
	Address
	Information
	Data
}


func problemFourEs () {
	variableCitizen := CitizenMultiStruct{
		Address: Address{
			Barangay: "Pasonanca",
			City: "Zamboanga City",
			ZipCode: 7000,
		},
		Information: Information{
			Name: "Marlo",
			Age: 32,
		},
		Data: Data{
			Name: "Mernie",
			Age: 32,
		},
	}

	variableCitizen.Input()

	fmt.Println(variableCitizen.Information.Name)
	fmt.Println(variableCitizen.Data.Name)
}

// 5️⃣ Embedding + Interfaces (Hard)

// Embed a struct so that the outer struct automatically satisfies an interface without explicitly implementing all methods.

// Hints:
// Interface satisfaction · method sets · implicit implementation

type Inputter interface {
	Input()
}

type CitizenInterface struct {
	Address
}

func (ci CitizenInterface) Input () {
	fmt.Println("Citizen Interface: ", ci.Address)
}

func problemFiveEs () {
	variableCitizen := CitizenInterface{
		Address: Address{
			Barangay: "Pasonanca",
			City: "Zamboanga City",
			ZipCode: 7000,
		},
	}

	variableCitizen.Input()
}