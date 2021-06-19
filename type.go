package main

import "fmt"
import "strconv"

func main() {
	var i int = 1
	var f float32 = 1.1
	fmt.Printf("%T %T\n", i, f)
	fmt.Printf("%v %f\n", i, f)

	i2 := int(f)
	fmt.Printf("%T %v\n", i2, i2)

	var s string = "100"
	ans, err := strconv.Atoi(s)
	fmt.Printf("%T %v %T\n", ans, ans, err)

}
