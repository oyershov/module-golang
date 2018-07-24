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
	Msg      string
	ChainLen int
	Chain    *b.Blockchain
}
