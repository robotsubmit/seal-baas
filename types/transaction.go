package types

import (
	"errors"

	"github.com/d5c5ceb0/newchain/crypto"
)

const (
	DefaultChainID byte = 1
)

//go:generate msgp -tests=false

type TxContent struct {
	ChainID   byte           `msg:"chainID"`
	Nonce     uint64         `msg:"nonce"`
	Sender    crypto.Address `msg:"sender"`
	Recipient crypto.Address `msg:"recipient"`
	Value     int64          `msg:"value"`
	Fee       int64          `msg:"fee"`
	Data      []byte         `msg:"data"`
}

func (this *TxContent) Hash() (crypto.Digest, error) {
	data, err := this.Marshal()
	if err != nil {
		return crypto.Digest{}, err
	}
	return crypto.Hash(data), nil
}

func (this *TxContent) Marshal() ([]byte, error) {
	return this.MarshalMsg(nil)
}

func (this *TxContent) Unmarshal(r []byte) error {
	rem, err := this.UnmarshalMsg(r)
	if len(rem) != 0 {
		return errors.New("unmarshal error.")
	}

	return err
}

type Transaction struct {
	TxContent `msg:"txContent"`
	Signature []byte `msg:"signature"`
}

func NewTransaction(nonce uint64, from, to crypto.Address, value, fee int64, data []byte) *Transaction {
	d := TxContent{
		ChainID:   DefaultChainID,
		Nonce:     nonce,
		Sender:    from,
		Recipient: to,
		Value:     value,
		Fee:       fee,
		Data:      data,
	}

	return &Transaction{TxContent: d}
}

func (this *Transaction) Marshal() ([]byte, error) {
	return this.MarshalMsg(nil)
}

func (this *Transaction) Unmarshal(r []byte) error {
	rem, err := this.UnmarshalMsg(r)
	if len(rem) != 0 {
		return errors.New("unmarshal error.")
	}

	return err
}

func (this *Transaction) Hash() (crypto.Digest, error) {
	data, err := this.Marshal()
	if err != nil {
		return crypto.Digest{}, err
	}
	return crypto.Hash(data), nil
}

func (this *Transaction) Sign(privkey *crypto.PrivateKey) ([]byte, error) {
	hash, err := this.TxContent.Hash()
	if err != nil {
		return nil, err
	}

	return crypto.Sign(hash[:], privkey)
}

func (this *Transaction) AttachSignature(sig []byte) {
	this.Signature = sig
}

func (this *Transaction) Verify(pubkey *crypto.PublicKey) error {
	hash, err := this.TxContent.Hash()
	if err != nil {
		return err
	}

	return pubkey.Verify(hash[:], this.Signature)
}

func (this *Transaction) Validation() error {
	return nil
}

func (this *Transaction) GetNonce() uint64             { return this.Nonce }
func (this *Transaction) GetSender() crypto.Address    { return this.Sender }
func (this *Transaction) GetRecipient() crypto.Address { return this.Recipient }
func (this *Transaction) GetValue() int64              { return this.Value }
func (this *Transaction) GetFee() int64                { return this.Fee }
func (this *Transaction) GetData() []byte              { return this.Data }
func (this *Transaction) GetSignature() []byte         { return this.Signature }
