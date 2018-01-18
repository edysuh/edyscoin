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

func NewBlock(prevHash [32]byte, nonce int64) Block {
	return Block{time.Now(), prevHash, nonce, []*Transaction{} }
}

// TODO get merkle root of transactions then hash the whole struct
func (b *Block) Hash() [32]byte {
	marsh, err := json.Marshal(b)
	if err != nil {
		log.Fatal("error in marshalling into json")
	}
	return sha256.Sum256(marsh)
}
