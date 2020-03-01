package main

// Import the format package to print to console.
import "fmt"

// Function that adds two integers together.
func add(x int, y int) int {
	return x + y
}

func main() {
	// Prints a string and then the result of adding two plus two.
	fmt.Printf("Two plus two is %d!\n", add(2, 2))
}
