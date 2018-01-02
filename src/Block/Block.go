package Block

import (
	// "bytes"
	"crypto/sha256"
	"encoding/gob"
	"io"
	"log"
	"fmt"
	"hash"
	"os"
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

// inefficient file name concat?
func (b *Block) Save() {
	chaindataDir := "chaindata"
	if _, err := os.Stat("./src/" + chaindataDir); os.IsNotExist(err) {
		os.Mkdir("./src/" + chaindataDir, 0777)
	}
	indexStr := fmt.Sprintf("%06d", b.Index)
	// filename := chaindataDir + "/" + indexStr
	filename := indexStr

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("failed to create file -> ", err)
	}
	defer file.Close()

	w := io.Writer(file)
	// gob.Register(sha256.digest{})
	enc := gob.NewEncoder(w)
	err = enc.Encode(b)
	if err != nil {
		log.Fatal("failed to encode block to file -> ", err)
	}
	file.Sync()
}

func CreateBlock(index int,
				 ts time.Time,
				 prev_hash hash.Hash,
				 data string,
				 nonce string) Block {
	return Block{index, ts, sha256.New(), prev_hash, data, nonce}
}

func CreateFirstBlock() Block {
	return CreateBlock(0, time.Now(), nil, "", "")
}

// #check if chaindata folder exists.
// chaindata_dir = 'chaindata'
// if not os.path.exists(chaindata_dir):
//   #make chaindata dir
//   os.mkdir(chaindata_dir)
//   #check if dir is empty from just creation, or empty before
// if os.listdir(chaindata_dir) == []:
//   #create first block
//   first_block = create_first_block()
//   first_block.self_save()

func saveBlock() {
	chaindataDir:= "chaindata"
	// needs to take in an error
	// if !os.IsExist() {
		os.Mkdir(chaindataDir, 744)
	// }
}

// def sync():
//   node_blocks = []
//   #We're assuming that the folder and at least initial block exists
//   chaindata_dir = 'chaindata'
//   if os.path.exists(chaindata_dir):
//     for filename in os.listdir(chaindata_dir):
//       if filename.endswith('.json'): #.DS_Store sometimes screws things up
//         filepath = '%s/%s' % (chaindata_dir, filename)
//         with open(filepath, 'r') as block_file:
//           block_info = json.load(block_file)
//           block_object = Block(block_info) #since we can init a Block object with just a dict
//           node_blocks.append(block_object)
// return node_blocks

func sync() {
	// var nodeBlocks []Block
	// chaindataDir := "chaindata"
	// if !os.IsExist() { }
	// files, err := os.Readdir(-1)
}
