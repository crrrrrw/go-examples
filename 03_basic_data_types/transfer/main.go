package main

import (
	"fmt"
	"strconv"
)

/**
基本数据类型的相互转换
1. 不同类型变量之间赋值时必须显示转换，不支持自动转换。这个与java、c等不同。
2. 数据类型转换范围可以由大到小，也可以有小到大。
3. 数据转换大转小时，编译不会报错，可能发生数据溢出问题

基本数据类型 -> string类型
1. fmt.Sprintf
2. strconv包

string类型 -> 基本数据类型
1. 必须确保数据有效，否则将会转化为一个默认值

**/
func main() {
	var n int16 = 100
	// var f float64 = n //cannot use n (type int) as type float64 in assignment
	var f float64 = float64(n)
	fmt.Println("f is", f)

	var n1 int32 = int32(n)
	var n2 int8 = int8(n)
	fmt.Printf("n1 is %d, n2 is %d\n", n1, n2)
	fmt.Printf("n 数据类型为 %T\n", n) // 原数据类型不变

	// 数据范围由大到小时，如果超出了类型范围，不会编译报错，但会做溢出处理，结果会变
	var n3 int16 = 255
	var n4 int8 = int8(n3)
	fmt.Printf("n3 is %d, n4 is %d\n", n3, n4) // n3 is 255, n4 is -1

	/**
	// n3 + 10的数据类型 = n3的数据类型
	var n5 int32
	n5 = n3 + 10 // cannot use n3 + 10 (type int16) as type int32 in assignment
	**/

	// 基本数据类型 -> string类型
	// 方式一 fmt.Sprintf
	var num1 int = 100
	var num2 float64 = 45.2
	var b bool = true
	var c byte = 'c'
	var str string
	str = fmt.Sprintf("%d , %f, %t, %c", num1, num2, b, c)
	fmt.Printf("str is %q, 类型是 %T \n", str, str) // str is "100 , 45.200000, true, c", 类型是 string

	// 方式二 strconv 包
	var str2 string
	str2 = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str2 is %q, 类型是 %T \n", str2, str2) // str2 is "100", 类型是 string
	str2 = strconv.FormatFloat(num2, 'f', 3, 64)
	fmt.Printf("str2 is %q, 类型是 %T \n", str2, str2) // str2 is "45.200", 类型是 string
	str2 = strconv.FormatBool(b)
	fmt.Printf("str2 is %q, 类型是 %T \n", str2, str2) // str2 is "true", 类型是 string
	str2 = strconv.Itoa(num1)
	fmt.Printf("str2 is %q, 类型是 %T \n", str2, str2) // str2 is "100", 类型是 string

	// string类型 -> 基本数据类型
	var str3 = "true"
	var b1 bool
	b1, _ = strconv.ParseBool(str3)
	fmt.Printf("b1 is %v, 类型是 %T \n", b1, b1) // b1 is true, 类型是 bool

	var str4 = "123"
	var num3 int64
	num3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("num3 is %v, 类型是 %T \n", num3, num3) // num3 is 123, 类型是 int64

	var str5 = "123.4567"
	var f1 float64
	f1, _ = strconv.ParseFloat(str5, 64)
	fmt.Printf("f1 is %v, 类型是 %T \n", f1, f1) // f1 is 123.4567, 类型是 float64

	f1, _ = strconv.ParseFloat("hello", 64)
	fmt.Printf("f1 is %v, 类型是 %T \n", f1, f1) // f1 is 0, 类型是 float64
}
