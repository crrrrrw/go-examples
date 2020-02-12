package main

import "fmt"

/**
1. golang 字符串统一使用UTF-8编码
2. 字符串不可变，不可修改。
3. 字符串的两种表达形式：
	- 双引号 ""
	- 反引号 ``
4. 字符串拼接方式
**/
func main() {
	var name string = "Tom"
	fmt.Println("name is ", name)
	fmt.Println("name is ", name)

	// 字符串不可修改
	var str1 string = "Hello"
	// str1[0]=’w‘ // cannot assign to str1[0]
	fmt.Println("str1 is ", str1)

	// 普通，双引号表示
	str2 := "abc\n123"
	fmt.Println("str2 is ", str2)

	// 反引号表示，以原生形式展示
	str3 := `
	func main() {
		var name string = "Tom"
		fmt.Println("name is ", name)
		fmt.Println("name is ", name)
	}
	`
	fmt.Println("str3 is ", str3)

	// 字符串拼接
	var str4 = "hello," + "world"
	fmt.Println("str4 is ", str4)
}
