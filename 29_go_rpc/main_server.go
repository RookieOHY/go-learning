package main

import (
	"go-learning/29_go_rpc/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
go rpc 概念梳理：
	RPC:
		概念：远程过程调用
		核心：通信协议(TCP HTTP2 Grpc)、序列化（JSON、ProtoBuf）
		(RPC经常携带者服务注册，治理，监控等)
		使用：
			内置的net/rpc包来实现
	将对象注册为rpc服务端的规则：
		服务端方法是可以导出的
		方法的类型可导出
		方法必须有2个参数（参数的类型可导出，要么是自建）
		方法的返回是error
	定义格式：
		func (t *T) MethodName(argType T1, replyType *T2) error
		参数argType由客户端提供；后者由服务端返回（一定是指针类型）

*/
// main 服务端(基于TCP、HTTP协议的RPC服务端)
func main001() {
	//注册服务(服务名+具体的服务对象)
	rpc.RegisterName("MathService", new(server.MathService))
	//处理http
	rpc.HandleHTTP()
	//在端口1234开发Tcp链接
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	//rpc.Accept(listen)
	http.Serve(listen, nil)
}

/*
②基于json的tcp rpc服务端
*/
func main() {
	rpc.RegisterName("MathService", new(server.MathService))
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error", e)
	}
	for {
		conn, error := l.Accept()
		if error != nil {
			log.Println("jsonrpc.Serve: accept:", error.Error())
			return
		}
		go jsonrpc.ServeConn(conn)
	}
}
