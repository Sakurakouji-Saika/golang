package main

import "fmt"

func main() {
	fmt.Print("1\n")
	defer fmt.Print("2\n")
	defer fmt.Print("3\n")
	defer fmt.Print("4\n")
	fmt.Print("5\n")
}
