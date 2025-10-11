package main

import (
	"fmt"
)

func main(){
	fmt.Println("welcome")

	primes:= [6]int{1,2,3,4,5,6}

	var s[]int = primes[1:4];
	fmt.Println(s)

	names:= [4]string{
		"john",
		"paul",
		"george",
		"meloni",
	}

	fmt.Println(names)
	a:= names[0:2]
	b:= names[1:3]

	fmt.Println(a,b)

	b[0] = "xxxx"

	fmt.Println(a,b)

	//A slice does not store any data, it just describes a section of an underlying array.

    //Changing the elements of a slice modifies the corresponding elements of its underlying array.

    // Other slices that share the same underlying array will see those changes.

	

}