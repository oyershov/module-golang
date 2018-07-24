package main

import (
	b "./blockchain"
	s "./server"
	"time"
)

func main() {

	var myNode *s.Server = s.CreateServer(0, "me", "127.0.0.1:8000")
	go myNode.Receive()

	bc := b.NewBlockchain()
	bc.AddBlock("Alice sends 1 BTC to Peter")
	bc.AddBlock("Peter sends 0.0004 BTC to Anna")

	time.Sleep(5 * time.Second)
	req := &s.Info{"Current status", 2, bc}
	go myNode.Send(req)

	bc.AddBlock("Anna sends 0.028 ETH to Alice")
	time.Sleep(5 * time.Second)
	req = &s.Info{"Current status", 3, bc}
	go myNode.Send(req)

	/*
		time.Sleep(5 * time.Second)
		req = &s.Info{"msg", 5, []s.Block{s.Block{"message L"}, s.Block{"message   KOK"}}}
		go myNode.Send(req)
	*/
}
