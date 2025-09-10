package Section3

import "fmt"

// Structs in Go
type Person struct {
	Name string
	Age  int
	DOB  struct {
		Day   int
		Month int
		Year  int
	}
	Degree      string
	Citizanship string
	Is_underage bool
}

func Structs() {
	students := make(map[int]Person)
	for i := 0; i < 2; i++ {
		var name, degree, born string
		var age, day, month, year int
		fmt.Print("Enter your Name:")
		fmt.Scanln(&name)
		fmt.Print("Enter your Degree:")
		fmt.Scanln(&degree)
		fmt.Print("Enter your placer of birth:")
		fmt.Scanln(&born)
		fmt.Print("Enter your Age:")
		fmt.Scanln(&age)
		fmt.Print("Enter your DOB Day:")
		fmt.Scanln(&day)
		fmt.Print("Enter your DOB Month:")
		fmt.Scanln(&month)
		fmt.Print("Enter your DOB Year:")
		fmt.Scanln(&year)
		students[i] = Person{
			Name:        name,
			Degree:      degree,
			Citizanship: born,
			Age:         age,
			DOB: struct {
				Day   int
				Month int
				Year  int
			}{
				Day:   day,
				Month: month,
				Year:  year,
			},
		}
		fmt.Println("Student added successfully!")
	}
	fmt.Printf("Students Data: %v", students)
}
