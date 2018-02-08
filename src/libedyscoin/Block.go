package libedyscoin

import (
	"crypto/sha256"
	"encoding/json"
	// "fmt"
	"log"
	"time"
)

type Block struct {
	TimeStamp    time.Time      `json:"ts"`
	PrevHash     [32]byte       `json:"ph"`
	Nonce        int64          `json:"nc"`
	Transactions []Transaction `json:"tx"`
}

func NewBlock(prevHash [32]byte, nonce int64) Block {
	return Block{time.Now(), prevHash, nonce, []Transaction{}}
}

// TODO get merkle root of transactions then hash the whole struct
func (b *Block) Hash() [32]byte {
	// str := fmt.Sprintf("%#v", b)
	marsh, err := json.Marshal(b)
	if err != nil {
		log.Fatal("error in marshalling into json ->", err)
	}
	// fmt.Println("marsh: ", marsh)
	// fmt.Printf("--------\nblock %+v marsh %s\n", b, marsh)
	return sha256.Sum256(marsh)
	// return sha256.Sum256(([]byte)(str))
}
