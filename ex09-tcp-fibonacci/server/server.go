package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net"
	"time"
)

type Res struct {
	Num  *big.Int
	Time time.Duration
}

var ServerCash = make(map[int]*big.Int)

var port = "127.0.0.1:8001"

func Calc(n int) *big.Int {
	fn := make(map[int]*big.Int)

	for i := 0; i <= n; i++ {
		var f = big.NewInt(0)
		if i <= 2 {
			f.SetUint64(1)
		} else {
			f = f.Add(fn[i-1], fn[i-2])
		}
		fn[i] = f
	}
	return fn[n]
}

func JsonEncoder(conn net.Conn, answer *Res) {
	encoder := json.NewEncoder(conn)
	e := encoder.Encode(answer)
	if e != nil {
		fmt.Println("Encode error: ", e)
		return
	}
}

func JsonDecoder(conn net.Conn) (*Res, bool) {
	var num int
	decoder := json.NewDecoder(conn)
	e := decoder.Decode(&num)
	switch e {
	case io.EOF:
		return nil, false
	}
	fmt.Println("From client: ", num)

	var answer Res
	sTime := time.Now()
	if ServerCash[num] != nil {
		answer.Num = ServerCash[num]
	} else {
		answer.Num = Calc(num)
		ServerCash[num] = answer.Num
	}
	eTime := time.Since(sTime)
	answer.Time = eTime

	return &answer, true
}

func HandleConn(conn net.Conn) {
	fmt.Println("Accepted connection")
	for {
		answer, success := JsonDecoder(conn)
		if !success {
			conn.Close()
			return
		}
		JsonEncoder(conn, answer)
	}
}

func main() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Can`t create server")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		go HandleConn(conn)
	}
}
