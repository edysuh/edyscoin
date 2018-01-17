package blockchain

import (
	"testing"
	"block"
)

func TestNewTransaction(t *testing.T) {
	c := NewBlockChain()
	i1 := c.NewTransaction(block.Transaction{"s1", "r1", 1000})
	if i1 != 1 {
		t.Errorf("i1: %d", i1)
	}
	i2 := c.NewTransaction(block.Transaction{"s2", "r2", 2000})
	if i2 != 1 {
		t.Errorf("i2: %d", i2)
	}
	i3 := c.NewTransaction(block.Transaction{"s3", "r3", 3000})
	if i3 != 1 {
		t.Errorf("i3: %d", i3)
	}
}

func TestNewBlock(t *testing.T) {
	c := NewBlockChain()
	c.NewTransaction(block.Transaction{"s1", "r1", 1000})
	c.NewTransaction(block.Transaction{"s2", "r2", 2000})
	c.NewTransaction(block.Transaction{"s3", "r3", 3000})
	bl := c.NewBlock()
	// hash := bl.ComputeHash()
	// tarr := []block.Transaction{block.Transaction{"s1", "r1", 1000},
	// 						   block.Transaction{"s2", "r2", 2000},
	// 						   block.Transaction{"s3", "r3", 3000}}
	// var prev [32]byte
	// struc := block.Block{1, bl.TimeStamp, prev, 0, tarr}
	// if bl != struc {
	// 	t.Errorf("new block was incorrectly created")
	// }
	t.Logf("%v\n", bl)
}

// func TestProofOfWork(t *testing.T) {
// 	var last int64 = 234
// 	sol := ProofOfWork(last)
// 	t.Logf("%v", sol)
// }

func TestValidChain(t *testing.T) {
	c := NewBlockChain()
	c.NewTransaction(block.Transaction{"s1", "r1", 1000})
	c.NewBlock()
	c.NewTransaction(block.Transaction{"s2", "r2", 2000})
	c.NewBlock()
	c.NewTransaction(block.Transaction{"s3", "r3", 3000})
	c.NewBlock()
	tf := ValidChain(c)
	if tf != true {
		// c.Display()
		t.Error()
	}
}
