package sender

import (
	"fmt"
	usr "shamir/internal/user"

	crypto "github.com/Altermilk/cryptoMath"
)

type Sender struct {
	usr.User
	m int32
}

func (s *Sender) MakeMsg(msg string) {
	s.Buf.Runes = []rune(msg)
	fmt.Println("msg: ", s.Buf.Runes)
}

func (s *Sender) SetM(r rune) {
	s.m = int32(r)
}

func (s *Sender) CountX1(p int32, rBuf *usr.Buffer) {
	rBuf.Put(int32(crypto.Modularizate(int(s.m), int(s.C), int(p))))
}

func (s *Sender) CountX3(p int32, rBuf *usr.Buffer) {
	rBuf.Put(int32(crypto.Modularizate(int(s.Buf.Buf), int(s.D), int(p))))
	s.Buf.ClearBuf()
}