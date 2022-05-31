package main

import (
	"fmt"
	"math"
)

func fgcd(a, b float64) float64 {
	if a < b {
		return fgcd(b, a)
	}
	if math.Abs(b) < 0.001 {
		return a
	} else {
		return (fgcd(b, a-math.Floor(a/b)*b))
	}

}

func main() {
	var a float64 = 0
	var b float64 = 0
	fmt.Print("Enter number1: \n")
	n, err := fmt.Scan(&a)
	if n != 1 || err != nil {
		fmt.Print("Not a number \n")
	} else {
		fmt.Print("Enter number2: \n")
		n, err = fmt.Scan(&b)
		if n != 1 || err != nil {
			fmt.Print("Not a number \n")
		} else {
			f := fgcd(math.Abs(a), math.Abs(b))
			fmt.Printf("Highest Common Factor is: ")
			fmt.Println(f)
		}
	}
}
