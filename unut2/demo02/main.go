package main
import "fmt"

//全局变量
var n7 = 100
var n8 = 9.7

//一次性申明全局变量
var(
	n9= 1000
	n10 = "name"
)

func main(){
	//{}局部变量
	var num int = 18
	fmt.Println(num)

	//2.不赋值，使用默认值打印	
	var num2 int
	fmt.Println(num2)

	//3.没写变量类型，会自动推断类型
	var num3 = "tom"
	fmt.Println(num3)

	//4.省略var关键字，注意 := 不能写为 = 
	sex := "男"
	fmt.Println(sex)

	fmt.Println("-----------------")

	var n1,n2,n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4,name,n5 = 10,"jack",7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

	n6,height := 6.9,100.6
	fmt.Println(n6)
	fmt.Println(height)

	fmt.Println(n7)
	fmt.Println(n8)

	fmt.Println(n9)
	fmt.Println(n10)
}