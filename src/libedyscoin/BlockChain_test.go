package libedyscoin

import (
	"log"
	"testing"
	"reflect"
)

// redundant, but ok
func TestNewBlockChain(t *testing.T) {
	bc := NewBlockChain()

	if bc.Head == nil {
		log.Fatal("NewBlockChain header not set properly")
	}
	if !reflect.DeepEqual(bc.Transactions, []*Transaction{}) {
		log.Fatal("NewBlockChain transactions slice not initialized properly")
	}
}

func TestMine(t *testing.T) {
	// t.Skip()
	bc := NewBlockChain()
	bc.SetDifficulty(1)
	bc.NewTransaction(&Transaction{"a", "b", 1})
	tru := bc.Mine()
	bc.NewTransaction(&Transaction{"b", "c", 2})
	tru = bc.Mine()
	if tru {
		bc.DisplayBlockChain()
	}
	log.Println("----------------")
	tru = bc.ValidChain()
	if !tru {
		t.Error("Not a valid chain")
	}
}

func TestValidProof(t *testing.T) {
	t.Skip()
}

func TestValidChain(t *testing.T) {
	t.Skip()
}
