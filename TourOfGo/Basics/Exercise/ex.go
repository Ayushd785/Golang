package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println("welcome")

	var ans float64= sqrt(1069.0);
	fmt.Println(ans)
}

func sqrt(x float64) float64 {
	var z float64 = x

	for i:=0;i<10;i++{
		newz := z-  (z*z -x)/(2*z);
		if math.Abs(newz-z) < 1e-5{
			return newz;
		}

		z = newz;
		fmt.Println(z)

	}
	return z;

}