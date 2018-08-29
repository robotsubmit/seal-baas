package pow

import (
	"sync"
	"time"

	"github.com/d5c5ceb0/newchain/types"
)

type Chain interface {
	AddBlock(block *types.Block) error
	CurrentBlock() (*types.Block, error)
}

type TxPool interface {
	GetTransactions() []*types.Transaction
}

type PowServer struct {
	mu    sync.Mutex
	chain Chain
	pool  TxPool
	start bool
	quit  chan interface{}
}

func NewPowServer(chain Chain, pool TxPool) *PowServer {
	return &PowServer{
		chain: chain,
		pool:  pool,
		start: false,
		quit:  make(chan interface{}, 1),
	}
}

func (this *PowServer) Start() {
	if this.start {
		return
	}

	go this.Mining()
	this.start = true
}

func (this *PowServer) Stop() {
	if !this.start {
		return
	}
	this.quit <- true
	this.start = false
}

func (this *PowServer) Mining() {
out:
	for {
		select {
		case <-this.quit:
			break out
		default:
		}

		parentBlock, _ := this.chain.CurrentBlock()
		b := types.NewBlock(parentBlock, nil)
		this.solveDifficulty(b)
		b.Validation()
		this.calcDifficulty()
		this.chain.AddBlock(b)
		time.Sleep(5 * time.Second) //TODO

	}
}

func validateBlock(bhash *types.Block) bool {
	return true
}

func (this *PowServer) solveDifficulty(b *types.Block) *types.Block {
	for i := uint64(0); ; i++ {
		b.ChangeNonce(i)
		if !validateBlock(b) {
			continue
		} else {
			break
		}
	}

	return b
}

func (this *PowServer) calcDifficulty() {
}
