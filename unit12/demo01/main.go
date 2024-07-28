package main
import(
	"fmt"
	"time"
	"strconv"
)

func test(){
	for i := 1;i <= 10;i++{
		fmt.Println("hello golang +" + strconv.Itoa(i))
		time.Sleep(time.Second*2)
	}
}

func main(){//主线程
	go test()//开启一个协程

	for i := 1;i <= 10;i++{
		fmt.Println("hello zzz + " + strconv.Itoa(i))
		//阻塞一秒
		time.Sleep(time.Second)
	}
}