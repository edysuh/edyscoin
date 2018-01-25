package libedyscoin

import (
	"log"
	"testing"
	"reflect"
)

// redundant, but ok
func TestNewBlockChain(t *testing.T) {
	bc := NewBlockChain()

	if bc.head == nil {
		log.Fatal("NewBlockChain header not set properly")
	}
	if !reflect.DeepEqual(bc.transactions, []*Transaction{}) {
		log.Fatal("NewBlockChain transactions slice not initialized properly")
	}
}

func TestMine(t *testing.T) {
	// t.Skip()
	bc := NewBlockChain()
	bc.SetDifficulty(2)
	bc.Mine()
}

func TestValidProof(t *testing.T) {
	t.Skip()
}

func TestValidChain(t *testing.T) {
	t.Skip()
}
