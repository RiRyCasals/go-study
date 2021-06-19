package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%3v", i*j)
		}
		fmt.Println()
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		} else if i == 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
}
