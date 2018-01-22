package libedyscoin

import (
	"bytes"
	"fmt"
)

type CBlock struct {
	block *Block
	next  *CBlock
}

type BlockChain struct {
	head         *CBlock
	tail         *CBlock
	difficulty   int
	transactions []*Transaction
}

func NewBlockChain() BlockChain {
	diff := 4
	// TODO does the genesis block need proof of work?
	genesis := NewBlock([32]byte{}, 0) 
	cblock := &CBlock{&genesis, (*CBlock)(nil)}
	return BlockChain{cblock, cblock, diff, []*Transaction{}}
}

func (bc *BlockChain) NewTransaction(tact Transaction) {
	bc.transactions = append(bc.transactions, &tact)
}

// validate the current transactions into a new block to the blockchain
// hash the prev last block, generate a nonce (just zero for now),
// and check for valid proof; append to chain and reset curr transactions
func (bc *BlockChain) Mine() {
	block := NewBlock(bc.tail.block.Hash(), (int64)(0))
	block.Transactions = bc.transactions

	for !bc.ValidProof(block) {
		block.Nonce++
	}

	bc.transactions = []*Transaction{&Transaction{"s1", "r1", 1000}}
	cblock := &CBlock{&block, (*CBlock)(nil)}
	bc.tail.next = cblock
	bc.tail = cblock
}

func (bc *BlockChain) SetDifficulty(d int) {
	bc.difficulty = d
}

// is valid proof if first D bytes are zeros, or the guessed hash < 2^(32-D),
// where D is difficulty
func (bc *BlockChain) ValidProof(block Block) bool {
	guess := block.Hash()
	fmt.Printf("%v %v\n", block.Nonce, guess)
	if bytes.HasPrefix(make([]byte, 32), guess[:bc.difficulty]) {
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
		if currHash != curr.next.block.PrevHash ||
				!bytes.HasPrefix(make([]byte, 32), currHash[:bc.difficulty]) {
			return false
		}
		curr = curr.next
	}
	return true
}
