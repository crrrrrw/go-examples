package main

import "fmt"

// 指针。语法看着跟C一致
func main()  {
	x := 1
	p := &x // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2 // equivalent to x = 2
	fmt.Println(x) // "2"

	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil) // "true false false"

	fmt.Println(f() == f()) // "false" 每次返回一定不同

	v := 1
	incr(&v) // v+1 = 2
	fmt.Println(incr(&v)) // "3" (v=3)
}

// 返回一个局部变量的地址
func f() *int {
	v := 1
	return &v
}

// 入参是一个指针参数，增加指针参数所指向的变量的值
func incr(p *int) int {
	*p++ // 增加指针参数所指向的变量的值;并没有改变p的值
	return *p
}