package server

import (
	"encoding/json"
	"fmt"
	"net"
)

type Server struct {
	id   int
	name string
	host string
}

var servers = [...]string{"127.0.0.1:8001", "127.0.0.1:8002", "127.0.0.1:8003"}

func CreateServer(id int, name string, host string) *Server {
	return &Server{id, name, host}
}

func (server *Server) Send(req *Info) {

	for _, s := range servers {
		conn, err := net.Dial("tcp", s)
		//conn, err := net.Dial("tcp", "10.30.8.105:8080")

		if err != nil {
			fmt.Println("Err: ", err)
			return
		}

		encoder := json.NewEncoder(conn)
		e := encoder.Encode(*req)

		fmt.Println("Encode error: ", e)
		conn.Close() // we're finished
	}
}

func (server *Server) Receive(p string) {
	ln, err := net.Listen("tcp", p)
	if err != nil {
		fmt.Println("Can`t create server")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println("Accepted connection")

	d := json.NewDecoder(c)
	var msg Info
	err := d.Decode(&msg)

	fmt.Println("Msg: ", msg, err)

	c.Close()
}
