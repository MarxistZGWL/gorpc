package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
)

type Client struct {
	conn io.ReadWriteCloser
}

func NewClient() *Client {
	return &Client{}
}

var DefaultClient = NewClient()

func (client *Client) receive() {
	var err error
	for err == nil {
		var reply string
		json.NewDecoder(client.conn).Decode(&reply)
		log.Println("client: receive", reply)
	}
}

func (client *Client) send(content string) {
	log.Println("client: send", content)
	json.NewEncoder(client.conn).Encode(&content)
}

func Dial(addr string) *Client {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("client: dial error", err)
		return nil
	}
	client := &Client{
		conn: conn,
	}
	go client.receive()
	return client
}
