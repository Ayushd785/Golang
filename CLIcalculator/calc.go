package main

import "fmt"

func main(){
	for {
		fmt.Println("select options between 1 to 4")
		var option int
		fmt.Scan(&option)
		if option == 5 {
			fmt.Println("Thank you for playing")
			break
		}
		if option > 4 || option < 1 {
			println("please enter a valid option")
			return
		}
		var b float64
		var a float64
		fmt.Println("enter a:")
		fmt.Scan(&a)
		fmt.Println("enter b:")
		fmt.Scan(&b)

		var ans float64
		switch option {
		case 1:
			ans = add(a,b);
		case 2:
			ans = subt(a,b);
		case 3:
			ans = multiply(a,b);
		case 4:
			ans = divide(a,b);
		}

		fmt.Println("ans is:", ans)
	}

}

func add(a float64, b float64) float64 {
	return a + b
}

func subt(a float64, b float64) float64 {
	return a-b
}

func multiply(a float64, b float64) float64 {
	return a*b
}

func divide(a float64, b float64) float64 {
	return a/b
}







