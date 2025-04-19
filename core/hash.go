package core

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

type Hash [32]uint8

func (h Hash) IsZero() bool {
	for _, b := range h {
		if b != 0 {
			return false
		}
	}
	return true
}

func (h Hash) Bytes() []byte {
	b := make([]byte, 32)
	copy(b, h[:])
	return b
}

func (h Hash) String() string {
	return hex.EncodeToString(h.Bytes())
}

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		panic(fmt.Sprintf("HashFromBytes: expected 32 bytes, got %d", len(b)))
	}
	var value [32]uint8
	copy(value[:], b[:32])
	return Hash(value)
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, token); err != nil {
		panic(fmt.Sprintf("RandomBytes: failed to read random bytes: %v", err))
	}
	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
