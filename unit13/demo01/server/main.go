package main

import(
	"fmt"
	"net"
)

func process(conn net.Conn){
	defer conn.Close()

	for{

		buf := make([]byte,1024)

		n,err := conn.Read(buf)
		if err!= nil{
			return
		}
		fmt.Println(string(buf[0:n]))
	}
}

func main(){
	fmt.Println("服务器端启动了。。")
	listen,err := net.Listen("tcp","127.0.0.1:8888")
	if err != nil{
		fmt.Println("监听失败,err",err)
		return
	}
	for{
		conn,err2 := listen.Accept()
		if err2 != nil{
			fmt.Println("客户端等待失败,err2",err2)
		}else{
			//连接成功
			fmt.Printf("等待连接成功,con=%v,接收到的客户端信息：%v \n",conn,conn.RemoteAddr().String())
		}
			//准备一个协程，协程处理客户端服务请求：
			go process(conn)
	}

}