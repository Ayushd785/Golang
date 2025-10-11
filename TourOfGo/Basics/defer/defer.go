package main

import (
	"fmt"
)

func main(){
		// This 'defer' statement schedules fmt.Println("world") to run later,
	// just before the main function exits.

	defer fmt.Println("world")

	
	// This line executes immediately.
	fmt.Println("hello")




	// stacking defer

	fmt.Println("counting")

	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// Execution Order:
// 1. fmt.Println("hello") is called, printing "hello".
// 2. The main function is now finished and about to return.
// 3. The deferred call, fmt.Println("world"), is now executed, printing "world".