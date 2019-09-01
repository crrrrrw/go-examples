package main

import "fmt"

/*
 “=” 用于赋值， 与其他语言一致
*/
func main() {
	var x int
	x = 1          // 有名称的变量
	fmt.Println(x) // 1

	p := &x
	*p = 2         // 非直接赋值
	fmt.Println(x) // 2

	// 自增自减，与其他语言一样
	v := 1
	v++
	fmt.Println(v) // 2
	v--
	fmt.Println(v) // 1

	a := 1
	b := 2
	a, b = b, a       // 多行赋值，此句交换a,b
	fmt.Println(a, b) // 2 1

	fmt.Println(fib(10)) // 55

}

// 斐波那契数列
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
