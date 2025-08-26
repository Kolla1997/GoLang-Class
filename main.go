package main

import (
	"GoLang_Class/Section1"
	"GoLang_Class/Section2"
	"GoLang_Class/Section4"
	"GoLang_Class/Tasks"

	"fmt"
)

func main() {
	var temp float32
	var day int
	var limit int
	fmt.Scanln("Please enter the day number (1-7):", &day)
	fmt.Println("Hello World") // Day 1 Task
	Section1.Variables()
	Tasks.Task_day2() // Day 2 Task
	fmt.Print("Please enter the temperature:")
	if _, err := fmt.Scanln(&temp); err == nil {
		Section4.If_else(temp)
	} else {
		fmt.Println("Error in reading the temperature:", err)
	}
	fmt.Print("Please enter the day number (1-7):")
	if _, err := fmt.Scanln(&day); err == nil {
		Section4.Switch_case(day)
	} else {
		fmt.Println("Error in reading the day number:", err)
	}
	fmt.Print("Please eneter a number to all numbers upto that number:")
	if _, err := fmt.Scanln(&limit); err == nil {
		Section4.ControlFlow(limit)
	} else {
		fmt.Println("Error in reading the prime number:", err)
	}

	// calling function with return values
	fmt.Println("Addition with return value:", Section2.Add(10, 20))
	fmt.Println("Subtraction with return value:", Section2.Sub(20, 10))
	fmt.Println("Multiplication with return value:", Section2.Mult(15, 20))

}
