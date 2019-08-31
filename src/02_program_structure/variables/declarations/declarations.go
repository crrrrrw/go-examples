package main

import "fmt"

const a = 3.0 // 声明一个常量

func main() {
	// 声明局部变量
	var b = 5.0
	var c = add(a,b)
	fmt.Printf("a=%f, b=%f, c=%f",a,b,c)
}

// 声明一个函数
func add(a float64, b float64) float64{
	return a + b
}
