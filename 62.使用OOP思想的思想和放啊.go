package main

import "fmt"

type Person_oop struct {
	name string
	age  int
}

func (per Person_oop) eat() {
	fmt.Println("eat...")
}

func (per Person_oop) sleep() {
	fmt.Println("sleep...")
}

func (per Person_oop) work() {
	fmt.Println("work...")
}
func main() {
	per := Person_oop{
		name: "tom",
		age:  20,
	}

	fmt.Printf("per : %v\n", per)
	per.eat()
	per.sleep()
	per.work()
}
