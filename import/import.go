package main

/*
 Import format package for printing to console.
 Import math random package for generating a random integer.
*/
import (
	"fmt"
	"math/rand"
)

func main() {
	// Prints a string and random integer to console.
	fmt.Println("Printing random number: ", rand.Intn(10))
}
