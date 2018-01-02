package Block

import (
	"testing"
	"block"
)

func TestBlock(t *testing.T) {
	b := block.CreateFirstBlock()
	b.Save()
}
