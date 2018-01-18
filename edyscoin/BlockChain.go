package edyscoin

import (
	"bytes"
)

var DIFFICULTY = 4

type CBlock struct {
	block *Block
	next  *CBlock
}

type BlockChain struct {
	head         *CBlock
	tail         *CBlock
	transactions []*Transaction
}

func NewBlockChain() BlockChain {
	// does the genesis block need proof of work?
	genesis := NewBlock([32]byte{}, 0) 
	cblock := &CBlock{&genesis, (*CBlock)(nil)}
	return BlockChain{cblock, cblock, []*Transaction{}}
}

func (bc *BlockChain) NewTransaction(tact Transaction) {
	bc.transactions = append(bc.transactions, &tact)
}

// validate the current transactions into a new block to the blockchain
// hash the prev last block, generate a nonce, and check for valid proof
func (bc *BlockChain) ProofOfWork() int64 {
	block := NewBlock(bc.tail.block.Hash(), (int64)(0))
	block.transactions = bc.transactions

	for !bc.ValidProof(block) {
		block.nonce++
	}

	bc.transactions = []*Transaction{}
	return block.nonce
}

// is valid proof if first D bytes are zeros, or the guessed hash < 2^(32-D),
// where D is DIFFICULTY
func (bc *BlockChain) ValidProof(block Block) bool {
	guess := block.Hash()
	if bytes.HasPrefix([]byte{}, guess[:DIFFICULTY]) {
		return true
	}
	return false
}

// is valid chain if every block has the correct prev hash, 
// and the correct nonce to solve proof of work
func (bc *BlockChain) ValidChain() bool {
	curr := bc.head
	for curr.next != nil {
		currHash := curr.block.Hash()
		if currHash != curr.next.block.prevHash ||
				!bytes.HasPrefix([]byte{}, currHash[:DIFFICULTY]) {
			return false
		}
		curr = curr.next
	}
	return true
}

// creates a new block with the current transactions and runs proof
// of work to validate the new block into the chain
// func (bc *BlockChain) Mine() {
// 	block := NewBlock(bc.tail.block.Hash(), (int64)(0))
// 	block.transactions = bc.transactions
// 	bc.transactions = []*Transaction{}

// }
