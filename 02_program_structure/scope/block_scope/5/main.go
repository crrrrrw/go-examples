package main

import (
	"fmt"
	"os"
)

func main() {

	/*
		if f, err := os.Open("aa.txt"); err != nil { // 编译错误：f declared and not used
			fmt.Println(err)
		}
		f.Stat() // 编译错误: undefined f
		f.Close()    // 编译错误: undefined f
	*/

	// 解决上面问题，必须要在结构体外声明
	f, err := os.Open("aa.txt")
	if err != nil {
		fmt.Println(err)
	}
	f.Stat()
	f.Close()

	// 如果不想在外部块声明，可以放到else块中
	if f, err := os.Open("aa.txt"); err != nil {
		fmt.Println(err)
	} else {
		// f and err are visible here too
		f.Stat()
		f.Close()
	}

}
