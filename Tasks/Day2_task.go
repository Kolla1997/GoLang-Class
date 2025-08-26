package Tasks

import "fmt"

func add_cal(x, y int) {
	fmt.Println("Addition is: ", x+y)
}

func product_cal(x, y int) {
	fmt.Println("Product is: ", x-y)
}

func Task_day2() {
	var x, y int
	fmt.Print("Please provide X value:")
	fmt.Scan(&x)
	fmt.Print("Please provide Y value:")
	fmt.Scan(&y)
	add_cal(x, y)
	product_cal(x, y)
}
