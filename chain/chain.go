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

func NewBlockChain(db storage.Database) *BlockChain {
	st := NewStore(db)
	bc := &BlockChain{
		blockDb: st,
	}

	if cur, err := bc.CurrentBlock(); err != nil {
		bc.genesisBlock = types.NewGenesisBlock()
		bc.currentBlock = bc.genesisBlock
		bc.height = 0
	} else {
		bc.height = cur.GetHeight()
		bc.currentBlock = cur
		bc.genesisBlock, _ = bc.GetBlockByHeight(0)
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
		blocks = append(blocks, b)
	}

	return blocks, nil
}

func (this *BlockChain) AddBlock(block *types.Block) error {
	if block.Validation() != nil {
		return errors.New("block verify error")
	}

	err := this.blockDb.StoreBlock(block)
	if err != nil {
		return err
	}

	this.currentBlock = block
	this.height = block.GetHeight()

	return err
}

func (this *BlockChain) CurrentBlock() (*types.Block, error) {
	if this.currentBlock != nil {
		return this.currentBlock, nil
	}

	b, err := this.blockDb.GetCurrentBlock()
	if err != nil {
		return nil, errors.New("no current block")
	}

	return b, nil
}

func (this *BlockChain) CurrentBlockHash() (crypto.Digest, error) {
	if this.currentBlock != nil {
		hash, _ := this.currentBlock.Hash()
		return hash, nil
	}

	b, err := this.blockDb.GetCurrentBlock()
	if err != nil {
		return crypto.Digest{}, errors.New("no current block")
	}

	hash, _ := b.Hash()

	return hash, nil
}

func (this *BlockChain) CurrentHeight() (uint64, error) {
	if this.currentBlock != nil {
		return this.currentBlock.GetHeight(), nil
	}

	b, err := this.blockDb.GetCurrentBlock()
	if err != nil {
		return 0, errors.New("no current block")
	}

	return b.GetHeight(), nil
}
