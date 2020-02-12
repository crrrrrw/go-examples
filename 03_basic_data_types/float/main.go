package main

import "fmt"

/**
1.浮点数=符号位+指数位+尾数位；
2.存储过程可能精度损失
3.浮点类型有固定的范围和字段长度，不受OS影响
4.浮点类型默认为float64
5.两种表示形式：
	- 十进制形式
	- 科学计数法形式
**/
func main() {

	var price float32 = 28.23
	fmt.Printf("price is %f \n", price)
	var price2 float64 = 292.232
	fmt.Printf("price2 is %f \n", price2)

	// 精度损失
	var num1 float32 = -123.12345678
	var num2 float64 = -123.12345678
	fmt.Printf("num1 is %f , num2 is %f\n", num1, num2)

	// 默认 float64 类型
	var num3 = 1.1
	fmt.Printf("num3 数据类型为 %T \n", num3)

	// 十进制形式
	num4 := 5.23
	num5 := .123
	fmt.Printf("num4 is %f , num5 is %f\n", num4, num5)

	// 科学计数法形式
	num6 := 1.1234e2  // => 1.1234 * 10^2
	num7 := 1.1234e-2 // => 1.1234 * 10^-2
	fmt.Printf("num6 is %f \n", num6)
	fmt.Printf("num7 is %f \n", num7)
}
