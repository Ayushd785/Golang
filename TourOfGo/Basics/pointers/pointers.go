package main

import (
	"fmt"
)

func main(){
	fmt.Println("welcome")

	i := 44;
	p:= &i;
	fmt.Println(p); // this tells the ram locatons of I where p points to it
	fmt.Println(*p) // this is actually what the location contains 

	j := 90
	p = &j

	*p = *p/2;  // ----> what it actually does it edits the j 
	fmt.Println(j); 
	fmt.Println(*p)
	// pointers 
}