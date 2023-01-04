package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	id    int
	name  string
	age   int
	email string
}

type Customer struct {
	id, age     int
	name, email string
}

func main() {
	var tom Person

	tom.id = 10
	tom.name = "tom"
	tom.age = 20
	tom.email = "admin@qq.com"
	fmt.Printf("tom:%v\n", tom)

	var Nums [3]Customer

	for i := 0; i < 3; i++ {
		Nums[i].id = i
		Nums[i].name = "name:" + strconv.Itoa(i)
		Nums[i].email = "admin@" + strconv.Itoa(i)
		Nums[i].age = i * 10
	}

	for _, v := range Nums {
		fmt.Printf("id:%v\n", v.id)
		fmt.Printf("name:%v\n", v.name)
		fmt.Printf("email:%v\n", v.email)
		fmt.Printf("age:%v\n\n", v.age)
	}

}
