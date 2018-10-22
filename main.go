package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// 1st way
	// alex := person{"Alex", "Anderson"}

	//2nd way
	// hendro := person{firstName: "Segy", lastName: "Pratama"}

	//3rd way
	// var hendro person
	// hendro.firstName = "Segy"
	// hendro.lastName = "Pratama"
	// fmt.Println(hendro)
	// fmt.Printf("%+v", hendro)

	hendro := person{
		firstName: "Segy",
		lastName:  "Pratama",
		contact: contactInfo{
			email:   "segyhendropratama@gmail.com",
			zipCode: 16112,
		},
	}

	// fmt.println(hendro)
	fmt.Printf("%+v", hendro)

}
