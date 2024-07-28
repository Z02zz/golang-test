package main
import(
	"fmt"
	"io/ioutil"
)

func main(){
	//在下面不需要open，close，封装在Readfile文件内
	content,err := ioutil.ReadFile("d:/text01.txt")

	if err != nil {//读取有误	
		fmt.Println("读取有误，错误为",err)
	}
	//读取成功
	fmt.Printf("%v",string(content))
}

