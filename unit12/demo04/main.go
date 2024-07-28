package main
import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func writeData(intChan chan int){
	defer wg.Done()
	for i:=1;i<=10;i++{
		intChan <- i
		fmt.Println("写入数据为：",i)
		time.Sleep(time.Second*3)
	}
	close(intChan)
}

func readData(intChan chan int){
	defer wg.Done()
	//遍历
	for v := range intChan{
		fmt.Println("读取操作为：",v)
		time.Sleep(time.Second)
	}
}

func main(){//主线程
	//定义管道
	intChan := make(chan int,10)

	wg.Add(2)

	go writeData(intChan)
	go readData(intChan)

	wg.Wait()
}