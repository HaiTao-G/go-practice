package main

import "fmt"

func test() {
	a := 10
	fmt.Println("a =", a)
}

/*
作用域：就是变量产生作用的范围。
定义在{}里面的变量就是局部变量，同时也只在{}里面有效。
执行到定义变量那句话，才开始给变量分配空间，离开作用域之后自动释放
*/

func main() {
	fmt.Println("a =", a)

	{
		i := 10
		fmt.Println("i = ", i)
	}
	fmt.Println("i = ", i)
	if b := 3; b == 3 {
		fmt.Println("b = ", b)
	}
	fmt.Println("b = ", b)
}
