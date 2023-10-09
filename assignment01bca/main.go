// Tayyab Qaisar
// 20I-0590
// Assignment # 1

package main


import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CurrentHash = block.CreateHash()
	return block
}

func (b *Block) CreateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (bc *Blockchain) DisplayBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s\nNonce: %d\nPrevious Hash: %s\nCurrent Hash: %s\n\n",
			block.Transaction, block.Nonce, block.PreviousHash, block.CurrentHash)
	}
}

func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(bc.Blocks) {
		bc.Blocks[index].Transaction = newTransaction
		bc.Blocks[index].CurrentHash = bc.Blocks[index].CreateHash()
	}
}

func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		if bc.Blocks[i].PreviousHash != bc.Blocks[i-1].CurrentHash {
			return false
		}
	}
	return true
}

func main() {
	genesisBlock := NewBlock("Genesis Block", 0, "")
	blockchain := &Blockchain{Blocks: []*Block{genesisBlock}}

	blockchain.Blocks = append(blockchain.Blocks, NewBlock("Alice to Bob", 123, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash))
	blockchain.Blocks = append(blockchain.Blocks, NewBlock("Bob to Carol", 456, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash))

	fmt.Println("Blockchain:")
	blockchain.DisplayBlocks()

	// Change the transaction of the second block
	blockchain.ChangeBlock(1, "Modified Transaction")

	fmt.Println("Blockchain after modification:")
	blockchain.DisplayBlocks()

	if blockchain.VerifyChain() {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
