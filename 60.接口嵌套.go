package main

import "fmt"

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

// 组合接口
type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
}

func (fish Fish) fly() {
	fmt.Printf("fly...\n")
}

func (fish Fish) swim() {
	fmt.Printf("swim...\n")
}

func main() {
	v := Fish{}

	v.fly()
	v.swim()
}
