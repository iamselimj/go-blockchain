package core

type Blockchain struct {
	storage   Storage
	headers   []*Header
	validator Validator
}

func NewBlockchain(gen *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		storage: NewMeroryStorage(),
	}
	bc.validator = NewBlockValidator(bc)
	err := bc.addBlockWithoutValidation(gen)

	return bc, err
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	return (nil)
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return (height < bc.Height())
}

func (bc *Blockchain) Height() uint32 {
	return (uint32(len(bc.headers) - 1))
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)
	return (nil)
}
