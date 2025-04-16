package network

type NetAddr string

type RPC struct {
	From    NetAddr
	To      NetAddr
	Payload []byte
}
