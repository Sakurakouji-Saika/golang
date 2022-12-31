package main

import "fmt"

func test1() {
	fmt.Println("我既没有参数，也没有返回值")
}

func sum(a int, b int) int {
	return a + b
}

func test2() (name string, age int, sex string) {
	name = "昵称"
	age = 18
	sex = "Man"
	return name, age, sex
}
func test3() (name string, age int, sex string) {
	name = "test3"
	age = 19
	sex = "女"
	return
}

func test4() (string, int) {
	n := "test4"
	a := 18

	return n, a
}

func main() {
	test1()
	fmt.Println(sum(4, 6))

	name, age, sex := test2()
	name1, age1, sex1 := test3()
	name2, age2 := test4()
	fmt.Printf("%v,%v,%v\n", name, age, sex)
	fmt.Printf("%v,%v,%v\n", name1, age1, sex1)
	fmt.Printf("%v,%v\n", name2, age2)
}
