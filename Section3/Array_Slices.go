package Section3

import (
	"fmt"
	"reflect"
)

func Arrays() {
	var numbs [5]int // declaring an empty array of integers with size 5
	for i := 0; i < len(numbs); i++ {
		numbs[i] = i + 1
	}
	fmt.Println(numbs)

	var names [5]string = [5]string{"Dinesh", "kolla", "Golang", "Class", "Hello"}
	fmt.Println(names)

	var marks = [5]float32{56.4, 76.8, 45.0, 55, 12.67}
	fmt.Println(marks)

	if reflect.TypeOf(names) == reflect.TypeOf(marks) {
		fmt.Println("Both are of same type")
	} else {
		fmt.Println("Both are of different types")
	}
	for i, val := range names {
		fmt.Println(i, "-", val)
	}
}

// Slices in Go
func Slices() {
	var numbs [5]int = [5]int{10, 20, 30, 40, 50}
	slice := numbs[:3]
	fmt.Println("Slice:", slice)
	fmt.Println("Length of the slice:", len(slice))
	fmt.Println("Capacity of the slice:", cap(slice))
	// appending values to a slice
	slice = append(slice, 90, 78, 23, 45, 67)
	fmt.Println("Slice after append:", slice)
	fmt.Println("Length of the slice after append:", len(slice))
	fmt.Println("Capacity of the slice after append:", cap(slice))

	// creating a slice using make function
	slice2 := make([]int, 5, 10) // creates a slice of int with length 5 and capacity 10
	copy(slice2, slice)          // copying values from slice to slice2
	fmt.Println("Slice2:", slice2)
	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2))

	// Remove an element from slice
	i := 2
	slice2 = append(slice2[:i], slice2[i+1:]...) // removing the 3rd element
	fmt.Println("Slice2 after removing 3rd element:", slice2)
}
