package main

import (
	b "./blockchain"
	s "./server"
	"time"
)

func main() {

	var myNode *s.Server = s.CreateServer(1, "Some node", "127.0.0.1:8001")
	var myNode2 *s.Server = s.CreateServer(2, "Another node", "127.0.0.1:8002")
	var myNode3 *s.Server = s.CreateServer(3, "Another node 2", "127.0.0.1:8003")
	go myNode.Receive(":8001")
	go myNode2.Receive(":8002")
	go myNode3.Receive(":8003")

	bc := b.NewBlockchain()
	bc.AddBlock("Alice sends 1 BTC to Peter")
	bc.AddBlock("Peter sends 0.0004 BTC to Anna")

	time.Sleep(5 * time.Second)
	req := &s.Info{2, "127.0.0.1:8001", "Send", bc}
	go myNode.Send(req)

	bc.AddBlock("Anna sends 0.028 ETH to Alice")
	time.Sleep(5 * time.Second)
	req = &s.Info{3, "127.0.0.1:8002", "Send", bc}
	go myNode2.Send(req)

	/*
		time.Sleep(1 * time.Second)
		req := &s.Info{4, "ip:port", "somecode S", []s.Block{s.Block{"message A1"}, s.Block{"message A2"}}}
		go myNode.Send(req)

		time.Sleep(1 * time.Second)
		req = &s.Info{5, "ip:port", "somecode L", []s.Block{s.Block{"message B1"}, s.Block{"message B2"}}}
		go myNode2.Send(req)

		time.Sleep(1 * time.Second)
		req = &s.Info{5, "ip:port", "somecode S", []s.Block{s.Block{"message C1"}, s.Block{"message C2"}}}
		go myNode3.Send(req)
	*/
	time.Sleep(2 * time.Second)

}
