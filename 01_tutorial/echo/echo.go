package main
// 单行注释
import (
	"fmt" 
	"os"
)

func main(){
	var s,sep string
	for i:=1;i<len(os.Args);i++ { // for循环，不用带小括号
		s+=sep+os.Args[i] // os.Args是命令行参数
		sep=" "
	} 
	fmt.Println(s)

	for { // 无限循环
		fmt.Println("in loop")
		fmt.Println("out loop")
		break;
	}

	for true { // while 循环
		fmt.Println("in loop")
		fmt.Println("out loop")
		break;
	}
}
