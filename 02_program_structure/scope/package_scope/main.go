package main

import (
	"fmt"
	"go-examples/02_program_structure/scope/package_scope/visiblity"
)

func main() {
	fmt.Println(visiblity.Name)
	// fmt.Println(visiblity.age) // 名称小写， 编译会报错
	visiblity.PrintInfo()
}
