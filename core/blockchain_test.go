package core

import (
	"testing"

	crypto "github.com/iamselimj/go-blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesisBlock(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)

	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesisBlock(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHashBlockchain(t *testing.T) {
	privKey, _ := crypto.GeneratePrivateKey()

	block := &Block{
		Header: &Header{
			Version:  1,
			PrevHash: RandomHash(),
			Height:   1,
		},
		Transactions: []Transaction{
			{Data: []byte("Test Data")},
		},
	}

	err := block.Sign(privKey)
	assert.Nil(t, err, "La signature du bloc ne devrait pas échouer")

	blockchain := &Blockchain{
		headers: []*Header{},
		storage: NewMemoryStorage(),
	}
	blockchain.headers = append(blockchain.headers, block.Header)

	validator := NewBlockValidator(blockchain)

	blockchain.SetValidator(validator)

	err = validator.ValidateBlock(block)
	assert.Nil(t, err, "Le bloc ne devrait pas échouer à la validation")

	hasher := BlockHasher{}
	calculatedHash := block.Hash(hasher)
	assert.False(t, calculatedHash.IsZero(), "Le hash du bloc ne devrait pas être vide")
}
