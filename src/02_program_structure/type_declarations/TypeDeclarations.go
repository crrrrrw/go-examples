package main

import "fmt"

/*
语法：
type name underlying-type
*/
// 以下定义 Celsius 和 Fahrenheit 两种类型，底层类型都是 float64 型
type Celsius float64
type Fahrenheit float64

// 定义常量：绝对0度，0度，沸点
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC)       // 编译报错：类型不匹配

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	// fmt.Println(c == f)          // 编译报错：类型不匹配
	fmt.Println(c == Celsius(f)) // "true"! ; 这种情况未改变f的值，只是把类型改变了
}

// 设摄氏度转华氏度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// 华氏度转摄氏度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
