package main

import (
	"math/rand"
	"time"
	ch "shamir/internal/channel"

)

func main(){
	var rnd = rand.New(rand.NewSource(time.Now().Unix()))
	channel := ch.OpenChannel("Andrew", "Bob", rnd)
	channel.Snd.MakeMsg("Hello, dude!")
	channel.Send(rnd)
}
