package block

import (
	"crypto/sha256"
	"time"
)

type Transaction struct {
	Sender   string
	Receiver string
	Amount   int64
}

type Block struct {
	Index        int
	TimeStamp    time.Time
	PrevHash     [32]byte
	Nonce        int64
	Transactions []Transaction
}

func NewBlock(index int, prev [32]byte, nonce int64, tact []Transaction) Block {
	return Block{index, time.Now(), prev, nonce, tact}
}

func NewFirstBlock() Block {
	var noprev [32]byte
	return NewBlock(0, noprev, 0, []Transaction{})
}

// eventually implement merkle tree
func (b *Block) ComputeHash() [32]byte {
	payload := string(b.Index) + b.TimeStamp.String() + string(b.PrevHash[:]) + string(b.Nonce)
	return sha256.Sum256([]byte(payload))
}

// TODO move this to blockchain command
func (b *Block) Save() {
	CreateFile(b)
}
