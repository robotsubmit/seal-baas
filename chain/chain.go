package chain

import (
	"errors"
	"sync"

	"github.com/d5c5ceb0/newchain/crypto"
	"github.com/d5c5ceb0/newchain/storage"
	"github.com/d5c5ceb0/newchain/types"
)

type BlockChain struct {
	mu           sync.RWMutex
	blockDb      *Store
	genesisBlock *types.Block
	currentBlock *types.Block
	height       uint64
}

func NewBlockChain(db *storage.Database) *BlockChain {
	st := NewblockDb(db)
	bc := &BlockChain{
		blockDb: st,
	}

	if cur, err := bc.CurrentBlock(); err != nil {
		bc.genesisBlock = NewGenesisBlock()
		bc.currentBlock = bc.genesisBlock
		bc.height = 0
	} else {
		bc.height = cur.GetHeight()
		bc.currentBlock = cur
		bc.genesisBlock = bc.GetBlockByHeight(0)
	}

	return bc
}

func (this *BlockChain) GetBlockByHash(hash *crypto.Digest) (*types.Block, error) {
	return this.blockDb.GetBlockByHash(hash)
}
func (this *BlockChain) GetBlockByHeight(height uint64) (*types.Block, error) {
	return this.blockDb.GetBlockByHeight(height)
}

func (this *BlockChain) GetBlocks(heights []uint64) ([]*types.Block, error) {
	var blocks []*types.Block
	for _, height := range heights {

		b, err := this.blockDb.GetBlockByHeight(height)
		if err != nil {
			return nil, errors.New("get block error")
		}
		append(blocks, b)
	}

	return blocks, nil
}

func (this *BlockChain) AddBlock(block *types.Block) error {
	if block.Validation() != nil {
		return errors.New("block verify error")
	}

	err := this.blockDb.blockDbBlock(block)
	if err != nil {
		return err
	}

	this.currentBlock = block
	this.height = block.GetHeight()

	return err
}

func (this *BlockChain) CurrentBlock() crypto.Digest {
	if this.CurrentBlock != nil {
		return this.CurrentBlock, nil
	}

	b, err := this.blockDb.GetCurrentBlock()
	if err != nil {
		return nil, errors.New("no current block")
	}

	return b, nil
}

func (this *BlockChain) CurrentBlockHash() (crypto.Digest, error) {
	if this.CurrentBlock != nil {
		return this.CurrentBlock.Hash(), nil
	}

	b, err := this.blockDb.GetCurrentBlock()
	if err != nil {
		return nil, errors.New("no current block")
	}

	return b.Hash(), nil
}

func (this *BlockChain) CurrentHeight() (crypto.Digest, error) {
	if this.CurrentBlock != nil {
		return this.CurrentBlock.GetHeight(), nil
	}

	b, err := this.blockDb.GetCurrentBlock()
	if err != nil {
		return nil, errors.New("no current block")
	}

	return b.GetHeight(), nil
}
