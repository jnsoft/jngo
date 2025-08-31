package merkle

import (
	"crypto/sha256"
	"crypto/sha3"
)

type HashFunction interface {
	Hash(data []byte) []byte
	Name() string
}

type SHA256Hash struct{}

type SHA3_256Hash struct{}

func (h SHA256Hash) Hash(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

func (h SHA256Hash) Name() string {
	return "SHA256"
}

func (h SHA3_256Hash) Hash(data []byte) []byte {
	sum := sha3.Sum256(data)
	return sum[:]
}

func (h SHA3_256Hash) Name() string {
	return "SHA3-256"
}
