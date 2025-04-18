package net

type Transport interface {
	Connect(peer Transport) error
	Disconnect(peer NetAddr)
	Consume() <-chan RPC
	Broadcast(payload []byte) error
	SendMessage(to NetAddr, payload []byte) error
	Addr() NetAddr
}
