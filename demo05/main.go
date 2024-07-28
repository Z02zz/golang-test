package main
import "fmt"

func main(){
	var c1 byte = 'a'
	fmt.Println(c1)

	var c2 byte = '6'
	fmt.Println(c2)

	//本质上是一个整数，可以参与直接运算
	var c3 byte = '('
	fmt.Println(c3)

	//汉字字符对应unicode码
	//对应的码值为20013，byte类型溢出，可以用int
	//golang的字符对应使用的是UTF-8，utf-8是unicode其中一种编码方案

	var c5 byte = 'A'
	fmt.Printf("%c",c5)
}