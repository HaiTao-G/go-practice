package main

import "fmt"

func main() {
	a := 10
	str := "mike"

	f1 := func() {
		fmt.Printf("a= %d, str =%s\n", a, str)
	}

	f1()

	type FuncType func()
	var f2 FuncType
	f2 = f1
	f2()

	func ()  {
		fmt.Printf("a = %d,str = %s\n",a,str)
	}()

	
}
