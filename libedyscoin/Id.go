package libedyscoin

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

const IdBytes = 256/8

type Id [IdBytes]byte

func NewId(addr string) Id {
	return sha256.Sum256([]byte(addr + time.Now().String()))
}

func (id *Id) ToString() string {
	return fmt.Sprintf("%v", *id)
}

func (id *Id) Equals(o Id) bool {
	a := ([IdBytes]byte)(*id)
	b := ([IdBytes]byte)(o)
	return bytes.Equal(a[:], b[:])
}
