package main

import (
	"GoLang_Class/Section1"
	"GoLang_Class/Section4"

	"fmt"
)

func main() {
	var temp float32
	var day int
	var limit int
	fmt.Scanln("Please enter the day number (1-7):", &day)
	fmt.Println("Hello World")
	Section1.Variables()
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

}
