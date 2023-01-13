package main

import "fmt"

type Pet_3 interface {
	eat()
	sleep()
	// 定义了一个接口有吃和睡方法
}

type Dog_2 struct {
	name string
	age  int
	// 定义了一个狗类型，有名字和年龄
}

func (dog Dog_2) eat() {
	fmt.Println("dog eat...")
	//关于狗的接口实现内容，吃
}

func (dog Dog_2) sleep() {
	fmt.Println("dog sleep...")
	//关于狗的接口实现内容，睡觉
}

type Cat struct {
	name string
	age  int
}

func (dog Cat) eat() {
	fmt.Println("Cat eat...")
}

func (dog Cat) sleep() {
	fmt.Println("Cat sleep...")
}

type Person_3 struct {
	name string
}

func (Per Person_3) care(per Pet_3) {
	per.eat()
	per.sleep()
}

func main() {
	dog := Dog_2{}
	cat := Cat{}
	per := Person_3{}

	per.care(dog)
	per.care(cat)
}
