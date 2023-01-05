package main

import "fmt"

func 结构体指针() {
	type Person struct {
		id   int
		name string
		age  int
	}

	tom := Person{
		id:   101,
		name: "tom",
		age:  18,
	}

	var p_person *Person
	p_person = &tom

	fmt.Printf("tom:%v\n", tom)
	fmt.Printf("p_person_地址:%p\n", p_person)
	fmt.Printf("p_person_指针内容:%v\n", *p_person)
}

func 结构体指针_new实例化() {
	type Person struct {
		id   int
		name string
		age  int
	}

	var p_person = new(Person)
	fmt.Printf("----------------\n")
	fmt.Printf("p_person:%T\n", p_person)

	p_person.id = 1
	p_person.name = "昵称"
	p_person.age = 18

	fmt.Printf("p_person-value:%v\n", p_person)
}
func main() {
	结构体指针()
	结构体指针_new实例化()
}
