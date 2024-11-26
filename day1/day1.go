package main

import "fmt"

//常量
const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)


func main() {
	var a int = 10
	fmt.Printf("12313")
	fmt.Println(a)
	//匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("n4=", n4)
	fmt.Println("PB=",PB)
}

func foo() (int, string) {
	return 10, "Q1mi"
}
