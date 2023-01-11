// 一个类型实现多个接口
package main

import "fmt"

type Pet interface {
	eat(string) string
}
type CatM struct {
	name string
}

// func (Cat CatM) eat(name string) string {
func (Cat *CatM) eat(name string) string {
	Cat.name = "花花....\n"
	fmt.Printf("name-eat : %v \n", name)
	return "吃的很好\n"
}

func main() {
	// cat := CatM{
	cat := &CatM{
		name: "画画...\n",
	}

	s := cat.eat("哗哗...\n")

	fmt.Printf("s: %v \n", s)
	fmt.Printf("dog: %v \n", cat)
}
