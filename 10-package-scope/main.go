package main

import "fmt"

// in greetings.go we can access the variables and function declared in main.go

var score = 99.5

func main() {

	// var score = 99.5 If we declarre the varible within the main function it will be available only within the main function and not globally (is not in the package scope)

	// We can't use the variable points outside the main function but we need to run main.go and greetings.go in order to see the results
	sayHello("James Bond")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}
