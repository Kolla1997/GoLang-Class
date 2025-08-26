package Section4

import "fmt"

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
}
