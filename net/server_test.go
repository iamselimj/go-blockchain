package net

import (
	"testing"
)

func TestServerStartStop(t *testing.T) {
	lt := NewLocalTransport("LOCAL")
	rt := NewLocalTransport("REMOTE")

	lt.Connect(rt)
	rt.Connect(lt)

	opts := ServerOpts{
		Transports: []Transport{lt, rt},
	}
	s := NewServer(opts)

	go s.Start()
	s.Stop()
}
