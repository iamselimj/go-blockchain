package main

import (
	"time"

	net "github.com/iamselimj/go-blockchain/net"
)

func main() {
	lt := net.NewLocalTransport("LOCAL")
	rt := net.NewLocalTransport("REMOTE")

	lt.Connect(rt)
	rt.Connect(lt)

	go func() {
		for {
			rt.SendMessage(lt.Addr(), []byte("Hello from remote!"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := net.ServerOpts{
		Transports: []net.Transport{lt, rt},
	}
	s := net.NewServer(opts)
	s.Start()
}
