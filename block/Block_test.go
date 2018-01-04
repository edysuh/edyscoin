package block

import (
	"testing"
)

func TestBlock(t *testing.T) {
	b := NewFirstBlock()
	b.Save()
}
