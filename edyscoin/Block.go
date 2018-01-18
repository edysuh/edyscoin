package edyscoin

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"time"
)

type Block struct {
	timeStamp    time.Time
	prevHash     [32]byte
	nonce        int64
	transactions []*Transaction
}

// dooes block creation only occur within the chain,
// thus this method should stay as a member of BlockChain?
func NewBlock(prevHash [32]byte, nonce int64) Block {
	return Block{time.Now(), prevHash, nonce, []*Transaction{} }
}

// get merkle root of transactions then hash the whole struct
func (b *Block) CalculateHash() [32]byte {
	// merkle := Merkle(b.transactions)
	// marsh, err := json.Marshal(Block{b.timeStamp, b.prevHash, b.nonce, merkle})
	marsh, err := json.Marshal(b)
	if err != nil {
		log.Fatal("error in marshalling into json")
	}
	return sha256.Sum256(marsh)
}
