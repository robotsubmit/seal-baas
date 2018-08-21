package types

import (
	"errors"

	"github.com/d5c5ceb0/newchain/crypto"
)

//go:generate msgp -tests=false

type Transaction struct {
	ChainID   uint64         `msg:"chainID"`
	Nonce     uint64         `msg:"nonce"`
	Sender    crypto.Address `msg:"sender"`
	Recipient crypto.Address `msg:"recipient"`
	Value     int64          `msg:"value"`
	Fee       int64          `msg:"fee"`
	Data      []byte         `msg:"data"`
	Signature []byte         `msg:"-"`
}

func NewTransaction() {}

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

func (this *Transaction) Sign() {
}

func (this *Transaction) Verify() {
}

func (this *Transaction) GetNonce() uint64      { return this.Nonce }
func (this *Transaction) GetSender() Address    { return this.Sender }
func (this *Transaction) GetRecipient() Address { return this.Recipient }
func (this *Transaction) GetValue() int64       { return this.Value }
func (this *Transaction) GetFee() int64         { return this.Fee }
func (this *Transaction) GetData() int64        { return this.Data }
func (this *Transaction) GetSignature() []byte  { return this.Signature }
