package core

import (
	"fmt"
)

type Validator interface {
	ValidateBlock(b *Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (bv *BlockValidator) ValidateBlock(b *Block) error {
	if err := b.Verify(); err != nil {
		return fmt.Errorf("block verification failed: %w", err)
	}
	if b.Height != bv.bc.Height()+1 {
		return fmt.Errorf("invalid block height")
	}
	return nil
}
