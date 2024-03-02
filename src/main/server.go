package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("server: accept error", err)
		}
		go server.serveConn(conn)
	}
}

func (server *Server) serveConn(conn io.ReadWriteCloser) {
	server.readRequest(conn)
	server.sendRequest(conn, "hello")
}

func (server *Server) readRequest(conn io.ReadWriteCloser) {
	var content string
	json.NewDecoder(conn).Decode(&content)
	log.Println("server: read msg: ", content)
}

func (server *Server) sendRequest(conn io.ReadWriteCloser, reply string) {
	json.NewEncoder(conn).Encode(&reply)
}
