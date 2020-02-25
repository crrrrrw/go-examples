package main

import "fmt"

/**
1. 切片是数组的一个引用，是引用类型，因此是引用传递。
2. 切片使用类似数组，遍历、访问、求长度都跟数组一样。
3. 切片的长度是动态变化，类似java上的ArrayList
4. slice底层结构体
	type slice struct {
		ptr *[2]int // 指针数组指向引用数组的地址
		len int // 长度
		cap int	// 容量
	}
5. 声明一个切片
	- 方式1：定义一个切片，让切片引用一个已创建好的数组
	- 方式2：通过 make 来创建切片： var name []type = make([]type, len, [cap])
	- 二者不同点：方式1可以使用slice和arr都改变数组元素；
				但是方式2 make创建的切片，隐式的创建了一个数组，但是这个数组只能slice访问，make方法底层维护，对外不可见
	- 方式3：定义一个切片，直接指定具体的数组，原理类似于make方式
**/
func main() {

	// 数组
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	// 方式1：声明一个切片名为 slice1;引用了数组 intArr;引用 intArr 下标 1到3，左闭右开 [1,3)
	slice1 := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Printf("slice1=%v, 长度=%d, 容量=%d\n", slice1, len(slice1), cap(slice1))

	// 内存地址分析
	fmt.Printf("intArr[1]地址=%p,值=%d\n", &intArr[1], intArr[1])
	fmt.Printf("slice1[0]地址=%p,值=%d\n", &slice1[0], slice1[0])
	slice1[0] = 66
	fmt.Printf("intArr[1]地址=%p,值=%d\n", &intArr[1], intArr[1])
	fmt.Printf("slice1[0]地址=%p,值=%d\n", &slice1[0], slice1[0])

	// 方式2：make 创建切片
	var slice2 []float64 = make([]float64, 5, 10)
	fmt.Printf("slice2=%v, 长度=%d, 容量=%d\n", slice2, len(slice2), cap(slice2)) // 初始化会有默认值

	// 方式3
	var slice3 []string = []string{"aa", "bb", "cc"}
	fmt.Printf("slice3=%v, 长度=%d, 容量=%d\n", slice3, len(slice3), cap(slice2))
}
