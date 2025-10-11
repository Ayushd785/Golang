package main

import (
	"fmt"
)

func main(){
	var a[2] string;
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1,2,3,4,5,6}
	primes[1] = 4;
	fmt.Println(primes[1])
	fmt.Println(primes)

}