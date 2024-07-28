package main

import(
	"fmt"
	"os"
	"bufio"
	"io"
)

func main(){

	//打开文件
	file, err := os.Open("d:/text01.txt")
	if err != nil {
		fmt.Println("文件打开失败,err",err)
	}

	defer file.Close()

	//创建一个流
	reader := bufio.NewReader(file)

	//读取操作
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取操作成功，已全部读取完毕")
}