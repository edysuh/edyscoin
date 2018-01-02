package block

import (
	// "bytes"
	// "crypto/sha256"
	// "encoding/gob"
	// "io"
	// "log"
	// "fmt"
	"hash"
	// "os"
	"time"
)

type Block struct {
	Index     int
	Ts        time.Time
	Hash      hash.Hash
	Prev_hash hash.Hash
	Data      string
	Nonce     string
}

// // inefficient file name concat?
// func (b *Block) Save() {
// 	chaindataDir := "chaindata"
// 	if _, err := os.Stat("./src/" + chaindataDir); os.IsNotExist(err) {
// 		os.Mkdir("./src/" + chaindataDir, 0777)
// 	}
// 	indexStr := fmt.Sprintf("%06d", b.Index)
// 	// filename := chaindataDir + "/" + indexStr
// 	filename := indexStr

// 	file, err := os.Create(filename)
// 	if err != nil {
// 		log.Fatal("failed to create file -> ", err)
// 	}
// 	defer file.Close()

// 	w := io.Writer(file)
// 	// gob.Register(sha256.digest{})
// 	enc := gob.NewEncoder(w)
// 	err = enc.Encode(b)
// 	if err != nil {
// 		log.Fatal("failed to encode block to file -> ", err)
// 	}
// 	file.Sync()
// }

// func CreateBlock(index int,
// 				 ts time.Time,
// 				 prev_hash hash.Hash,
// 				 data string,
// 				 nonce string) Block {
// 	return Block{index, ts, sha256.New(), prev_hash, data, nonce}
// }

// func CreateFirstBlock() Block {
// 	return CreateBlock(0, time.Now(), nil, "", "")
// }

// func saveBlock() {
// 	chaindataDir:= "chaindata"
// 	// needs to take in an error
// 	// if !os.IsExist() {
// 		os.Mkdir(chaindataDir, 744)
// 	// }
// }

// func sync() {
// 	// var nodeBlocks []Block
// 	// chaindataDir := "chaindata"
// 	// if !os.IsExist() { }
// 	// files, err := os.Readdir(-1)
// }
