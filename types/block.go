package types

import (
	"errors"

	"github.com/d5c5ceb0/newchain/crypto"
)

//go:generate msgp -tests=false

type Header struct {
	ChainID    byte        `msg:"id"`
	ParentHash crypto.Hash `msg:"parent"`
	TxRootHash crypto.Hash `msg:"txRoot"`
	//ReceiptHash crypto.Hash `msg:"receiptHash"`
	//AccountHash crypto.Hash `msg:"accountHash"`
	Height      uint64         `msg:"height"`
	Nonce       uint64         `msg:"nonce"`
	CreatedTime uint64         `msg:"time"`
	Coinbase    crypto.Address `msg:"coinbase"`
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

func NewBlock() Block {
	return Block{}
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
		return nil, err
	}

	return pubkey.Verify(hash[:], this.Signature)

}

func (this *Block) AttachSignature(sig []byte) {
	this.Signature = sig
}
