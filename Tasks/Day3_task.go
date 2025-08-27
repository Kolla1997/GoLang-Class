package Tasks

import "fmt"

func if_else(x, y int) {
	fmt.Println("In If-Else function")
	if x < y {
		fmt.Println(x, " is less than ", y)
	} else if x == y {
		fmt.Println(x, " is equal to ", y)
	} else {
		fmt.Println(x, " is greater than ", y)
	}
}

func forLoop(x, y int) {
	a, b := x, y
	if x > y {
		a, b = y, x
	}
	fmt.Println("In For Loop function")
	for i := a; i <= b; i++ {
		for j := a; j <= b; j++ {
			fmt.Println(i, "X", j, "=", i*j)
		}
	}
	fmt.Println()
}

func Task_day3() {
	var x, y int
	fmt.Println("Day 2 Task")
	fmt.Println()
	fmt.Print("Please provide X value: ")
	if _, err := fmt.Scanln(&x); err == nil {
		fmt.Println("X value is: ", x)
	} else {
		fmt.Println("Error in reading X value: ", err)
	}
	fmt.Print("Please provide Y value: ")
	if _, err := fmt.Scanln(&y); err == nil {
		fmt.Println("Y value is: ", y)
	} else {
		fmt.Println("Error in reading Y value: ", err)
	}

	if_else(x, y)
	forLoop(x, y)
}
