package main

import (
	"fmt"
)

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat...")
}

func (a Animal) sleep() {
	fmt.Println("sleep...")
}

type Dog_63 struct {
	a_var    Animal //可以理解为继承
	color_63 string
}

type Cat_63 struct {
	a_var    Animal
	color_63 string
}

func main() {
	dog := Dog_63{
		a_var: Animal{
			name: "dog",
			age:  2,
		},
		color_63: "金色",
	}

	cat := Cat_63{
		a_var: Animal{
			name: "Cat",
			age:  5,
		},
		color_63: "白色",
	}

	dog.a_var.eat()
	dog.a_var.sleep()

	fmt.Println("dog.clor:", dog.color_63)
	fmt.Println("dog.a_var.age:", dog.a_var.age)

	cat.a_var.eat()
	cat.a_var.sleep()

}
