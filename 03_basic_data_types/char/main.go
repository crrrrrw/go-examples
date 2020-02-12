package main

import "fmt"

/**
1. golang 没有专门的字符类型，若要存储单个字符，一般用 byte 类型，超出byte数据范围的用较大的数据类型比如int。
2. golang 的字符串是由单个字节连接起来的。
3. golang 的字符使用UTF-8编码
4. golang 字符串存储到计算机过程:
	存储：字符->ASCII码->二进制
	读取: 二进制->ASCII码->字符
**/
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1 is ", c1)
	fmt.Println("c2 is ", c2)
	fmt.Printf("c1 is %c, c2 is %c \n", c1, c2)

	// var c3 byte = '中'// constant 20013 overflows byte
	var c3 int = '中'
	fmt.Printf("c3 is %c , c3 is %d \n", c3, c3)

	var n = 100 + 'a'
	fmt.Println("n is ", n)
}
