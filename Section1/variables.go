package Section1

import "fmt"

func Variables() {
	// in go when a variable is declared and not initialized it takes the zero value of that type
	// for int it's 0, for float it's 0.0
	// for string it's "" (a empty string)
	// for bool  the default value is false
	// for pointers the default value is nil
	// ----------------------------------

	a := 10 // short hand declaration
	var b = 20
	var x, y int = 20, 5       // explicit declaration
	var name string = "Dinesh" // explicit declaration with type
	age := 25
	fmt.Println("Name:", name, "age:", age, "sum of a,b:", a+b)
	fmt.Println("Multiplication of x, y is:", x*y)

	var Fname string

	if Fname == "" {
		fmt.Println("Please provide your first name!")
	} else {
		fmt.Println("The provided first name is: ", Fname)
	}
}
