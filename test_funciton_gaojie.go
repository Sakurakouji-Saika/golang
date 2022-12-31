package main

import "fmt"

//高级函数

func sayJello(name string) {
	fmt.Printf("Hello ,%s", name)
}

func test(name string, f func(string)) {
	//将接收到的形参1，2合并起来
	f(name)
	//f函数是sayJello，函数内的参数使用 test 的 name。
}

func main() {
	// sayJello 指定函数名
	test("tom", sayJello)
}
