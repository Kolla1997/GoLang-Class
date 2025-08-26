package Section2

import "fmt"

func add(x, y int) {
	fmt.Println("Addition is: ", x+y)
}

func product(x, y int) {
	fmt.Println("Product is: ", x-y)
}

func Task_sec2() {
	var x, y int
	fmt.Print("Please provide X value:")
	fmt.Scan(&x)
	fmt.Print("Please provide Y value:")
	fmt.Scan(&y)
	add(x, y)
	product(x, y)
}
