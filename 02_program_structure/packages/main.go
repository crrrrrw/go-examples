package main

import (
	"fmt"
	pack "go-examples/02_program_structure/packages/mylib" // 包路径可以随意，他会扫描 pacage 这行代码的 包名，坑了我好久
)

func main() {
	fmt.Println(pack.Multy(1.5, 23.2))
}
