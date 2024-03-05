package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce int
	previousHash string
	timestamp int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block{
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp   	%d\n", b.timestamp)
	fmt.Printf("nonce    		%d\n", b.nonce)
	fmt.Printf("previouse_hash  %s\n", b.previousHash)
	fmt.Printf("transactions   	%s\n", b.transactions)
	
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain{
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("=", 25))
}

type Blockchain struct {
	transactionPool []string
	chain []*Block
}

func NewBlockChain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "İnit hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func init(){
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := NewBlockChain()
	blockChain.Print()
	blockChain.CreateBlock(5, "Hash 1")
	blockChain.Print()
	blockChain.CreateBlock(5, "Hash 2")
	blockChain.Print()

}

