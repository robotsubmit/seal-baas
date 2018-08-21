package types

import "errors"

//go:generate msgp -tests=false

type Transaction struct {
	Nonce     uint64         `msg:"nonce"`
	Sender    crypto.Address `msg:"sender"`
	Recipient crypto.Address `msg:"recipient"`
	Value     int64          `msg:"value"`
	Fee       int64          `msg:"fee"`
	Data      []byte         `msg:"data"`
	Signature []byte         `msg:"signature"`
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

func (this *Transaction) Hash()

func (this *Transaction) Verify()
