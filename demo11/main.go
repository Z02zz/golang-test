package main
import "fmt"
import "strconv"

func main(){
	var n1 int = 10
	// var n2 float32 = 4.78
	// var n3 bool = false
	// var n4 byte = 'a'
	
	var s1 string = fmt.Sprintf("%d",n1)
	fmt.Printf("%T,%v",s1,s1)

	var n2 int = 100
	var s2 string = strconv.FormatInt(int64(n2),10)
	fmt.Printf("%T,%v",s2,s2)

	var n3 float32 = 3.14
	var s3 string = strconv.FormatFloat(float64(n3),'f',9,32)
	fmt.Printf("%T,%v",s3,s3)

	var n4 bool = true
	var s4 string = strconv.FormatBool(n4)
	fmt.Printf("%T,%v",s4,s4)

	var s5 string = "true"
	var b bool
	b , _ = strconv.ParseBool(s5)
	fmt.Printf("%T,%v",b,b)
}