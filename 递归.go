package main

// 阶层
func jc(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * (n - 1)
	}
}

func main() {
	n := 5
	r := jc(n)
	print(r)
}
