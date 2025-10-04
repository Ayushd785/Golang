// Every executable Go program must start with 'package main'.
package main

// The 'import' statement lists the packages this file uses.
// We need 'fmt' for printing and 'math/rand' for random numbers.
import (
	"fmt"
	"math/rand"
)

func main() {
	// We call functions from imported packages using 'packageName.FunctionName()'.
	// Here, we print a string and a random integer up to 9.
	fmt.Println("My favorite number is", rand.Intn(10))
}