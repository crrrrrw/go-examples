package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 65535
	var i8 int8 = 1<<7 - 1    // 1 字节 -128~127
	var i16 int16 = 1<<15 - 1 // 2 字节
	var i32 int32 = 1<<31 - 1 // 4 字节
	var i64 int64 = 1<<63 - 1 // 8 字节

	fmt.Print(a, i8, i16, i32, i64, "\n") // 65535 127 32767 2147483647 9223372036854775807
	fmt.Print(i8, i8+1, i8*i8, "\n")      // 127 -128 1

	var ui8 uint8 = 1<<8 - 1
	var ui16 uint16 = 1<<16 - 1
	var ui32 uint32 = 1<<32 - 1
	var ui64 uint64 = 1<<64 - 1

	fmt.Print(ui8, ui16, ui32, ui64, "\n") // 255 65535 4294967295 18446744073709551615
	fmt.Print(ui8, ui8+1, ui8*ui8, "\n")   // 55 0 1

	var n1 = 100
	fmt.Printf("n1 的类型: %T \n", n1) // 整形数据默认为 int 类型

	var n2 int64 = 10
	fmt.Printf("n1 的类型: %T ，占用字节数: %d \n", n2, unsafe.Sizeof(n2))

}
