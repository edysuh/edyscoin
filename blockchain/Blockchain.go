package blockchain

import (
	"block"
	"crypto/sha256"
	"strconv"
)

var LEADING_ZEROS = 4

type Blockchain struct {
	chain            []*block.Block
	currTransactions []block.Transaction
}


func NewBlockchain() Blockchain {
	genesis := block.NewFirstBlock()
	chain := []*block.Block{&genesis}
	return Blockchain{chain, make([]block.Transaction, block.SIZE)}
}

// proof 
func (bc *Blockchain) NewBlock() block.Block {
	prev := bc.chain[len(bc.chain)-1]
	newLast := block.NewBlock(prev.Index+1, prev.ComputeHash(), 0, bc.currTransactions)
	bc.currTransactions = make([]block.Transaction, 32)
	return newLast
}

func (bc *Blockchain) NewTransaction(tact block.Transaction) int {
	bc.currTransactions = append(bc.currTransactions, tact)
	return len(bc.chain)
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
	for i := 0; i < LEADING_ZEROS; i++ {
		if guess[i] != 0 {
			return false
		}
	}
	return true
}

func (bc *Blockchain) ValidChain(c Blockchain) bool {
	for i := 1; i < len(c.chain); i++ {
		if c.chain[i-1].PrevHash != c.chain[i].ComputeHash() {
			return false
		}
		if !ValidProof(c.chain[i-1].Nonce, c.chain[i].Nonce) {
			return false
		}
	}
	return true
}
