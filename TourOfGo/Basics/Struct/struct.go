package main

import (
	"fmt"
)

type vertex struct{
	x int
	y int
}

func main(){
	fmt.Println("welcome")

	tt:= vertex{1,2}

	x := tt.x
	y := tt.y

	fmt.Println(tt)
	fmt.Println(x)
	fmt.Println(y)

	p := &tt

	p.x = 23
	fmt.Println(tt)
}

