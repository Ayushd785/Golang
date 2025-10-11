package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	fmt.Println("welcome to the go")

	os := runtime.GOOS;

	switch os{
	case "tt":
		fmt.Println("your os is t")

	case "linux":
		fmt.Println("your os is linux")

	default:
		fmt.Printf("%s. \n", os)
	}


	fmt.Println("second switch example")
	fmt.Println("when is saturday?")

	today:= time.Now().Weekday();
	switch time.Saturday{
	case today + 0:
		fmt.Println("today")
	case today +1:
		fmt.Println("tomorrow")
	case today +2:
		fmt.Println("day after tomorrow")
	default:
		fmt.Println("too far")
	}


	
	

}

