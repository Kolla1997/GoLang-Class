package Section1

import "fmt"

func Variables() {
	a := 10                    // short hand declaration
	var b = 20                 // explicit declaration
	var name string = "Dinesh" // explicit declaration with type
	age := 25
	fmt.Println("Name:", name, "age:", age, "sum of a,b:", a+b)
}
