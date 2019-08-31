package main

import "fmt"

/*
var name type = expression
*/
func main() {
	// 声明单个变量
	var s string
	fmt.Println(s) // ""
	// 声明变量列表，三个 int 类型
	var i, j, k int // int, int, int
	// 声明 bool, float64, string 类型，并赋值
	var b, f, str = true, 2.3, "four" // bool, float64, string

	fmt.Println(i, j, k) // 0 0 0
	fmt.Println(b, f, str) // true 2.3 four
}
