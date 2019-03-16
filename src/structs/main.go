package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Smith"
	// fmt.Printf("Alex fullname is %s \n", alex)

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Alder",
	// 	contactInfo: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 98074,
	// 	},
	// }

	// jim.updateName("James")
	// jim.print()

	mySlice := []string{"1", "2", "3", "4", "5"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(mySlice []string) {
	mySlice[0] = "11"
}

// func (p *person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// func (p person) print() {
// 	fmt.Printf("%+v ", p)
// }
