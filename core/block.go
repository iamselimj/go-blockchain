package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	crypto "github.com/iamselimj/go-blockchain/crypto"
)

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature
	hash         Hash
}

func NewBlock(h *Header, txs []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txs,
	}
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
	if err != nil {
		return err
	}
	b.Validator = privKey.PublicKey()
	b.Signature = sig
	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("Block has no signature\n")
	}
	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("Block has invalid signature\n")
	}
	return nil
}

func (b *Block) Hash(hasher Hasher[*Block]) Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}
	return b.hash
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return (enc.Encode(w, b))
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return (dec.Decode(r, b))
}

func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)
	return (buf.Bytes())
}
