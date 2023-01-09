package main

import "fmt"

type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

type Mobile struct {
	model string
}

func (m Mobile) read() {
	fmt.Println("m.model:", m.model)
	fmt.Println("read-Mobile...\n")
}

func (m Mobile) write() {
	fmt.Println("m.model:", m.model)
	fmt.Println("write-Mobile...\n")
}

func (c Computer) read() {
	fmt.Println("c.name:", c.name)
	fmt.Println("read...\n")
}

func (c Computer) write() {
	fmt.Println("c.name:", c.name)
	fmt.Println("write...\n")
}

func main() {

	var usb USB //使用定义的接口，必须全部实现,如果少一个就会报错，比如删除write方法在运行，就会报错
	usb = Computer{
		name: "联想",
	}

	usb.read()
	usb.write()
}
