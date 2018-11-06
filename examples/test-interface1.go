package main

import "fmt"

func main() {
	var x *int = nil
	var y interface{} = x
	fmt.Printf("x == nil ? %+v\n", x == nil)
	fmt.Printf("y == nil ? %+v\n", y == nil)
}
