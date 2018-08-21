package crypto

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	x := Hash([]byte{0})
	fmt.Println(x)
	fmt.Println(x.Size())
	fmt.Println(x.ToBytes())
	fmt.Println(x.ToHexString())
	y := x.ToBytes()
	var z Digest
	z.FromBytes(y)
	fmt.Println(x.CompareTo(z))
	fmt.Println(x.Reverse())
	fmt.Println(x.Marshal())
	a, _ := x.Marshal()
	var b Digest
	b.Unmarshal(a)
	fmt.Println("-", b)

}
