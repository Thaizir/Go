package main

import "fmt"

func main() {

	// x := 0

	// there's no while loop in Go

	// for x <= 5 {
	// 	fmt.Println("the value of x is: ", x)
	// 	x++ // adds 1 to x
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("the value of x is: ", i)
	// }

	// loop through an slice (like a for in)
	names := []string{"Link", "Zelda", "Sheik", "Ganondorf"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// This form is similar to the for in in JavaScript
	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	// If we don't want to use the index we can substitute it with the "_"
	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}

	fmt.Println(names)

}
