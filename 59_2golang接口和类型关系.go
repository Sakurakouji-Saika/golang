// // 多个类型实现一个接口
package main

import "fmt"

type Pet_2 interface {
	eat()
}

type Dog_3 struct {
}

type Cat_3 struct {
}

func (dog Dog_3) eat() {
	fmt.Println("dog eat...")
}

func (cat Cat_3) eat() {
	fmt.Println("cat eat...")
}

func main() {

	dog := Dog_3{}
	cat := Cat_3{}
	dog.eat()
	cat.eat()

	//实现多态
	var p Pet_2
	p = Cat_3{}
	p.eat()
	p = Dog_3{}
	p.eat()
}
