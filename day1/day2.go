package main

import "fmt"

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14

	//%T操作变量所属类型
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)

	// %d 对应整形格式输出
	// %s 对应字符串格式输出
	// %c 对应字符格式输出
	// %f 对应浮点格式输出
	fmt.Printf("a = %d, b = %s, c = %c, d = %f\n", a, b, c, d)

	//%v自动匹配格式输出
	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)

}
