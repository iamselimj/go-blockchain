package core

import (
	"crypto/sha256"
)

type Hasher[T any] interface {
	Hash(T) Hash
}

type BlockHasher struct {
}

func (BlockHasher) Hash(b *Block) Hash {
	h := sha256.Sum256(b.HeaderData())
	return Hash(h)
}
