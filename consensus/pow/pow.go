package pow

import (
	"fmt"
	"time"

	"github.com/d5c5ceb0/newchain/types"
)

type Chain interface {
	AddBlock(block *types.Block) error
}
type PowServer struct {
	chain Chain
}

func NewPowServer(chain Chain) *PowServer {
	return &PowServer{
		chain: chain,
	}
}

func (this *PowServer) Start() {
	go this.Mining()
}

func (this *PowServer) Mining() {
	for i := uint64(0); i < 10000; i++ {
		b := types.NewBlock(i)
		this.chain.AddBlock(b)
		fmt.Println("height", i)
		time.Sleep(10 * time.Second)

	}
}
