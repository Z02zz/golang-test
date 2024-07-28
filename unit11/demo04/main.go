package main
import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	file,err := os.OpenFile("d:/demo.txt",os.O_RDWR|os.O_APPEND|os.O_CREATE,0666)

	if err != nil{
		fmt.Println("打开文件失败",err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < 10;i++{
		writer.WriteString("hello\n")
	}

	writer.Flush()
}