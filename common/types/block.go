package types

import (
	"github.com/d5c5ceb0/newchain/crypto"
)

//go:generate msgp -tests=false

type Header struct {
	ChainID    byte
	ParentHash crypto.Hash
	TxRootHash crypto.Hash
	//ReceiptHash crypto.Hash
	//AccountHash crypto.Hash
	Height      uint64
	Nonce       uint64
	CreatedTime uint64
	Coinbase    crypto.Address
	ExtraData   []byte
}

type Block struct {
	Header
	txs []Transaction
	//Receipt
	//Account
}
