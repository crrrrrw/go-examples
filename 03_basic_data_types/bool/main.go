package main

import (
	"fmt"
	"unsafe"
)

/**
1. 占用1字节
2. 取值只能为 true or false,不能用 1，0
**/
func main() {
	var b = false
	fmt.Println("b = ", b)
	fmt.Println("b sizeof ", unsafe.Sizeof(b))

	// var b1 bool = 1 // cannot use 1 (type int) as type bool in assignment
}
