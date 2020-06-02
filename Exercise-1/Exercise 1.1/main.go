// In this exercise, print a random number, between 1 and 5, of stars (*) to the console.
// Create a new folder and add a main.go file to it.
// In main.go, add the main package name to the top of the file:
package main

// Now, add the imports we'll use in this file:
import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Create a main() function:
func main() {
	// Seed the random number generator:
	rand.Seed(time.Now().UnixNano())
	// Generate a random number between 0 and 4 then add 1 to get a number between 1 and 5:
	r := rand.Intn(5) + 1
	// Use the string repeater to create a string with the number of stars we need:
	stars := strings.Repeat("*", r)
	// Print the string with the stars to the console with a new line character at the end and close the main() function:
	fmt.Println(stars)
}
