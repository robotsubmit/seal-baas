package crypto

import (
	"crypto/sha256"
	"errors"

	"golang.org/x/crypto/ripemd160"
)

const (
	AddrLength = 20
)

//go:generate msgp -tests=false

type Address [AddrLength]byte

func GenAddress(code []byte) (Address, error) {
	temp := sha256.Sum256(code)
	md := ripemd160.New()
	f := md.Sum(temp[:])

	if len(f) != AddrLength {
		return Address{}, errors.New("length of address error")
	}

	var addr Address
	copy(addr[:], f[:])

	return addr, nil
}

func (this *Address) Marshal() ([]byte, error) {
	return this.MarshalMsg(nil)
}

func (this *Address) Unmarshal(r []byte) error {
	rem, err := this.UnmarshalMsg(r)
	if len(rem) != 0 {
		return errors.New("unmarshal error.")
	}

	return err
}

func (this *Address) Verify() error {
	return nil
}
