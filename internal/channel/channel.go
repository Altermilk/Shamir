package channel

import (
	"fmt"
	"math/rand"

	rcv "shamir/internal/reciever"
	snd "shamir/internal/sender"
	crypto "github.com/Altermilk/cryptoMath"
	"github.com/go-datastructures/go-datastructures/queue"
)

type Channel struct {
	Rcv      rcv.Reciever
	Snd      snd.Sender
	que      queue.Queue
	msgParts []rune
	P int32
}

func (C *Channel) SetP(rnd *rand.Rand) {
	C.P = rune(crypto.GetRandomSimpleNum(*rnd))
}

func (C *Channel) SetPrivateKeys(rnd *rand.Rand) {
	C.Rcv.SetPrivateKeys(rnd, C.P)
	C.Snd.SetPrivateKeys(rnd, C.P)
}

func (C *Channel) CountParts(msg []rune) {
	msgl := len(msg)
	t_ := 0
	t := msgl
	F := false
	if msgl>2 && msgl%2!=0{
		t_ ++
		msgl --
		F = true
	}
	
	size := 1
	for {
		if t%2 == 0{
			size *=2
			t /= 2
		}else{
			t = msgl/size
			t_ += msgl - t*size
			break
		}
	}
	
	
	n:=0
	for i := 0; i < t; i++ {
		sizeP := 0
		if msgl - size*i < size{
			break
		}else{
			sizeP = size
		}
		m := make([]rune, sizeP)
		for j := 0; j < size; j++ {
			m[j] = msg[n]
			n++
		}
		C.que.Put(m)
	}
	if F{
		msgl++
		sizeP := msgl - n*size
		m := make([]rune, sizeP)
		for j := 0; j < size; j++ {
			m[j] = msg[n]
			n++
		}
		C.que.Put(m)
	}
}

func OpenChannel(rName, sName string, rnd *rand.Rand) *Channel {
	var C Channel
	C.SetP(rnd)
	C.Rcv = rcv.Reciever{}
	C.Snd = snd.Sender{}
	C.Rcv.Name, C.Snd.Name = rName, sName
	C.que = *queue.New(100)
	fmt.Println("P = ", C.P)
	return &C
}

func (C *Channel) SendPart(msg []rune, rnd *rand.Rand) []rune {
	C.Snd.SetPrivateKeys(rnd, C.P)
	C.Rcv.SetPrivateKeys(rnd, C.P)
	for i := range msg {
		C.Snd.SetM(msg[i])
		C.Snd.CountX1(C.P, &C.Rcv.Buf)
		C.Rcv.CountX2(C.P, &C.Snd.Buf)
		C.Snd.CountX3(C.P, &C.Rcv.Buf)
		C.Rcv.CountX4(C.P)
	}

	return C.Rcv.Buf.Runes
}

func (C *Channel) Send(rnd *rand.Rand) {
	C.CountParts(C.Snd.Buf.Runes)
	C.Snd.Buf.ClearRunes()
	for C.que.Len() > 0 {
		msgPart, _ := C.que.Peek()
		msgPartRunes, ok := msgPart.([]rune)
		if !ok {
		  panic("Unexpected type in queue")
		}
		fmt.Println("\n> msg part: ", msgPartRunes)
		C.msgParts = append(C.msgParts, C.SendPart(msgPartRunes, rnd)...)
		C.que.Get(1) // Удалить элемент из очереди здесь
	
		if(C.que.Len() == 0){ // Проверить длину очереди здесь
		  break
		}
		
		fmt.Println(">> Decoded part: ", C.msgParts)
		C.Rcv.Buf.ClearRunes()
	  }
	
	fmt.Println("\n---------> [ " + string(C.msgParts) + " ]")
}
