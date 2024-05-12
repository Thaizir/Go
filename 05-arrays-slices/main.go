package main

import "fmt"

func main() {

	// To declare an array we have to set the length of the array and the type of the elements in the array, which means we cannot change the lenght of the array by adding of removing elements. We use curly braces to create the array
	// var ages [3]int = [3]int{20, 30, 33}
	var ages = [3]int{20, 30, 33} // Go infers the type

	// len is the method we use to get the lenght of the array
	fmt.Println(ages, len(ages))

	names := [4]string{"Link", "Zelda", "Sheik", "Ganondorf"}
	fmt.Println(names, len(names))

	names[1] = "Pepito"
	fmt.Println(names, len(names))

	// Slices: are more like typical arrays in other languages (use arrays under the hood)
	var scores = []int{100, 50, 6, 8, 63, 75}
	scores[2] = -1
	fmt.Println(scores)

	// Append items to slices
	// append doens't change the original slice, it returns a new slice for us. If we want to edit the original slice we need to make the old slice equals to the append structure
	scores = append(scores, 85)
	fmt.Println(scores)

	// slice ranges
	rangeOne := names[1:3]  // Just take the values between the 1st and 3rd element in the array including the position 1 but not the position 3.
	rangeTwo := names[2:]   // Takes all the elements from Kira (including Kira) until the end of the array
	rangeThree := names[:3] // includes all the elements until the position 3 not including the position 3.

	// We can append elements into the rages
	rangeOne = append(rangeOne, "Salmen")

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)
}
