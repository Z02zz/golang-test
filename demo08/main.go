package main
import "fmt"

func main(){
	var s1 string = "111111111"
	fmt.Println(s1)
	var s2 string = "222222"
	fmt.Println(s2)
	//反引号``有特殊字符
	var s3 string = `
	package main
	import "fmt"

	func main(){
	var flag01 bool = true
	fmt.Println(flag01)

	var flag02 bool = false
	fmt.Println(flag02)

	var flag03 bool = 5<8
	fmt.Println(flag03)
}
`
fmt.Println(s3)
//字符串拼接
var s5 string = "abc" + "def"
s5 += "hijk"
fmt.Println(s5)

//+号保留在上一行最后
var s6 string = "abc" +"abc" +"abc" +"abc" +"abc" +
"abc" +"abc" +"abc" +"abc" +"abc" +"abc" +
"abc" 
fmt.Println(s6)
}