package block

import (
	"fmt"
	"encoding/gob"
	"io"
	"log"
	"os"
)

func CreateDir() string {
	chaindataDir := "chaindata"
	if _, err := os.Stat(chaindataDir); os.IsNotExist(err) {
		os.Mkdir(chaindataDir, 0777)
	}
	return chaindataDir
}

// TODO cleanup
func CreateFile(b *Block) {
	dir := CreateDir()
	indexStr := fmt.Sprintf("%06d", b.Index)
	filename := dir + "/" + indexStr

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("failed to create file -> ", err)
	}
	defer file.Close()

	w := io.Writer(file)
	enc := gob.NewEncoder(w)
	err = enc.Encode(b)
	if err != nil {
		log.Fatal("failed to encode block to file -> ", err)
	}
	file.Sync()
}

// TODO use a read method from DataHandler
func Sync() {
	// var nodeBlocks []Block
	// chaindataDir := "chaindata"
	// if !os.IsExist() { }
	// files, err := os.Readdir(-1)
}
