package main

import (
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Println("dial: error")
	}
	go Accept(lis)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	addr := lis.Addr().String()
	call(addr, wg)
	wg.Wait()
}
func call(addr string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	client := Dial(addr)
	client.send("I am a client")
	time.Sleep(5 * time.Second)
}
