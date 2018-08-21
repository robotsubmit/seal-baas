package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

const (
	SelectedSigAlgo = "p256r1"
)

//go:generate msgp -tests=false

type PublicKey []byte
type PrivateKey []byte

func GenerateKey() (PrivateKey, PublicKey, error) {
	var prikey PrivateKey
	var pubkey PublicKey

	switch SelectedSigAlgo {
	case "p256r1":
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return nil, nil, errors.New("Generate key pair error")
		}

		prikey = privateKey.D.Bytes()
		pubkey = elliptic.Marshal(elliptic.P256(), privateKey.PublicKey.X, privateKey.PublicKey.Y)
	}

	return prikey, pubkey, nil
}

func Sign(hash []byte, prikey *PrivateKey) (sig []byte, err error) {
	if len(hash) != DigestLength {
		return nil, errors.New("length of hash is error")
	}

	switch SelectedSigAlgo {
	case "p256r1":
		privateKey := new(ecdsa.PrivateKey)
		privateKey.Curve = elliptic.P256()
		privateKey.D.SetBytes([]byte(*prikey))

		var r, s *big.Int
		r, s, err = ecdsa.Sign(rand.Reader, privateKey, hash[:])
		if err != nil {
			fmt.Printf("Sign error\n")
			return
		}
		copy(sig[:], r.Bytes())
		copy(sig[len(r.Bytes()):], s.Bytes())
		return
	}

	return nil, errors.New("no signature algorithm.")
}

func (this *PublicKey) Verify(hash []byte, sig []byte) error {

	switch SelectedSigAlgo {
	case "p256r1":
		x, y := elliptic.Unmarshal(elliptic.P256(), []byte(*this))
		pubkey := new(ecdsa.PublicKey)
		pubkey.Curve = elliptic.P256()
		pubkey.X = new(big.Int).Set(x)
		pubkey.Y = new(big.Int).Set(y)

		var r, s big.Int
		r.SetBytes(sig[:len(sig)])
		s.SetBytes(sig[len(sig):])

		if ecdsa.Verify(pubkey, hash[:], &r, &s) {
			return nil
		} else {
			return errors.New("Verify failed.")
		}
	}

	return errors.New("Verify failed.")
}

func (this *PublicKey) ToAddress() (Address, error) {
	return GenAddress([]byte(*this))
}

func (this *PublicKey) Marshal() ([]byte, error) {
	return this.MarshalMsg(nil)
}

func (this *PublicKey) Unmarshal(r []byte) error {
	rem, err := this.UnmarshalMsg(r)
	if len(rem) != 0 {
		return errors.New("unmarshal error.")
	}

	return err
}
