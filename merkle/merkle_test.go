package merkle

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestMerkleTree_SHA256(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}
	tree := NewMerkleTree(data, SHA256Hash{})
	if tree.Root == nil {
		t.Fatal("Root should not be nil")
	}
	if len(tree.Leaves) != 4 {
		t.Errorf("Expected 4 leaves, got %d", len(tree.Leaves))
	}
}

func TestMerkleTree_SHA3_256(t *testing.T) {
	data := [][]byte{
		[]byte("x"),
		[]byte("y"),
	}
	tree := NewMerkleTree(data, SHA3_256Hash{})
	if tree.Root == nil {
		t.Fatal("Root should not be nil")
	}
	if len(tree.Leaves) != 2 {
		t.Errorf("Expected 2 leaves, got %d", len(tree.Leaves))
	}
}

func TestMerkleTree_OddLeaves(t *testing.T) {
	data := [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
	}
	tree := NewMerkleTree(data, SHA256Hash{})
	if len(tree.Leaves) != 4 {
		t.Errorf("Expected 4 leaves after padding, got %d", len(tree.Leaves))
	}
}

func TestGenerateProofAndVerify(t *testing.T) {
	data := [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
	}
	tree := NewMerkleTree(data, SHA256Hash{})

	for i := range data {
		proof, err := tree.GenerateProof(i)
		if err != nil {
			t.Fatalf("GenerateProof failed: %v", err)
		}
		if proof.Index != i {
			t.Errorf("Proof index mismatch: got %d, want %d", proof.Index, i)
		}
		ok := VerifyProof(data[i], proof, tree.Root.Hash, SHA256Hash{})
		if !ok {
			t.Errorf("Proof verification failed for index %d", i)
		}
	}
}

func TestGenerateProof_InvalidIndex(t *testing.T) {
	data := [][]byte{[]byte("a"), []byte("b")}
	tree := NewMerkleTree(data, SHA256Hash{})
	_, err := tree.GenerateProof(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}
	_, err = tree.GenerateProof(2)
	if err == nil {
		t.Error("Expected error for out-of-range index")
	}
}

func TestVerifyProof_InvalidProof(t *testing.T) {
	data := [][]byte{[]byte("a"), []byte("b")}
	tree := NewMerkleTree(data, SHA256Hash{})
	proof, _ := tree.GenerateProof(0)
	// Tamper with proof
	proof.Path[0].Hash = []byte("invalid")
	ok := VerifyProof(data[0], proof, tree.Root.Hash, SHA256Hash{})
	if ok {
		t.Error("Expected verification to fail for tampered proof")
	}
}

func TestHashFunctionNames(t *testing.T) {
	tests := []struct {
		hf     HashFunction
		expect string
	}{
		{SHA256Hash{}, "SHA256"},
		{SHA3_256Hash{}, "SHA3-256"},
	}

	for _, tt := range tests {
		if got := tt.hf.Name(); got != tt.expect {
			t.Errorf("HashFunction.Name() = %s, want %s", got, tt.expect)
		}
	}
}

func TestMerkleNodeHashConsistency(t *testing.T) {
	data := []byte("hello")
	hf := SHA256Hash{}
	node := newMerkleNode(nil, nil, data, hf)
	expected := hf.Hash(data)
	if !bytes.Equal(node.Hash, expected) {
		t.Errorf("Leaf node hash mismatch")
	}
	left := newMerkleNode(nil, nil, []byte("a"), hf)
	right := newMerkleNode(nil, nil, []byte("b"), hf)
	parent := newMerkleNode(left, right, nil, hf)
	combined := append(left.Hash, right.Hash...)
	expectedParent := hf.Hash(combined)
	if !bytes.Equal(parent.Hash, expectedParent) {
		t.Errorf("Parent node hash mismatch")
	}
}

func TestRootHashHexEncoding(t *testing.T) {
	data := [][]byte{[]byte("x"), []byte("y")}
	tree := NewMerkleTree(data, SHA256Hash{})
	rootHex := hex.EncodeToString(tree.Root.Hash)
	if len(rootHex) != 64 {
		t.Errorf("Expected 64 hex chars for SHA256, got %d", len(rootHex))
	}
}
