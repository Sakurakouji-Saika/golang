package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func f1(a int) {
	a = 100
}

func f2(s []int) {
	s[0] = 114514
}

func f3(args ...int) {
	for _, v := range args {
		fmt.Println("v:", v)
	}
}

func f4(name string, ok bool, args ...int) {
	fmt.Println("name:", name)
	fmt.Println("ok:", ok)
	for _, v := range args {
		fmt.Println("v:", v)
	}
}

func main() {
	fmt.Println(sum(2, 3))

	x := 200
	f1(x)
	fmt.Println(x)

	s := []int{1, 2, 3}
	f2(s)
	fmt.Println(s)

	f3(1, 2, 3)
	f3(4, 5, 6, 7, 8)

	f4("测试昵称", true, 9, 8, 7, 6, 5)
}
