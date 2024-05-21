package reciever

import (
	usr "shamir/internal/user"
	crypto "github.com/Altermilk/cryptoMath"
)

type Reciever struct {
	usr.User
}

func (r *Reciever) CountX2(p int32, sBuf *usr.Buffer) {
	sBuf.Put(int32(crypto.Modularizate(int(r.Buf.Buf), int(r.C), int(p))))
	r.Buf.ClearBuf()
}

func (r *Reciever) CountX4(p int32) {
	r.Buf.Put(int32(crypto.Modularizate(int(r.Buf.Buf), int(r.D), int(p))))
	r.Buf.PutRunes()
	r.Buf.ClearBuf()
}
