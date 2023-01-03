package main

import "fmt"

func main() {
	var my_ip *int
	fmt.Printf("ip-value: %v \n", my_ip) //nil
	fmt.Printf("ip-type: %T  \n", my_ip) //nil

	i := 100
	my_ip = &i
	fmt.Printf("ip-取地址: %v \n", my_ip)    //nil
	fmt.Printf("ip-取指针的值: %v \n", *my_ip) //nil

	var sp *string
	var s string = "hello"
	sp = &s
	fmt.Printf("sp-value : %v\n", *sp)
	fmt.Printf("sp-type : %T\n", sp)
}
