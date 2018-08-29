package types

import (
	"errors"
	"time"

	"github.com/d5c5ceb0/newchain/crypto"
)

//go:generate msgp -tests=false

type Header struct {
	ChainID    byte          `msg:"id"`
	ParentHash crypto.Digest `msg:"parent"`
	TxRootHash crypto.Digest `msg:"txRoot"`
	Height     uint64        `msg:"height"`
	Nonce      uint64        `msg:"nonce"`
	Timestamp  int64         `msg:"time"`
}

func (this *Header) Marshal() ([]byte, error) {
	return this.MarshalMsg(nil)
}

func (this *Header) Unmarshal(r []byte) error {
	rem, err := this.UnmarshalMsg(r)
	if len(rem) != 0 {
		return errors.New("unmarshal error.")
	}

	return err
}

func (this *Header) Hash() (crypto.Digest, error) {
	data, err := this.Marshal()
	if err != nil {
		return crypto.Digest{}, err
	}
	return crypto.Hash(data), nil
}

type Block struct {
	Header    `msg:"header"`
	Txs       []Transaction `msg:"txs"`
	Signature []byte        `msg:"sig"`
}

func NewBlock(parentBlock *Block, txs []*Transaction) *Block {
	pHash, _ := parentBlock.Hash()
	return &Block{
		Header: Header{
			ChainID:    parentBlock.ChainID,
			ParentHash: pHash,
			TxRootHash: crypto.Digest{}, //TODO use tx merkle tree
			Height:     parentBlock.Height + 1,
			Nonce:      0,
			Timestamp:  time.Now().Unix(),
		},
	}
}

func (this *Block) ChangeNonce(nonce uint64) *Block {
	this.Nonce = nonce
	return this
}

func (this *Block) Marshal() ([]byte, error) {
	return this.MarshalMsg(nil)
}

func (this *Block) Unmarshal(r []byte) error {
	rem, err := this.UnmarshalMsg(r)
	if len(rem) != 0 {
		return errors.New("unmarshal error.")
	}

	return err
}

func (this *Block) Hash() (crypto.Digest, error) {
	data, err := this.Marshal()
	if err != nil {
		return crypto.Digest{}, err
	}
	return crypto.Hash(data), nil
}

func (this *Block) Sign(privkey *crypto.PrivateKey) ([]byte, error) {
	hash, err := this.Header.Hash()
	if err != nil {
		return nil, err
	}

	return crypto.Sign(hash[:], privkey)
}

func (this *Block) Verify(pubkey *crypto.PublicKey) error {
	hash, err := this.Header.Hash()
	if err != nil {
		return err
	}

	return pubkey.Verify(hash[:], this.Signature)

}

func (this *Block) AttachSignature(sig []byte) {
	this.Signature = sig
}

func (this *Block) Validation() error {
	return nil
}

func (this *Block) GetTransactionByHash(hash *crypto.Digest) (Transaction, error) {
	for _, tx := range this.Txs {
		h, err := tx.Hash()
		if err != nil {
			return Transaction{}, err
		}
		if hash.CompareTo(h) != 0 { //cache hash in Transaction
			return tx, nil
		}
	}

	return Transaction{}, errors.New("no transaction")
}

func NewGenesisBlock() *Block {
	genesisHeader := &Header{
		ChainID:    DefaultChainID,
		ParentHash: crypto.Digest{},
		TxRootHash: crypto.Digest{},
		Height:     0,
		Nonce:      0,
		Timestamp:  time.Date(2018, time.August, 0, 0, 0, 0, 0, time.UTC).Unix(),
	}
	// genesis block
	genesisBlock := &Block{
		Header: *genesisHeader,
	}

	return genesisBlock
}

func (this *Block) GetChainID() byte               { return this.ChainID }
func (this *Block) GetParentHash() crypto.Digest   { return this.ParentHash }
func (this *Block) GetTxRootHash() crypto.Digest   { return this.TxRootHash }
func (this *Block) GetHeight() uint64              { return this.Height }
func (this *Block) GetNonce() uint64               { return this.Nonce }
func (this *Block) GetCreateTime() int64           { return this.Timestamp }
func (this *Block) GetHeader() *Header             { return &this.Header }
func (this *Block) GetTransactions() []Transaction { return this.Txs }
func (this *Block) GetSignature() []byte           { return this.Signature }
