package main

import (
	"fmt"
	"go-learning/29_go_rpc/server"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// main 客户端（基于TCP和HTTP的RPC客户端）
func main01() {
	//client, err := rpc.Dial("tcp", "localhost:1234")
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//入参
	args := server.Args{A: 7, B: 8}
	var reply int
	//服务端返回指针类型，也需要用指针类型接收
	err = client.Call("MathService.Add", args, &reply)
	if err != nil {
		log.Fatal("MathService.Add error:", err)
	}
	//打印结果
	fmt.Printf("MathService.Add: %d+%d=%d", args.A, args.B, reply)
}

/*
	基于json的tcp客户端
*/
func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//入参
	args := server.Args{A: 7, B: 8}
	var reply int
	//服务端返回指针类型，也需要用指针类型接收
	err = client.Call("MathService.Add", args, &reply)
	if err != nil {
		log.Fatal("MathService.Add error:", err)
	}
	//打印结果
	fmt.Printf("MathService.Add: %d+%d=%d", args.A, args.B, reply)
}
