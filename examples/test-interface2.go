package main

import "fmt"

func main() {
	x := "test"
	var y interface{} = x
	fmt.Printf("%+v\n", y)
}
