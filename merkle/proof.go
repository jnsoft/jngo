package merkle

import (
	"bytes"
	"errors"
)

type ProofItem struct {
	Hash     []byte
	Position string // "left" | "right"
}

type MerkleProof struct {
	Path  []ProofItem
	Index int
}

func (t *MerkleTree) GenerateProof(index int) (*MerkleProof, error) {
	if index < 0 || index >= len(t.Leaves) {
		return nil, errors.New("index out of range")
	}

	path := []ProofItem{}
	pos := index
	nodes := t.Leaves

	for len(nodes) > 1 {
		var next []*MerkleNode

		// If odd number of nodes, pad with last node
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := nodes[i+1]
			parent := newMerkleNode(left, right, nil, t.HashFunction)
			next = append(next, parent)

			// Only add sibling for the current position
			if pos == i {
				// left child, add right sibling
				path = append(path, ProofItem{Hash: right.Hash, Position: "right"})
			} else if pos == i+1 {
				// right child, add left sibling
				path = append(path, ProofItem{Hash: left.Hash, Position: "left"})
			}

		}

		pos = pos / 2
		nodes = next
	}

	return &MerkleProof{Path: path, Index: index}, nil
}

func VerifyProof(leafData []byte, proof *MerkleProof, rootHash []byte, hf HashFunction) bool {
	currentHash := hf.Hash(leafData)

	for _, item := range proof.Path {
		if item.Position == "left" {
			currentHash = hf.Hash(append(item.Hash, currentHash...))
		} else {
			currentHash = hf.Hash(append(currentHash, item.Hash...))
		}
	}

	return bytes.Equal(currentHash, rootHash)
	//return hex.EncodeToString(currentHash) == hex.EncodeToString(rootHash)
}
