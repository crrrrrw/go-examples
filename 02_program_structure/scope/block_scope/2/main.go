package main

import "fmt"

// 包变量，包内都可以访问的到
var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
