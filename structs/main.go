package main

import "fmt"



type person struct{
	firstName string
	lastName string

}



func main()  {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// //var alex person
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
}