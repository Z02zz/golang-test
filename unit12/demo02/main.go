package main
import(
	"fmt"
	"time"
)

func main(){
	for i := 1;i <= 5;i++{
		//启动一个协程
		//使用了匿名函数，直接调用
		go func(){
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second * 2)

}