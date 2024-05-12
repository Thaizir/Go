package main

import "fmt"

func main() {

	// maps are key/value pairs, similar to objects in JS and dictionaries in Python but we have to declare the type of the key and the value
	// in this case the key is a string and the value is a float64
	menu := map[string]float64{
		"pie":     5.99,
		"cake":    3.99,
		"donut":   2.99,
		"salad":   8.99,
		"muffin":  4.99,
		"cupcake": 2.99,
	}

	fmt.Println(menu)
	// we can use the key to get the value
	fmt.Println(menu["pie"]) // 5.99

	// loop through the map

	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as keys types

	phonebook := map[int]string{
		123456789: "Link",
		987654321: "Zelda",
		123456783: "Sheik",
		987654322: "Ganondorf",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[123456789])

	// Update values in the map
	phonebook[123456789] = "Link Prime"
	fmt.Println(phonebook)

	// add a new value into the map
	phonebook[000000000] = "Ruto"
	fmt.Println(phonebook)

	// delete a value in the map, we use the delete function to do that delete(mapName,keyToDelete)
	delete(phonebook, 000000000)
	fmt.Println(phonebook)
}
