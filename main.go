package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce int
	previousHash [32]byte
	timestamp int64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block{
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	for _, t := range b.transactions {
		t.Print()
	}

	return b
}

func (b *Block) Hash() [32]byte{
	m, _ :=json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{
		Timestamp int64 			`json:"timestamp"`
		Nonce int 					`json:"nonce"`
		PreviousHash [32]byte 		`json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp: b.timestamp,
		Nonce: b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("timestamp   	%d\n", b.timestamp)
	fmt.Printf("nonce    		%d\n", b.nonce)
	fmt.Printf("previouse_hash  %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
	
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain{
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("=", 25))
}

type Blockchain struct {
	transactionPool []*Transaction
	chain []*Block
}

func NewBlockChain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func init(){
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := NewBlockChain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

type Transaction struct {
	senderBlockChainAddress 	string
	recipientBlockchainAddress 	string 
	value 						float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction)  Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender blockchain address    %s\n", t.senderBlockChainAddress)
	fmt.Printf(" recipient blockchain address   %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                      %.1f", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender string `json:"sender_blockchain_address"`
		Recipient string `json:"recipient_blockchain_address"`
		Value float32 `json:"value"`
	}{
		Sender: t.senderBlockChainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value: t.value,
	})
}

