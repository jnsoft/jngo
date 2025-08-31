package merkle

import "crypto/sha256"

type HashFunction interface {
	Hash(data []byte) []byte
	Name() string
}

type SHA256Hash struct{}

func (h SHA256Hash) Hash(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

func (h SHA256Hash) Name() string {
	return "SHA256"
}
