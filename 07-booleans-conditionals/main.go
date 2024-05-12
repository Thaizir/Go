package main

import (
	"fmt"
)

func main() {

	age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// Conditional Statements: Works the same as in JS

	if age <= 30 {
		fmt.Println("Age is less than 30")
	} else if age >= 50 {
		fmt.Println("Age is greater than 50")
	} else {
		fmt.Println("Age is between 30 and 50")
	}

	// Continue and Break: Works the same as in JS

	names := []string{"Link", "Zelda", "Sheik", "Ganondorf"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	// There's no built in ternary operator in Go, it is a design choice made by the Go creators. One of the reasons being overuse and abuse of that operator to create extremely complex code that is hard to read. This decision improves over it and forces a cleaner and readable approach over that.

}
