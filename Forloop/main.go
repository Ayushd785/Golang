package main

import (
	"fmt"
)

func main(){
	for i := 0; i <= 10; i++ {
		if(i == 2){
			continue;
		}
		fmt.Println(i)
	}

	age :=3
	if(age>19){
		fmt.Println("nice you are adult")

	}else{
		fmt.Println("get out come after", 18-age, "years")
	}


	

}


