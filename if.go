package main

import "fmt"

func main() {
	n := 10
	if n > 5 {
		fmt.Println("greater than 5")
	} else if n == 5 {
		fmt.Println("is 5")
	} else {
		fmt.Println("less than 5")
	}
}
