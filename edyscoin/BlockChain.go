package edyscoin

import (

)

type CBlock struct {
	blockptr *Block
	next     *CBlock
}

type BlockChain struct {
	head         *CBlock
	tail         *CBlock
	transactions []*Transaction
}

func NewBlockChain() BlockChain {
	genesis := NewBlock([32]byte{}, 0) 
	cblock := &CBlock{&genesis, (*CBlock)(nil)}
	return BlockChain{cblock, cblock, []*Transaction{}}
}

func (bc *BlockChain) NewTransaction(tact Transaction) {
	bc.transactions = append(bc.transactions, &tact)
}

// proof of work
func (bc *BlockChain) ProofOfWork(nonce int64) int64 {

	return 0
}

// validate proof
func (bc *BlockChain) ValidateProof(nonce int64) bool {
	return false
}

// validate chain

// mine
