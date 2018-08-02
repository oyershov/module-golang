package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
	"time"
)

type Res struct {
	Num  *big.Int
	Time time.Duration
}

var port = "127.0.0.1:8001"

func JsonEncoder(conn net.Conn, i int64) {
	encoder := json.NewEncoder(conn)
	e := encoder.Encode(i)
	if e != nil {
		fmt.Println("Encode error: ", e)
		return
	}
}

func JsonDecoder(conn net.Conn) {
	var msg Res
	decoder := json.NewDecoder(conn)
	e := decoder.Decode(&msg)
	if e != nil {
		fmt.Println("Decode error: ", e)
		return
	}
	fmt.Printf("%s %d\n", msg.Time, msg.Num)
}

func main() {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return
		}
		JsonEncoder(conn, i)
		JsonDecoder(conn)
	}
}
