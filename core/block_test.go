package core

import (
	"fmt"
	"testing"

	crypto "github.com/iamselimj/go-blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	h := &Header{
		Version:  1,
		PrevHash: RandomHash(),
		Height:   height,
	}
	tx := Transaction{
		Data: []byte("test"),
	}
	return (NewBlock(h, []Transaction{tx}))
}

func TestHashBlock(t *testing.T) {
	b := randomBlock(0)
	fmt.Println(b.Hash((BlockHasher{})))
}

func TestSignBlock(t *testing.T) {
	privKey, _ := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey, _ := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())
	otherPrivKey, _ := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}
