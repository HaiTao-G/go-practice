package main

import (
	"fmt"
)

func main() {
	var falg bool
	falg = true
	fmt.Printf("flag = %t\n", falg)

	//bool类型不能转换为int
	//fmt.Printf("flag = %d\n",int(falg))

	// 0就是假，非0就是真
	// 整型也不能转换为bool
	// flag = bool(1)

	var ch byte
	ch = 'a'
    t := int(ch)
    fmt.Println("t = ",t)
    fmt.Print("t = \n",t)
    fmt.Printf("t =",t)
}
