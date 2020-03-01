package main

// Import the format package for printing to console.
import "fmt"

// Swaps the order of two given strings.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// Assigns the variables a and b to the results of the swap function.
	a, b := swap("that", "this")

	// Prints the result of the swap function by printin a and b in order.
	fmt.Println(a, b)
}
