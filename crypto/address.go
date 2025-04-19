package crypto

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func AddressFromBytes(b []byte) (Address, error) {
	if len(b) != 20 {
		return Address{}, fmt.Errorf("Address must be 20 bytes, got %d", len(b))
	}
	var a Address
	copy(a[:], b)
	return a, nil
}

func (a Address) String() string {
	return hex.EncodeToString(a.Bytes())
}

func (a Address) Bytes() []byte {
	b := make([]byte, 20)
	copy(b, a[:]) // Utilisation de copy pour une meilleure performance
	return b
}
