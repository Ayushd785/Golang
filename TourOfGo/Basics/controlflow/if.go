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
	fmt.Println()

	anspow := pow(3,2,10)
	fmt.Print(anspow)
}


func isAdult(age int) string{
	if age<18{
		return "not valid"
	}
	return "adult"
}

// this is a very powerfull feature of Go in which you can declare short statements in the if condition only which has the scope of just if and else block only 

func pow( x, n, lim float64) float64{
	if v:=math.Pow(x,n); v<lim{
		return v
	}
	return lim
}