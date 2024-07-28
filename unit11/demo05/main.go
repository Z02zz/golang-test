package main
import (
	"fmt"
	"io/ioutil"
)

func main(){
	//第一源文件
	file1Path := "d:/demo.txt"
	//目标文件
	file2Path := "d:/demo02.txt"

	//对文件读取
	content,err := ioutil.ReadFile(file1Path)
	if err != nil{
		fmt.Println("err")
		return
	}

	err = ioutil.WriteFile(file2Path,content,0666)
	if err != nil{
		fmt.Println("failed")
	}
}