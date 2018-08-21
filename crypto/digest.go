package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

//go:generate msgp -tests=false

const (
	SelectedHashAlgo = "sha256" //TODO selected by configuration file
	DigestLength     = 32
)

type Digest [DigestLength]byte

func sha256Sum(value []byte) Digest {
	digest := Digest{}
	res := sha256.Sum256(value)
	copy(digest[:], res[0:DigestLength])
	return digest
}

func Hash(value []byte) Digest {
	switch SelectedHashAlgo {
	case "sha256":
		return sha256Sum(value)
	}

	return Digest{}
}

func (this *Digest) Size() int {
	return DigestLength
}

func (this *Digest) ToBytes() []byte {
	data := make([]byte, DigestLength)
	copy(data, this[0:DigestLength])
	return data
}

func (this *Digest) ToHexString() string {
	return hex.EncodeToString(this.ToBytes())
}

func (this *Digest) ToString() string {
	return string(this.ToBytes())
}

func (this *Digest) FromString(str string) error {
	data := []byte(str)
	if len(data) != DigestLength {
		return errors.New("decode string error")
	}

	copy(this[0:DigestLength], data)
	return nil
}

func (this *Digest) FromHexString(str string) error {
	data, err := hex.DecodeString(str)
	if err != nil {
		return err
	}

	if len(data) != DigestLength {
		return errors.New("decode string error")
	}

	copy(this[0:DigestLength], data)
	return nil
}

func (this *Digest) FromBytes(b []byte) error {
	if len(b) != DigestLength {
		return errors.New("len(b) != 32")
	}
	copy(this[:], b)

	return nil
}

func (this *Digest) CompareTo(o Digest) int {
	for i := this.Size() - 1; i >= 0; i-- {
		if this[i] > o[i] {
			return 1
		}
		if this[i] < o[i] {
			return -1
		}
	}

	return 0
}

func (this *Digest) Reverse() Digest {
	var s Digest
	copy(s[:], this[:])
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (this *Digest) Marshal() ([]byte, error) {
	return this.MarshalMsg(nil)
}

func (this *Digest) Unmarshal(r []byte) error {
	rem, err := this.UnmarshalMsg(r)
	if len(rem) != 0 {
		return errors.New("unmarshal error.")
	}

	return err
}
