package core

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Header struct {
	Version   uint32
	DataHash  Hash
	PrevHash  Hash
	Timestamp int64
	Height    uint32
	Nonce     uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return (fmt.Errorf("Failed to encode version: %w", err))
	}
	if err := binary.Write(w, binary.LittleEndian, &h.PrevHash); err != nil {
		return (fmt.Errorf("Failed to encode previous hash: %w", err))
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return (fmt.Errorf("Failed to encode timestamp: %w", err))
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return (fmt.Errorf("Failed to encode height: %w", err))
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Nonce); err != nil {
		return (fmt.Errorf("Failed to encode nonce: %w", err))
	}
	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return (fmt.Errorf("Failed to decode version: %w", err))
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevHash); err != nil {
		return (fmt.Errorf("Failed to decode previous hash: %w", err))
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return (fmt.Errorf("Failed to decode timestamp: %w", err))
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return (fmt.Errorf("Failed to decode height: %w", err))
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Nonce); err != nil {
		return (fmt.Errorf("Failed to decode nonce: %w", err))
	}
	return nil
}
