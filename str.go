package main

import "fmt"
import "strings"

func main() {
	fmt.Println("Hello", "world")

	fmt.Println(`Hello
    world
    and
go`)

	fmt.Println("\"")
	fmt.Println(`"`)

	s := "Hello world"
	fmt.Println(s[0])
	fmt.Println(string(s[0]))
	fmt.Println(s[:3])

	fmt.Println(strings.Contains(s, "H"))
	fmt.Println(strings.Contains(s, "a"))
}
