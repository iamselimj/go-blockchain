package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
)

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (pk *PublicKey) Bytes() []byte {
	return (elliptic.MarshalCompressed(pk.key, pk.key.X, pk.key.Y))
}

func (pk *PublicKey) Address() (Address, error) {
	h := sha256.Sum256(pk.Bytes())
	return AddressFromBytes(h[len(h)-20:])
}
