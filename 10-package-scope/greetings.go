package main // we declare the package main in order to use it in the main.go file which is the entry point of the program

import "fmt"

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("Hello", n)
}

// Calling the variable score from main.go in greetings.go
func showScore() {
	fmt.Println("Your score is: ", score)
}
