package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func GeneratePrivateKey() (PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return PrivateKey{}, err
	}
	return PrivateKey{key: key}, nil
}

func (pk *PrivateKey) PublicKey() PublicKey {
	return PublicKey{key: &pk.key.PublicKey}
}

func (pk *PrivateKey) Sign(data []byte) (*Signature, error) {
	hash := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, pk.key, hash[:])
	if err != nil {
		return nil, err
	}
	return &Signature{r: r, s: s}, nil
}

func (pk *PrivateKey) Bytes() []byte {
	return (append(pk.key.D.Bytes(), pk.key.PublicKey.X.Bytes()...))
}
