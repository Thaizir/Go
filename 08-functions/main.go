package main

import (
	"fmt"
	"math"
)

// We can create functions outside of the main function and call them in the main function on in the future if we have other files

// We have to declare the type of the parameter
func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

// We can pass a function as a parameter to another function but we have to declare the type of the parameter that will be passed to the function in the argument as well
func cicleNames(n []string, f func(string)) {

	for _, v := range n {
		f(v)
	}

}

// you can also type the return type of a function
func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {

	// Calling functions
	sayGreeting("Link")
	sayBye("Link")

	names := []string{"Link", "Zelda", "Sheik", "Ganondorf"}
	cicleNames(names, sayGreeting)

	a1 := circleArea(5)
	a2 := circleArea(20.3)
	fmt.Println(a1, a2)

	// We can truncate floats by using the %f format specifier for example with %.5f we get the first 5 decimals
	fmt.Printf("Area 1 is: %.2f \n", a1)
	fmt.Printf("Area 2 is: %.2f \n", a2)

}
