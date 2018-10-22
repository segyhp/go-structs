package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	// 1st way
	// alex := person{"Alex", "Anderson"}

	//2nd way
	// hendro := person{firstName: "Segy", lastName: "Pratama"}

	//3rd way
	var hendro person

	hendro.firstName = "Segy"
	hendro.lastName = "Pratama"

	fmt.Println(hendro)
	fmt.Printf("%+v", hendro)

}
