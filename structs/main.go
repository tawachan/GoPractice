package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	sayaka := person{
		firstName: "Sayaka",
		lastName:  "Sasaki",
		contact: contactInfo{
			email:   "sayaka@gmail.com",
			zipCode: 12345,
		},
	}
	sayaka.print()
	sayakaPointer := &sayaka
	sayakaPointer.updateName("Saya")
	sayaka.print()

	colors := map[string]string{
		"white": "#FFFFFF",
	}

	tests := make(map[string]string)

	colors["red"] = "red"

	delete(colors, "red")

	fmt.Println(colors)

	fmt.Println(tests)
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (PointerToPerson *person) updateName(newFirstName string) {
	(*PointerToPerson).firstName = newFirstName
}
