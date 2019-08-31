package main

import "fmt"

// 指针。语法看着跟C一致
func main()  {
	p := new(int) // p是*int类型, 指向一个未命名的 int 变量
	fmt.Println(p) // 返回地址
	fmt.Println(*p) // "0"
	*p = 2 // 设置 p 指向的变量的值为2
	fmt.Println(*p) // "2"
}
