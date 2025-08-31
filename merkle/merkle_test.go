package merkle

import (
	"encoding/hex"
	"testing"
)

func TestNewMerkleNode_Leaf(t *testing.T) {
	hf := SHA256Hash{}
	data := []byte("leaf")
	node := NewMerkleNode(nil, nil, data, hf)
	expected := hf.Hash(data)
	if hex.EncodeToString(node.Hash) != hex.EncodeToString(expected) {
		t.Errorf("expected %x, got %x", expected, node.Hash)
	}
	if node.Left != nil || node.Right != nil {
		t.Error("leaf node should not have children")
	}
}

func TestNewMerkleNode_Internal(t *testing.T) {
	hf := SHA256Hash{}
	left := NewMerkleNode(nil, nil, []byte("left"), hf)
	right := NewMerkleNode(nil, nil, []byte("right"), hf)
	node := NewMerkleNode(left, right, nil, hf)
	expected := hf.Hash(append(left.Hash, right.Hash...))
	if hex.EncodeToString(node.Hash) != hex.EncodeToString(expected) {
		t.Errorf("expected %x, got %x", expected, node.Hash)
	}
	if node.Left != left || node.Right != right {
		t.Error("internal node children mismatch")
	}
}

func TestNewMerkleTree_EvenLeaves(t *testing.T) {
	hf := SHA256Hash{}
	data := [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d")}
	tree := NewMerkleTree(data, hf)
	if tree.Root == nil {
		t.Fatal("root should not be nil")
	}
	if len(tree.Leaves) != 4 {
		t.Errorf("expected 4 leaves, got %d", len(tree.Leaves))
	}
}

func TestNewMerkleTree_OddLeaves(t *testing.T) {
	hf := SHA256Hash{}
	data := [][]byte{[]byte("a"), []byte("b"), []byte("c")}
	tree := NewMerkleTree(data, hf)
	if tree.Root == nil {
		t.Fatal("root should not be nil")
	}
	if len(tree.Leaves) != 4 {
		t.Errorf("expected 4 leaves (odd padded), got %d", len(tree.Leaves))
	}
}

func TestMerkleTree_GenerateProof_Valid(t *testing.T) {
	hf := SHA256Hash{}
	data := [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d")}
	tree := NewMerkleTree(data, hf)
	for i := range data {
		proof, err := tree.GenerateProof(i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if proof == nil {
			t.Fatal("proof should not be nil")
		}
		if len(proof.Hashes) == 0 {
			t.Error("proof should contain sibling hashes")
		}
	}
}

func TestMerkleTree_GenerateProof_InvalidIndex(t *testing.T) {
	hf := SHA256Hash{}
	data := [][]byte{[]byte("a"), []byte("b")}
	tree := NewMerkleTree(data, hf)
	_, err := tree.GenerateProof(-1)
	if err == nil {
		t.Error("expected error for negative index")
	}
	_, err = tree.GenerateProof(2)
	if err == nil {
		t.Error("expected error for out-of-range index")
	}
}

func TestVerifyProof_Valid(t *testing.T) {
	hf := SHA256Hash{}
	data := [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d")}
	tree := NewMerkleTree(data, hf)
	for i, leaf := range data {
		proof, err := tree.GenerateProof(i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		valid := VerifyProof(leaf, proof, tree.Root.Hash, hf)
		if !valid {
			t.Errorf("proof verification failed for leaf %d", i)
		}
	}
}

func TestVerifyProof_Invalid(t *testing.T) {
	hf := SHA256Hash{}
	data := [][]byte{[]byte("a"), []byte("b")}
	tree := NewMerkleTree(data, hf)
	proof, err := tree.GenerateProof(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	invalidLeaf := []byte("not a")
	valid := VerifyProof(invalidLeaf, proof, tree.Root.Hash, hf)
	if valid {
		t.Error("expected proof verification to fail for invalid leaf")
	}
}

func TestSHA256Hash_Name(t *testing.T) {
	h := SHA256Hash{}
	if h.Name() != "SHA256" {
		t.Errorf("expected SHA256, got %s", h.Name())
	}
}

func TestSHA256Hash_Hash(t *testing.T) {
	h := SHA256Hash{}
	data := []byte("test")
	hash := h.Hash(data)
	if len(hash) != 32 {
		t.Errorf("expected 32 bytes, got %d", len(hash))
	}
}
