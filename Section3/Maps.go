package Section3

import "fmt"

func Maps() {
	// declaring and initializing a map
	student := map[string]int{
		"Dinesh": 25,
		"Ravi":   23,
		"Sita":   22,
	}
	fmt.Println("Student Map:", student)

	// adding a new key-value pair to the map
	student["Gita"] = 24
	fmt.Println("Student Map after adding Gita:", student)

	// updating a value in the map
	student["Ravi"] = 24
	fmt.Println("Student Map after updating Ravi's age:", student)

	// deleting a key-value pair from the map
	delete(student, "Sita")
	fmt.Println("Student Map after deleting Sita:", student)

	// accessing a value from the map
	age := student["Dinesh"]
	fmt.Println("Dinesh's age is:", age)

	// checking if a key exists in the map
	if val, ok := student["Sita"]; ok {
		fmt.Println("Sita's age is:", val)
	} else {
		fmt.Println("Sita is not found in the map")
	}

	// iterating over a map
	for name, age := range student {
		fmt.Println(name, "is", age, "years old")
	}
}
