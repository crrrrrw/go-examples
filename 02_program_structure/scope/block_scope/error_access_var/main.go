package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// 此处访问不了 x 变量，x是局部变量
	fmt.Println(x) // 编译报错： undefined: x
}
