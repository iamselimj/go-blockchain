package crypto

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"math/big"
)

type Signature struct {
	r *big.Int
	s *big.Int
}

func (sig *Signature) Verify(pubKey PublicKey, data []byte) bool {
	hash := sha256.Sum256(data)
	return (ecdsa.Verify(pubKey.key, hash[:], sig.r, sig.s))
}
