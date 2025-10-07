package main

import (
	"fmt"
	"math"
)

func sqrt(f float64) string {
	if f<0 {
		return sqrt(-f) + "i";
	}

	return fmt.Sprint(math.Sqrt(f))
}

func main(){
	ans := sqrt(32)
	fmt.Printf("ans is %v \n", ans)

	isvalid := isAdult(12);
	fmt.Print(isvalid)
}


func isAdult(age int) string{
	if age<18{
		return "not valid"
	}
	return "adult"
}