package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func cmp(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	type f1 func(int, int) int
	var ff f1
	ff = sum

	my_Out_Sum := ff(6, 3)
	fmt.Println("my_Out_Sum:", my_Out_Sum)

	ff = cmp
	my_Out_CMP := ff(114, 514)
	fmt.Println("my_Out_CMP:", my_Out_CMP)
}
