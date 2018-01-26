package libedyscoin

import (
	"bytes"
	"fmt"
)

var DIFFICULTY int = 2

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

// TODO does the genesis block need proof of work?
func NewBlockChain() *BlockChain {
	genesis := NewBlock([32]byte{}, 0)
	cblock := &CBlock{&genesis, (*CBlock)(nil)}
	return &BlockChain{cblock, cblock, DIFFICULTY, []*Transaction{}}
}

func (bc *BlockChain) NewTransaction(txn Transaction) {
	bc.transactions = append(bc.transactions, &txn)
}

func (bc *BlockChain) DisplayBlockChain() {
	for curr := bc.head; curr != nil; curr = curr.next {
		fmt.Printf("%+v\n", *curr.block)
	}
}

func (bc *BlockChain) ListTransactions() {
	for _, txn := range bc.transactions {
		fmt.Printf("%+v\n", *txn)
	}
}

func (bc *BlockChain) SetDifficulty(d int) {
	bc.difficulty = d
}

// validate the current transactions into a new block to the blockchain
// hash the prev last block, generate a nonce (just zero for now),
// and check for valid proof; append to chain and reset curr transactions
func (bc *BlockChain) Mine() bool {
	block := NewBlock(bc.tail.block.Hash(), (int64)(0))
	block.Transactions = bc.transactions

	for !bc.ValidProof(block) {
		block.Nonce++
	}

	bc.transactions = []*Transaction{&Transaction{"s1", "r1", 1000}}
	cblock := &CBlock{&block, (*CBlock)(nil)}
	bc.tail.next = cblock
	bc.tail = cblock
	return true
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
	// curr := bc.head
	for curr := bc.head; curr != nil; curr = curr.next {
		currHash := curr.block.Hash()
		if currHash != curr.next.block.PrevHash ||
				!bytes.HasPrefix(make([]byte, 32), currHash[:bc.difficulty]) {
			return false
		}
		// curr = curr.next
	}
	return true
}

// check if other blockchain longer than ours
// consensus dictates the longer chain is the correct chain
func (bcA *BlockChain) Consensus(bcB *BlockChain) error {
	if !bcB.ValidChain() {
		return fmt.Errorf("ERR: new blockchain is not valid!!")
	}

	Acurr, Bcurr := bcA.head, bcB.head
	for Acurr != nil || Bcurr != nil {
		if Acurr == nil {
			bcA.ReplaceChain(bcB)
		} else if Bcurr == nil {
			return fmt.Errorf("ERR: new blockchain is shorter than current blockchain!!")
		}
		Acurr, Bcurr = Acurr.next , Bcurr.next 
	}

	return nil
}

func (bcA *BlockChain) ReplaceChain(bcB *BlockChain) {
	bcA.head, bcA.tail = bcB.head, bcB.tail
}
