package net

type NetAddr string

type RPC struct {
	From    NetAddr
	To      NetAddr
	Payload []byte
}
