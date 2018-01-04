package blockchain

import (
	"block"
	"crypto/sha256"
	"strconv"
)

var LEADING_ZEROS = 4

type Blockchain struct {
	Genesis          *block.Block
	LastBlock        *block.Block
	CurrTransactions []block.Transaction
}

func NewBlockchain() Blockchain {
	genesis := block.NewFirstBlock()
	return Blockchain{&genesis, &genesis, make([]block.Transaction, block.SIZE)}
}

// proof 
func (bc *Blockchain) NewBlock() block.Block {
	prev := bc.LastBlock
	last := block.NewBlock(prev.Index+1, prev.ComputeHash(), "", bc.CurrTransactions)
	bc.CurrTransactions = make([]block.Transaction, 32)
	return last
}

func (bc *Blockchain) NewTransaction(tact block.Transaction) int {
	bc.CurrTransactions = append(bc.CurrTransactions, tact)
	return bc.LastBlock.Index+1
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
