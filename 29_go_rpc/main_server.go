package main

import (
	"go-learning/29_go_rpc/server"
	"log"
	"net"
	"net/rpc"
)

/*
go rpc 概念梳理：
	RPC:
		概念：远程过程调用
		核心：通信协议(TCP HTTP2 Grpc)、序列化（JSON、ProtoBuf）
		(RPC经常携带者服务注册，治理，监控等)
		使用：
			内置的net/rpc包来实现
*/
// main 服务端
func main() {
	//注册服务
	rpc.RegisterName("MathService", new(server.MathService))
	listen, err := net.Listen("TCP", "1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	rpc.Accept(listen)
}
