package server

import (
	"encoding/json"
	"fmt"
	"net"
)

func client() {

	conn, err := net.Dial("tcp", ":8080")

	fmt.Println("Err: ", err)

	encoder := json.NewEncoder(conn)
	encoder.Encode(1)
	conn.Close() // we're finished
}
