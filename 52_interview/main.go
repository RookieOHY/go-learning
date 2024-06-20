package main

import (
	"fmt"
	"sync"
)

type Client struct {
	id string
}

type clientManager struct {
	clientIdMapLock sync.RWMutex
	clients         map[string]*Client
}

func (manager *clientManager) Addclient(client *Client) {
	manager.clientIdMapLock.Lock()
	defer manager.clientIdMapLock.Unlock()
	manager.addclient2Group(client)
}

func (manager *clientManager) addclient2Group(client *Client) {
	manager.clientIdMapLock.Lock()
	defer manager.clientIdMapLock.Unlock()
	fmt.Println(666)
	if manager.clients == nil {
		manager.clients = make(map[string]*Client)
	}
	manager.clients[client.id] = client
}

func main() {
	manager := &clientManager{}

	//// 启动一个goroutine来添加客户端
	//go func() {
	testClient := &Client{id: "client1"}
	manager.Addclient(testClient)
	//}()

	//// 等待一段时间，确保goroutine有机会运行
	//time.Sleep(10 * time.Second)
	//
	//// 检查是否存在死锁情况
	//fmt.Println("Test finished.")
}

func main01() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	chA := make(chan int, 1)
	chB := make(chan int)

	chA <- 1

	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			<-chA
			fmt.Printf("%c", 'a'+i)
			chB <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			<-chB
			fmt.Printf("%d", i+1)
			chA <- 1
		}
	}()

	wg.Wait()
}
