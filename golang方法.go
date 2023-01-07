package main

import "fmt"

type person_方法 struct {
	name string
}

func (per person_方法) eat() {
	fmt.Println(per.name + " eating... ")
}

func (per person_方法) sleep() {
	fmt.Println(per.name + " sleep... ")
}

type Customer_1 struct {
	name string
}

func (customer Customer_1) login_11(name string, pwd string) bool {
	fmt.Printf("Customer.name : %v \n", customer.name)

	if name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}

}

func main() {
	// var per person_方法
	// per.name = "tom"
	// per.eat()
	// per.sleep()

	cus := Customer_1{
		name: "tom",
	}

	b := cus.login_11("tom", "123")

	fmt.Print(b)
}
