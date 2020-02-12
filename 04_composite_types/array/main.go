package main

import "fmt"

/**
数组定义和结构和其他语言差不多
1. 长度固定，占用连续内存，各元素类型相同
2. var arr []int，这时 arr 就是一个slice切片。
3. 数组创建好后，若没有赋值，则有默认值
	- 数值型默认为 0
	- bool型默认为 false
	- string型默认为 ""
4. golang 的数组属于值类型，和java不同，它不是引用类型。这一点非常特殊！！！
5. 数组长度也是数据类型的一部分。比如，int[3] 与 int[4] 不是同一类型。
**/
func main() {

	// 定义数组
	var a [6]int
	// 若未赋值，则数据为默认值
	fmt.Println(a) // [0 0 0 0 0 0]
	// 遍历赋值
	for i := 0; i < len(a); i++ {
		a[i] = i * i
	}
	fmt.Println(a)

	// 输出索引和元素
	for i, v := range a {
		fmt.Printf("i:%d, v:%d\n", i, v)
	}

	// 仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 声明数组并初始化
	var q [4]int = [4]int{6, 7, 8, 9}
	fmt.Println(q) // [6 7 8 9]

	// 使用 ... 在数组长度位置，数组长度可由数组元素个数决定
	r := [...]int{1, 3, 5, 7, 9}
	fmt.Printf("%v, len:%d, %T \n", r, len(r), r) // [1 3 5 7 9], len:5, [5]int

	var nums [5]float64
	var sum float64
	var avg float64
	// 终端输入 5 个数并计算其平均值
	for i := 0; i < 5; i++ {
		fmt.Printf("请输入第 %d 个数\n", i)
		fmt.Scanln(&nums[i])
	}
	for _, v := range nums {
		sum += v
	}
	avg = sum / float64(len(nums))
	fmt.Printf("sum:%f, avg:%f \n", sum, avg)

	// 初始化数据并指定下标
	names := [...]string{1: "Tom", 0: "Jerry", 2: "Jane"}
	fmt.Println(names) // [Jerry Tom Jane]

	// 测试数组值传递
	arr := [3]int{1, 2, 3}
	test(arr)
	fmt.Println(arr) // [1 2 3]

	test2(&arr)
	fmt.Println(arr) // [88 2 3]
}

func test(arr [3]int) {
	arr[0] = 6
}

// 传入指针类型数组，通过引用传递改原数组的值
func test2(arr *[3]int) {
	(*arr)[0] = 88
}
