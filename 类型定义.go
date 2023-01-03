package main

import "fmt"

func main() {
	type MyInt int
	var i MyInt
	i = 100
	fmt.Printf("i-value: %v \t,\t i-type: %T\n", i, i)
}
