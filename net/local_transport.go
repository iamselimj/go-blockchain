package net

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.RWMutex
	peers     map[NetAddr]*LocalTransport
	closed    bool
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC),
		peers:     make(map[NetAddr]*LocalTransport),
		closed:    false,
	}
}

func (lt *LocalTransport) Connect(peer Transport) error {
	lt.lock.Lock()
	defer lt.lock.Unlock()
	if lt.closed {
		return (fmt.Errorf("connection attempt on closed transport %s", lt.addr))
	}
	lt.peers[peer.Addr()] = peer.(*LocalTransport)
	return (nil)
}

func (lt *LocalTransport) Disconnect(peer NetAddr) {
	lt.lock.Lock()
	defer lt.lock.Unlock()
	if _, ok := lt.peers[peer]; ok {
		delete(lt.peers, peer)
	}
}

func (lt *LocalTransport) Consume() <-chan RPC {
	lt.lock.RLock()
	defer lt.lock.RUnlock()
	if lt.closed {
		return (nil)
	}
	return lt.consumeCh
}

func (lt *LocalTransport) Broadcast(payload []byte) error {
	lt.lock.RLock()
	defer lt.lock.RUnlock()
	if lt.closed {
		return (fmt.Errorf("broadcast failed: transport %s is closed", lt.addr))
	}

	for _, peer := range lt.peers {
		err := peer.SendMessage(peer.Addr(), payload)
		if err != nil {
			return (err)
		}
	}
	return (nil)
}

func (lt *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	lt.lock.RLock()
	defer lt.lock.RUnlock()
	if lt.closed {
		return (fmt.Errorf("send failed: transport %s is closed", lt.addr))
	}
	peer, ok := lt.peers[to]
	if !ok {
		return (fmt.Errorf("[SendMessage] %s: could not send message to %s\n", lt.addr, to))
	}
	peer.consumeCh <- RPC{
		From:    lt.addr,
		To:      to,
		Payload: payload,
	}
	return (nil)
}

func (lt *LocalTransport) Addr() NetAddr {
	return (lt.addr)
}

func (lt *LocalTransport) Close() {
	lt.lock.Lock()
	defer lt.lock.Unlock()
	if lt.closed {
		return
	}
	lt.closed = true
	close(lt.consumeCh)
}
