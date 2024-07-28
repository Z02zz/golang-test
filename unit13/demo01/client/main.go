package main

import(
	"fmt"
	"net"
	"bufio"
	"os"
)



func main()  {
	fmt.Println("客户端启动。。")

	//调用Dial函数
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil{
		fmt.Println("客户端连接失败")
		return
	}
	fmt.Println("连接成功conn:",conn)

	reader := bufio.NewReader(os.Stdin)

	str,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("终端输入失败,err",err)
	}

	n,err2 := conn.Write([]byte(str))

	if err2 != nil{
		fmt.Println("err2",err2)
	}
	fmt.Printf("%d",n)

	
}