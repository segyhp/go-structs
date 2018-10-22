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
	hendro := person{firstName: "Segy", lastName: "Pratama"}

	fmt.Println(hendro)
}
