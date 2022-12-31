package main

func test1_1() {
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	r := max(1, 2)

	print(r)
}

func test1_2() {
	//自己调用自己
	r := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(3, 4)
	print("\n", r)
}

func test1_3() {
	name := "tom"
	age := "20"

	r := func() string {
		return name + " " + age
	}

	msg := r()

	print("\ntest1_3:", msg)

}
func main() {
	test1_1()
	test1_2()
	test1_3()
}
