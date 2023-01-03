package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{1, 3, 5}
	var ptr [MAX]*int
	fmt.Println(ptr) //打印 [<nil> <nil> <nil>]
	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i]
		//  整数地址赋值给指针数组
	}

	fmt.Println(ptr)

	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
