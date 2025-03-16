package trie

import (
	"testing"
)

func TestTrieBasicOperations(t *testing.T) {
	trie := NewTrie[int]()

	// Test IsEmpty
	if !trie.IsEmpty() {
		t.Errorf("Expected trie to be empty")
	}

	// Insert elements
	val1, val2 := 10, 20
	trie.Put("apple", &val1)
	trie.Put("banana", &val2)

	// Test Size
	if trie.Size() != 2 {
		t.Errorf("Expected size 2, got %d", trie.Size())
	}

	// Test Contains
	if !trie.Contains("apple") {
		t.Errorf("Expected 'apple' to be in trie")
	}
	if trie.Contains("cherry") {
		t.Errorf("Did not expect 'cherry' to be in trie")
	}

	// Test Get
	if got := trie.Get("apple"); got == nil || *got != 10 {
		t.Errorf("Expected 10, got %v", got)
	}
	if got := trie.Get("banana"); got == nil || *got != 20 {
		t.Errorf("Expected 20, got %v", got)
	}
	if got := trie.Get("cherry"); got != nil {
		t.Errorf("Expected nil, got %v", got)
	}
}

func TestTrieDeletion(t *testing.T) {
	trie := NewTrie[int]()
	val := 42
	trie.Put("hello", &val)

	// Ensure key exists
	if !trie.Contains("hello") {
		t.Errorf("Expected 'hello' to be in trie")
	}

	// Delete the key
	trie.Delete("hello")

	// Ensure key is removed
	if trie.Contains("hello") {
		t.Errorf("Expected 'hello' to be deleted")
	}
	if trie.Size() != 0 {
		t.Errorf("Expected size 0, got %d", trie.Size())
	}
}

func TestTrieKeysWithPrefix(t *testing.T) {
	trie := NewTrie[int]()
	val1, val2, val3 := 1, 2, 3
	trie.Put("car", &val1)
	trie.Put("cat", &val2)
	trie.Put("dog", &val3)

	// Check prefix "ca"
	keys := trie.KeysWithPrefix("ca")
	expected := map[string]bool{"car": true, "cat": true}

	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}

	for _, key := range keys {
		if !expected[key] {
			t.Errorf("Unexpected key: %s", key)
		}
	}

	// Check prefix "do"
	keys = trie.KeysWithPrefix("do")
	if len(keys) != 1 || keys[0] != "dog" {
		t.Errorf("Expected ['dog'], got %v", keys)
	}
}

func TestTrieLongestPrefix(t *testing.T) {
	trie := NewTrie[int]()
	val1, val2 := 5, 10
	trie.Put("flower", &val1)
	trie.Put("flow", &val2)

	longestPrefix := trie.LongestPrefix()
	if longestPrefix != "flow" {
		t.Errorf("Expected 'flow', got '%s'", longestPrefix)
	}

	// Test when no common prefix exists
	trie.Put("dog", &val1)
	longestPrefix = trie.LongestPrefix()
	if longestPrefix != "" {
		t.Errorf("Expected '', got '%s'", longestPrefix)
	}
}

func TestTrieEdgeCases(t *testing.T) {
	trie := NewTrie[int]()

	// Insert and overwrite
	val1, val2 := 100, 200
	trie.Put("key", &val1)
	trie.Put("key", &val2)

	if got := trie.Get("key"); got == nil || *got != 200 {
		t.Errorf("Expected 200 after overwrite, got %v", got)
	}

	// Insert nil value should remove key
	trie.Put("key", nil)
	if trie.Contains("key") {
		t.Errorf("Expected 'key' to be removed")
	}

	// Check size after removal
	if trie.Size() != 0 {
		t.Errorf("Expected size 0 after removal, got %d", trie.Size())
	}
}
