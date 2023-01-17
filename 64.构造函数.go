package main

import "fmt"

type Person_64 struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person_64, error) {
	if name == "" {
		return nil, fmt.Errorf("name不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("Age不能小于0")
	}
	return &Person_64{name: name, age: age}, nil
}

func main() {
	person, err := NewPerson("Tom", -20)

	if err == nil {
		fmt.Println("Person:", *person)
	} else {
		fmt.Println("erro:", err)
	}
}
