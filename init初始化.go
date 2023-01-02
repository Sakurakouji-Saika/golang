package main

import "fmt"

var i int = initVar()

func init() {
	fmt.Print("init1... \n")
}

func init() {
	fmt.Print("init2... \n")
}

func initVar() int {
	fmt.Print("initVar \n")
	return 0
}

func main() {
	fmt.Print("main \n")
}
