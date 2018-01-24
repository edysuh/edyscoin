package libedyscoin

import (
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
