package edyscoin

import (
	"log"
	"testing"
	"reflect"
)

// redundant, but ok
func TestNewBlockChain(t *testing.T) {
	bc := NewBlockChain()

	if bc.header == nil {
		log.Fatal("NewBlockChain header not set properly")
	}
	if !reflect.DeepEqual(bc.transactions, []*Transaction{}) {
		log.Fatal("NewBlockChain transactions slice not initialized properly")
	}
}
