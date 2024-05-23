package channel

import (
	"fmt"
	"math/rand"
	"strconv"
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
	p        int64
	P int32
	pr       []rune
}

func (C *Channel) SetP(rnd *rand.Rand) {
	C.p = int64(crypto.GetRandomSimpleNum(*rnd))
	ps := strconv.FormatInt(int64(C.p), 10)
	C.pr = []rune(ps)
	fmt.Println("p = ", C.p, "pr := ", C.pr)
}

func (C *Channel) SetPrivateKeys(rnd *rand.Rand) {
	C.Rcv.SetPrivateKeys(rnd, int32(C.p))
	C.Snd.SetPrivateKeys(rnd, int32(C.p))
}

func (C *Channel) CountParts(msg []rune) {
	pl, msgl := len(C.pr), len(string(msg))
	var t, t_ int
	if pl<msgl{
		t = msgl/pl
		t_ = msgl - t*pl
		// if t_ > 0{
		// 	t++
		// }
	}else{
		t = 1
		t_ = pl - msgl
	}
	
	n:=0
	for i := 0; i < t; i++ {
		size := 0
		if msgl>=pl{
			if n <= t*pl{
				size = pl
			}else{
				size = t_
			}
		}else{
			size = pl - msgl
		}
		m := make([]rune, size)
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
	return &C
}

func (C *Channel) SendPart(msg []rune, rnd *rand.Rand) []rune {
	C.Snd.SetPrivateKeys(rnd, int32(C.p))
	C.Rcv.SetPrivateKeys(rnd, int32(C.p))
	for i := range msg {
		C.Snd.SetM(msg[i])
		C.P = int32(C.pr[i])
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
		fmt.Println("> msg part: ", msgPartRunes)
		C.msgParts = append(C.msgParts, C.SendPart(msgPartRunes, rnd)...)
		C.que.Get(1) // Удалить элемент из очереди здесь
	
		if(C.que.Len() == 0){ // Проверить длину очереди здесь
		  break
		}
		fmt.Println(">> Decoded part: ", C.msgParts)
		C.Rcv.Buf.ClearRunes()
		fmt.Println(C.que.Len())
	  }
	
	fmt.Println(string(C.msgParts))
}
