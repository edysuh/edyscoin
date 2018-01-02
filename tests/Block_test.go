package Block

import (
	"testing"
	"Block"
)

func TestBlock(t *testing.T) {
	b := Block.CreateFirstBlock()
	b.Save()
}
