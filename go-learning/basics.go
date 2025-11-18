package main

import "fmt"

// demonstrating variables, types, and basic operations
func main() {
	// Variable declarations
	var name string = "Go Programmer"
	var age int = 25
	var isLearning bool = true

	// Short variable declaration
	course := "Go Programming"
	progress := 75.5

	// Constants
	const language = "Go"

	fmt.Println("=== Go Basics ===")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Learning: %t\n", isLearning)
	fmt.Printf("Course: %s\n", course)
	fmt.Printf("Progress: %.1f%%\n", progress)
	fmt.Printf("Language: %s\n", language)

	// Basic arithmetic
	fmt.Println("\n=== Arithmetic ===")
	a, b := 10, 3
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)

	// Calling functions
	fmt.Println("\n=== Functions ===")
	sum := add(5, 7)
	fmt.Printf("5 + 7 = %d\n", sum)

	product, quotient := calculate(20, 4)
	fmt.Printf("20 * 4 = %d, 20 / 4 = %d\n", product, quotient)

	// Arrays and slices
	fmt.Println("\n=== Arrays and Slices ===")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)
	fmt.Println("First element:", numbers[0])
	fmt.Println("Length:", len(numbers))

	// Append to slice
	numbers = append(numbers, 6, 7, 8)
	fmt.Println("After append:", numbers)

	// Maps
	fmt.Println("\n=== Maps ===")
	scores := map[string]int{
		"Math":    95,
		"Science": 88,
		"English": 92,
	}
	fmt.Println("Scores:", scores)
	fmt.Println("Math score:", scores["Math"])

	// Control flow
	fmt.Println("\n=== Control Flow ===")
	grade := getGrade(scores["Math"])
	fmt.Printf("Grade for Math: %s\n", grade)

	// Loops
	fmt.Println("\n=== Loops ===")
	fmt.Print("Counting: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// Function with return value
func add(x, y int) int {
	return x + y
}

// Function with multiple return values
func calculate(x, y int) (int, int) {
	return x * y, x / y
}

// Function with if-else
func getGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}
