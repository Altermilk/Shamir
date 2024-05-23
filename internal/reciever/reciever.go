package reciever

import (
	usr "shamir/internal/user"
	crypto "github.com/Altermilk/cryptoMath"
)

type Reciever struct {
	usr.User
}

func (r *Reciever) CountX2(p int32, sBuf *usr.Buffer) {
	sBuf.Put(crypto.ModularizateRune(r.Buf.Buf, r.C, p))
	r.Buf.ClearBuf()
}

func (r *Reciever) CountX4(p int32) {
	r.Buf.Put(crypto.ModularizateRune(r.Buf.Buf, r.D, p))
	r.Buf.PutRunes()
	r.Buf.ClearBuf()
}
