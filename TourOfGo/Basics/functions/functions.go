package main

import (
	"fmt"
)

func calculate(x int, y int) int {
	return x+y
}

// A function can return any number of results

func swap(x,y string)(string,string){
	return y,x
}

func split(sum int)(x,y int){
	x = sum*4/9
	y = sum -x

	return 
}

func main(){
	ans := calculate(5,6)
	fmt.Printf("your answer is %d \n", ans)

	a,b:= swap("hello ","world ")
	fmt.Print(a,b)

	ans2,ans3 := split(32)
	fmt.Printf("your answers are %d and %d", ans2, ans3)
}