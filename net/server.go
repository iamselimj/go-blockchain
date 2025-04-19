package net

import (
	"log"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts

	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}),
	}
}

func (s *Server) Start() {
	s.initTransports()

	for {
		select {
		case rpc := <-s.rpcCh:
			log.Printf("Received RPC: %+v\n", rpc)
		case <-s.quitCh:
			log.Println("Server shutting down...")
			return
		}
	}
}

func (s *Server) Stop() {
	close(s.quitCh)
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for {
				select {
				case rpc := <-tr.Consume():
					s.rpcCh <- rpc
				case <-s.quitCh:
					return
				}
			}
		}(tr)
	}
}
