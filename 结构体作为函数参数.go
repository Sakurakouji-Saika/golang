package main

import "fmt"

type Person_1 struct {
	id   int
	name string
}

// 值传递拷贝了一份副本
func showPerson(person Person_1) {
	person.id = 1
	person.name = "like"
	fmt.Printf("showPerson-person:%v\n", person)
}

// 函数内如果传递的是指针，则可以在函数内修改函数外的参数
func showPerson_指针(person *Person_1) {
	person.id = 1
	person.name = "like"
	fmt.Printf("showPerson-person:%v\n", person)
}

func main() {
	person := Person_1{1, "tom"}
	fmt.Printf("person:%v\n", person)
	fmt.Print("---------------\n")
	showPerson(person)
	fmt.Print("---------------\n")
	fmt.Printf("person - 最后: %v\n", person)
	fmt.Print("---------------\n")

	showPerson_指针(&person)
	fmt.Printf("指针修改后 person:%v\n", person)
	fmt.Print("---------------\n")

}
