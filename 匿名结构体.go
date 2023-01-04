package main

import "fmt"

func main() {
	var dog struct {
		id   int
		name string
	}

	dog.id = 1
	dog.name = "花花"
	fmt.Printf("dog:%v\n", dog)
}
