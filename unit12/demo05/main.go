package main
import (
	"fmt"
)

func main(){
	var intChan2 chan  int  //管道具备只写性质
	intChan2 = make(chan int,3)
	intChan2 <- 20
	num := <- intChan2
	fmt.Println(intChan2,num)
}