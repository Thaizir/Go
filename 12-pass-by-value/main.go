package main

import "fmt"

// Go is Pass By Value language

// That means that when we pass a value to a function, the value is copied to the function and the original value is not changed or affected

// Group A: Strings, Ints, Floats, Booleans, Arrays, Structs
// Group B: Slices, Maps, Functions

func updateName(n string) {
	n = "Link"
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {

	// Group A: Strings, Ints, Floats, Booleans, Arrays, Structs >> Non Pointer Values

	// name := "Zelda"
	// updateName(name)
	// fmt.Println(name) // output: Zelda - because we pass by value and the original name variable "Link" is not affected
	// So for Group A only the copied value is affected
	// A solution for this is using pointers or overwriting the original variable with the new value

	// Group B: Slices, Maps, Functions >> Pointer Wrapper Values

	// We change the original map
	// What happens is when we create the map, Go stores the data in a memory block and also creates a pointer in another memory block that points to the data in the first memory block, so when we make a change in Group B element Go copies the pointer which points to the data in the first memory block and changes the data in the first memory bloc,
	menu := map[string]float64{
		"pie":     5.99,
		"cake":    3.99,
		"cupcake": 2.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
