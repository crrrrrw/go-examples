package main

import "fmt"

/**
指针的使用基本与c语言一致.
指针类型指向的值类型，二者必须类型一致。
值类型包括：int,float,bool,string,数组,结构体。
**/
func main() {
	var num int32 = 123
	fmt.Printf("num is %d, 地址是 %v\n", num, &num) // num is 123, 地址是 0xc42001807c

	var p *int32 = &num
	fmt.Printf("指针 p的值=%v，p的地址=%v，p指向的值=%v\n", p, &p, *p) // 指针 p的值=0xc42001807c，p的地址=0xc42000c030，p指向的值=123
}
