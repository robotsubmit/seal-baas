package chain

import (
	"sync"

	"github.com/d5c5ceb0/newchain/crypto"
	"github.com/d5c5ceb0/newchain/storage"
	"github.com/d5c5ceb0/newchain/types"
)

type BlockChain struct {
	mu           sync.RWMutex
	blockDb      storage.Database
	currentBlock *types.Block
	genesisBlock *types.Block
}

func NewBlockChain(db storage.Database) BlockChain {
	bc := &BlockChain{
		blockDb: db,
	}

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
