package main

import(
	"fmt"
)

func main(){
	fmt.Println("welcome")

	nums:= []int{1,2,3,4,5}

	fmt.Println(nums)
	
	r:= []bool{true, false, true, false}

	fmt.Println(r)

	tumtum :=[]pairs{
		{2,true},
		{3,false},
		{4,true},
		{5, false},

	}

	fmt.Println(tumtum)
	
}

type pairs struct{
	x int
	r bool
}