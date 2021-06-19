package main

import "fmt"

func main() {
	var n int
	fmt.Println("Input number")
	fmt.Scan(&n)

	switch {
	case n < 5:
		fmt.Println("less than 5")
	case n == 5:
		fmt.Println("is 5")
	case n > 5:
		fmt.Println("greater than 5")
	}
}
