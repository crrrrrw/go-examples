package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "y变量在{}内部"
		fmt.Println(y)
	}
	// fmt.Println(y) // 编译错误 undefined: y， y在作用域外了
	fmt.Println(x) // x可以访问
}
