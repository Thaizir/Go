package main

import "fmt"

// I can't declare variables using := outside the main() function.
// var someName := "Hello"

func main() {

	// // strings
	// var nameOne string = "Link"
	// var nameTwo = "Zelda"
	// var nameThree string // variable declared to being used in the future

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "Mario"
	// nameThree = "Princess"
	// fmt.Println(nameOne, nameTwo, nameThree)

	// // Another way to initialize variables using colon and equal sign
	// nameFour := "Ganondorf"

	// fmt.Println(nameFour)

	// Integers
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // bits and memory
	// // int8 between -127 and 127
	// var numOne int8 = 126
	// var numTwo int64 = -25678541236

	// //unsing int (just positive numbers)
	// var numThree uint8 = 126 //or uint64 for example

	// fmt.Println(numOne, numTwo, numThree)

	// floats

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 888.545684684648648 // used in the most since has better aproximation than float32
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
