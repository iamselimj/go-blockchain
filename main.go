package main

import (
	"time"

	network "github.com/iamselimj/go-blockchain/network"
)

func main() {
	lt := network.NewLocalTransport("LOCAL")
	rt := network.NewLocalTransport("REMOTE")

	lt.Connect(rt)
	rt.Connect(lt)

	go func() {
		for {
			rt.SendMessage(lt.Addr(), []byte("Hello from remote!"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{lt, rt},
	}
	s := network.NewServer(opts)
	s.Start()
}
