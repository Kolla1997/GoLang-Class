package Section2

import (
	"fmt"
	"time"
)

func If_else(temp float32) {
	if temp < 30.0 {
		fmt.Println("Todays weather is super cold!")
	} else if temp > 30.0 && temp <= 50.0 {
		fmt.Println("Todays weather is cold!")
	} else if temp > 50.0 && temp <= 70.0 {
		fmt.Println("Todays weather is moderate!")
	} else {
		fmt.Println("Todays weather is hot!")
	}

	fmt.Println("calling short if-else function")
	shortif_else()
}

func shortif_else() {
	if num := 5; num%2 == 0 {
		fmt.Println(num, " is even number")
	} else {
		fmt.Println(num, " is odd number")
	}
}

func Switch_case(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Please eneter the valid day number (1-7)")
	}
}

func ControlFlow(limit int) {
	for i := 1; i <= limit; i++ {
		if i%6 == 0 {
			fmt.Print("\n", i, " ")
		} else {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	fmt.Println("Calling while loop built using for loop")
	whileLoop()
	fmt.Println("Calling nested for-loop built using for loop")
	nested_for_loop()
}

func whileLoop() { // while loop in golang
	i := 0
	for i <= 10 {
		if i == 5 {
			fmt.Println("Breaking loop at i=5")
			break // break the loop when i=5 and come out of the loop
		}
		time.Sleep(1 * time.Second) // added time delay 1 second unit we see the output
		i++
	}
}

func nested_for_loop() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
