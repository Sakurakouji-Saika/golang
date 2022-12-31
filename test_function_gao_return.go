package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	ff := cal("+")
	r := ff(1, 2)
	fmt.Printf("r :%v\n", r)
	ff = cal("-")
	l := ff(1, 2)
	fmt.Printf("l :%v\n", l)
}
