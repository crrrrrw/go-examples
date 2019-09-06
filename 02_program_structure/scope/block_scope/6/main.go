package main

import "fmt"

var x string
var y string
var z string

func init() {
	x := "hello"                // := 语句将x声明为局部变量,而全局变量x未初始化
	fmt.Println("init() x=", x) // init() x= hello

	var y = "world"             // y 也是局部变量
	fmt.Println("init() y=", y) // init() z= world

	z = "nihao"                 // z 全局变量赋值
	fmt.Println("init() z=", z) // init() z= nihao
}

func main() {
	fmt.Println("main() x=", x) // main() x=
	fmt.Println("main() y=", y) // main() y=
	fmt.Println("init() z=", z) // init() z= nihao
}
