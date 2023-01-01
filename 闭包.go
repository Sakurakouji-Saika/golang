package main

func bb_cal(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}

	sub := func(a int) int {
		base -= a
		return base
	}
	return add, sub

}

func bb_add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	f := bb_add()
	r := f(10)
	print("10 : ", r, "\n")
	r = f(20)
	print("20 : ", r, "\n")
	r = f(30)
	print("30 : ", r, "\n")

	print("--------------\n")

	f1 := bb_add()
	r1 := f1(100)
	print("r1-f1 = 100 : ", r1, "\n")

	r1 = f1(200)
	print("r1-f1 = 200 : ", r1, "\n")

	//案例2
	addbb, subbb := bb_cal(100)
	r2 := addbb(100)
	print("案例2 r addbb: ", r2, "\n")
	r2 = subbb(50)
	print("案例2 r subbb: ", r2, "\n")

}
