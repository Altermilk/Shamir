package user

import (
	"fmt"
	"math/rand"

	crypto "github.com/Altermilk/cryptoMath"
)

type User struct {
	Name string
	C, D int32
	Buf  Buffer
}

type Buffer struct {
	Buf   rune
	Runes []rune
}

func (b *Buffer) ClearBuf() {
	b.Buf = 0 // Set the value of the Buffer to 0 to clear it.
}

func (b *Buffer) Put(data rune) {
	b.Buf = data // Assign the value of data to the Buffer.
}

func (b *Buffer) PutRunes() {
	b.Runes = append(b.Runes, rune(b.Buf))
}

func (b *Buffer) ClearRunes() {
	b.Runes = []rune{}
}

func (u *User) SetPrivateKeys(rnd *rand.Rand, p int32) {
	u.C = CountC(p, rnd)
	u.D = CountD(u.C, p)
	fmt.Println(u.Name, " keys :", u.C, " ", u.D)
}

func CountC(p int32, rnd *rand.Rand) int32 {
	for {
		C := rnd.Int31n(p)
		
		if D, _, _ := crypto.GcdRunes(C, p-1); D == 1{
			return C
		}
	}
}
func CountD(C, p int32) int32 {
	return crypto.ModInvRunes(C, p-1)
}
