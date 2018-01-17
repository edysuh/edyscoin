package blockchain

import (
	"block"
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
)

var LEADING_ZEROS = 4

type BlockChain struct {
	chain            []*block.Block
	currTransactions []block.Transaction
}

func (bc *BlockChain) Display() {
	for _, block := range bc.chain {
		fmt.Printf("%v\n", block)
	}
}

func NewBlockChain() BlockChain {
	genesis := block.NewFirstBlock()
	return BlockChain{[]*block.Block{&genesis}, []block.Transaction{}}
}

// proof/nonce
// TODO need to run proof of work in order to create a new block
func (bc *BlockChain) NewBlock() block.Block {
	prev := bc.LastBlock()
	newLast := block.NewBlock(prev.Index+1, prev.ComputeHash(), 0, bc.currTransactions)
	bc.chain = append(bc.chain, &newLast)
	bc.currTransactions = []block.Transaction{}
	return newLast
}

func (bc *BlockChain) NewTransaction(tact block.Transaction) int {
	bc.currTransactions = append(bc.currTransactions, tact)
	return len(bc.chain)
}

// unnecessary code atm
func (bc *BlockChain) NewTransactionWithValues(s, r string, a int64) int {
	bc.currTransactions = append(bc.currTransactions, block.Transaction{s, r, a})
	return len(bc.chain)
}

func (bc *BlockChain) LastBlock() *block.Block {
	return bc.chain[len(bc.chain)-1]
}

func ProofOfWork(lastproof int64) int64 {
	var proof int64 = 0
	for !ValidProof(lastproof, proof) {
		proof++
	}
	return proof
}

func ValidProof(lastproof, proof int64) bool {
	guess := sha256.Sum256([]byte(strconv.FormatInt(lastproof * proof, 10)))
	fmt.Printf("%v, %v, %v\n", lastproof, proof, guess)
	if bytes.Equal(guess[:4], []byte{0, 0, 0, 0}) {
		return true
	}
	return false
}

func ValidChain(c BlockChain) bool {
	for i := 1; i < len(c.chain); i++ {
		if c.chain[i-1].ComputeHash() != c.chain[i].PrevHash {
			return false
		}
		// fmt.Printf("%v\n", i)
		// fmt.Printf("%v, %v\n", c.chain[i-1].Nonce, c.chain[i].Nonce)
		// if !ValidProof(c.chain[i-1].Nonce, c.chain[i].Nonce) {
		// 	return false
		// }
	}
	return true
}
