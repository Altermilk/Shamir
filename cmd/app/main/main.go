package main

import (
	"math/rand"
	"time"
	ch "shamir/internal/channel"

)

func main(){
	var rnd = rand.New(rand.NewSource(time.Now().Unix()))
	channel := ch.OpenChannel("Andrew", "Bob", rnd)
	channel.Snd.MakeMsg("Hello, dude! IT WORKS :3")
	channel.Send(rnd)
}
