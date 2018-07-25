package server

import (
	b "../blockchain"
)

/*
type Block struct {
	Value string
}
*/
/*Info contains data the nodes send to each other
 */
type Info struct {
	Len    int
	Sender string
	Mcode  string
	//Blocks []Block
	Chain *b.Blockchain
}
