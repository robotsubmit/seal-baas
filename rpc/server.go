package rpc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/d5c5ceb0/newchain/types"
)

type Chain interface {
	AddBlock(block *types.Block) error
	CurrentBlock() (*types.Block, error)
}

type Server struct {
	chain Chain
}

func NewServer(chain Chain) *Server {
	return &Server{
		chain: chain,
	}
}

func (this *Server) Start() {
	http.HandleFunc("/", this.Handler)
	http.ListenAndServe(":8000", nil)
}

func (this *Server) Handler(w http.ResponseWriter, r *http.Request) {
	b, _ := this.chain.CurrentBlock()
	msg, _ := b.Marshal()
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println(hex.EncodeToString(msg))
	data, _ := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"result":  hex.EncodeToString(msg),
		"id":      1.0,
	})

	w.Write(data)
}
