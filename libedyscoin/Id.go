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

func (id *Id) String() string {
	str := "["
	for _, b := range id {
		c := fmt.Sprintf("%v ", b)
		str += c
	}
	return str + "]"
}
