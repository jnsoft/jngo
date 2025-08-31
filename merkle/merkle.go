package merkle

import (
	"encoding/hex"
	"errors"
	"sync"
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  []byte
}

type MerkleTree struct {
	Root         *MerkleNode
	HashFunction HashFunction
	Leaves       []*MerkleNode
}

type MerkleProof struct {
	Hashes [][]byte // Sibling hashes
	Index  int      // Index of leaf in original list
}

func NewMerkleNode(left, right *MerkleNode, data []byte, hf HashFunction) *MerkleNode {
	var node MerkleNode

	if left == nil && right == nil {
		node.Hash = hf.Hash(data)
	} else {
		combined := append(left.Hash, right.Hash...)
		node.Hash = hf.Hash(combined)
	}

	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleTree(data [][]byte, hf HashFunction) *MerkleTree {
	var leaves []*MerkleNode

	// must be even number of leaves
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, d := range data {
		leaves = append(leaves, NewMerkleNode(nil, nil, d, hf))
	}

	nodes := leaves

	for len(nodes) > 1 {
		var level []*MerkleNode
		var wg sync.WaitGroup
		var mutex sync.Mutex

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := nodes[i+1]

			wg.Add(1)
			go func(l, r *MerkleNode) {
				defer wg.Done()
				parent := NewMerkleNode(l, r, nil, hf)
				mutex.Lock()
				level = append(level, parent)
				mutex.Unlock()
			}(left, right)
		}

		wg.Wait()

		if len(level)%2 != 0 && len(level) != 1 {
			level = append(level, level[len(level)-1])
		}

		nodes = level
	}

	return &MerkleTree{Root: nodes[0], Leaves: leaves, HashFunction: hf}
}

func (t *MerkleTree) GenerateProof(index int) (*MerkleProof, error) {
	if index < 0 || index >= len(t.Leaves) {
		return nil, errors.New("index out of range")
	}

	path := [][]byte{}
	current := t.Leaves
	for len(current) > 1 {
		next := []*MerkleNode{}
		newIndex := index / 2

		for i := 0; i < len(current); i += 2 {
			left := current[i]
			right := current[i+1]

			if i == index || i+1 == index {
				sibling := right
				if index%2 != 0 {
					sibling = left
				}
				path = append(path, sibling.Hash)
			}

			parent := NewMerkleNode(left, right, nil, t.HashFunction)
			next = append(next, parent)
		}

		current = next
		index = newIndex
	}

	return &MerkleProof{Hashes: path, Index: index}, nil
}

func VerifyProof(leafData []byte, proof *MerkleProof, rootHash []byte, hf HashFunction) bool {
	currentHash := hf.Hash(leafData)
	index := proof.Index

	for _, siblingHash := range proof.Hashes {
		if index%2 == 0 {
			currentHash = hf.Hash(append(currentHash, siblingHash...))
		} else {
			currentHash = hf.Hash(append(siblingHash, currentHash...))
		}
		index /= 2
	}

	return hex.EncodeToString(currentHash) == hex.EncodeToString(rootHash)
}
