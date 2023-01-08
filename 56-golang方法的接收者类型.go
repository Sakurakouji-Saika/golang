package main

import "fmt"

func main() {
	type Person struct {
		name string
	}

	p1 := Person{
		name: "Tom",
	}

	p2 := &Person{
		name: "tom",
	}

	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p1: %T\n", p2)

	//p1 是值类型
	//p2 是指针类型
}
