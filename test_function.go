package main

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func comp(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}
func main() {
	println(sum(1, 2))

	r := comp(2, 114514)
	println(r)

}
