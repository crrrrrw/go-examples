package main

import "fmt"

func main() {
	x := 0
	// 匿名方法,定义一个方法处理变量
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
