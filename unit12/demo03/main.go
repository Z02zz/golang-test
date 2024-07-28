package main
import(
	"fmt"
	"sync"
)

var wg sync.WaitGroup //只定义

func main(){
	for i := 1; i<=5; i++{
		wg.Add(1)
		go func(n int){
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	//主线程阻塞
	wg.Wait()
}