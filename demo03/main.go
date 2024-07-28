package main
// import "fmt"
// import "unsafe"
import(
	"fmt"
	"unsafe"
)
func main(){
	var num1 int8 = 120
	fmt.Println(num1)

	var num3 byte= 28
	fmt.Printf("%T",num3)
	fmt.Println(unsafe.Sizeof(num3))
}