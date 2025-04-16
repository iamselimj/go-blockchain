package crypto

import "encoding/hex"

type Address [20]uint8

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		panic("Address must be 20 bytes")
	}
	var a Address
	for i := range b {
		a[i] = b[i]
	}
	return Address(a)
}

func (a Address) String() string {
	return (hex.EncodeToString(a.Bytes()))
}

func (a Address) Bytes() []byte {
	b := make([]byte, 20)
	for i := range 20 {
		b[i] = a[i]
	}
	return (b)
}
