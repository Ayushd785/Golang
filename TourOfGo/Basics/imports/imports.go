package main

// This is a "factored" import statement. It's the standard Go
// style for importing multiple packages cleanly.
import (
	"fmt"
	"math"
)

func main() {
	// fmt.Printf lets us format strings.
	// The verb '%g' is replaced by math.Sqrt(7), and '\n' adds a new line.
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}