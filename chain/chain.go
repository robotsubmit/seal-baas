package chain

import (
	"sync"

	"github.com/d5c5ceb0/newchain/crypto"
	"github.com/d5c5ceb0/newchain/storage"
	"github.com/d5c5ceb0/newchain/types"
)

type BlockChain struct {
	mu           sync.RWMutex
	blockDb      *Store
	currentBlock *types.Block
	genesisBlock *types.Block
	height       uint64
}

func NewBlockChain(db *storage.Database) *BlockChain {
	st := NewStore(db)
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

func (this *BlockChain) GetBlockByHash(hash crypto.Digest) *types.Block {
}
func (this *BlockChain) GetBlockByHeight(num uint64) *types.Block {
}
func (this *BlockChain) AddBlock(block types.Block) (uint64, error) {
}
func (this *BlockChain) CurrentBlockHash() crypto.Digest {
}
func (this *BlockChain) CurrentBlock() crypto.Digest {
}
func (this *BlockChain) CurrentHeight() crypto.Digest {
}
