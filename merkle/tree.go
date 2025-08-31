package merkle

import (
	"encoding/hex"
	"fmt"
	"strings"
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

func newMerkleNode(left, right *MerkleNode, data []byte, hf HashFunction) *MerkleNode {
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
		leaves = append(leaves, newMerkleNode(nil, nil, d, hf))
	}

	nodes := leaves

	for len(nodes) > 1 {

		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}

		var level []*MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := nodes[i+1]
			parent := newMerkleNode(left, right, nil, hf)
			level = append(level, parent)
		}
		nodes = level
	}

	return &MerkleTree{Root: nodes[0], Leaves: leaves, HashFunction: hf}
}

func (t *MerkleTree) GetRootHash() []byte {
	return t.Root.Hash
}

func (t *MerkleTree) String() string {
	return t.HashFunction.Name() + ":" + hex.EncodeToString(t.Root.Hash)
}

func (t *MerkleTree) PrintMerkleTree() {
	fmt.Println("Merkle Tree Structure:")
	printNode(t.Root, 0, "root")
}

func printNode(node *MerkleNode, level int, label string) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", level)
	hashHex := hex.EncodeToString(node.Hash)
	hashHex = hashHex[:8] + "..."

	fmt.Printf("%s[%s]: %s\n", indent, label, hashHex)

	if node.Left != nil || node.Right != nil {
		printNode(node.Left, level+1, "L")
		printNode(node.Right, level+1, "R")
	}
}
