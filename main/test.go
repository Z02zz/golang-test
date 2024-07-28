package main	//申明所在包，每个go文件必须有归属的包	
import "fmt"	//引入程序中需要用的包，为了使用包下的函数 比如Println

func main(){	//main主函数 程序的入口
	fmt.Println("Hello Golang!")	//在控制台打印一句话
}