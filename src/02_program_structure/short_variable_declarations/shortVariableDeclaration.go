package main

import (
	"fmt" 
)

/*
name := expression
短小，灵活，非常适合在用在局部变量声明以及初始化
*/
func main() {
	i := 100 // an int
	var boiling float64 = 100 // a float64
	var names []string
	var err error
	fmt.Println(i,boiling,names,err)

	// 声明多个短变量
	a, b:= 0, 1
	fmt.Println(a, b)

	// :=表示声明 =表示赋值
	a, b = b, a// 交换 a, b
	fmt.Println(a, b)

	print()
}

func print() {
	fmt.Println("准备输入一些短变量")
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", d, d)
	fmt.Printf("%v, %T \n", e, e)
	fmt.Printf("%v, %T \n", f, f)
	fmt.Printf("%v, %T \n", g, g)
}