package main
import "fmt"

func main(){
	//  \n转行
	fmt.Println("aaa\nbbb")

	// \b退格
	fmt.Println("aaa\bbbb")

	// \r光标回到本行开头，后续输入会替换原有字符
	fmt.Println("aaaaa\rbbb")

	// \t 制表符
	fmt.Println("aaa\tbbb")

	// \"
	fmt.Println("\"aaaaaabbbbbbbbbb\"")
	fmt.Println("'\"aa'aa'aab\\bbbbbbbbb\\")
}