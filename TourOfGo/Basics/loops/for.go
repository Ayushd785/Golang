package main

import (
	"fmt"
)

func main(){
	k := 4
	for i:=0;  i<2;  i++  {
		if k==12{
			break;
		}
		k++;
	}

	sum:=1;

	for ; sum<1000; {
		sum++
	}

	// in Go while is same as for there is no while in go
	// you just need to drop the semicolons to use the while 

	for sum>200 {
		sum--;
	}

	


	fmt.Printf("the K is %d \n" , k)
	fmt.Printf("the sum is %d" , sum)

}