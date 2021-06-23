//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

import "fmt"

// importing fmt package

//gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

// main method
func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("Square ", square, "Cube", cube)
}
