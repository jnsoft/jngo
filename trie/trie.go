package trie

import "strings"

// Trie represents an R-way trie.
type Trie[T any] struct {
	root *node[T]
	size int
}

type node[T any] struct {
	val  *T
	next [256]*node[T] // ASCII-based
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{}
}

func (t *Trie[T]) Size() int {
	return t.size
}

func (t *Trie[T]) IsEmpty() bool {
	return t.size == 0
}

func (t *Trie[T]) Contains(key string) bool {
	return t.get(t.root, key, 0) != nil
}

func (t *Trie[T]) Get(key string) *T {
	x := t.get(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

// Put inserts a key-value pair into the Trie. Removes the key if val is nil.
func (t *Trie[T]) Put(key string, val *T) {
	if val == nil {
		t.Delete(key)
		return
	}
	t.root = t.put(t.root, key, val, 0)
}

func (t *Trie[T]) Delete(key string) {
	t.root = t.delete(t.root, key, 0)
}

// KeysWithPrefix returns all keys that start with the given prefix.
func (t *Trie[T]) KeysWithPrefix(prefix string) []string {
	var results []string
	x := t.get(t.root, prefix, 0)
	t.collect(x, strings.Builder{}, prefix, &results)
	return results
}

// LongestPrefix returns the longest common prefix of all keys in the trie.
func (t *Trie[T]) LongestPrefix() string {
	var prefix strings.Builder
	x := t.root

	for x != nil {
		// Count the number of non-nil children
		count := 0
		var nextIndex int
		for i := 0; i < 256; i++ {
			if x.next[i] != nil {
				count++
				nextIndex = i // Record the index of the non-nil child
			}
		}

		// If there's more than one child or if this node contains a value, stop
		if count != 1 || x.val != nil {
			break
		}

		// Append the character corresponding to the single child and move deeper
		prefix.WriteByte(byte(nextIndex))
		x = x.next[nextIndex]
	}

	return prefix.String()
}

func (t *Trie[T]) get(x *node[T], key string, d int) *node[T] {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	c := key[d]
	return t.get(x.next[c], key, d+1)
}

func (t *Trie[T]) put(x *node[T], key string, val *T, d int) *node[T] {
	if x == nil {
		x = &node[T]{}
	}
	if d == len(key) {
		if x.val == nil {
			t.size++
		}
		x.val = val
		return x
	}
	c := key[d]
	x.next[c] = t.put(x.next[c], key, val, d+1)
	return x
}

func (t *Trie[T]) delete(x *node[T], key string, d int) *node[T] {
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.val != nil {
			t.size--
		}
		x.val = nil
	} else {
		c := key[d]
		x.next[c] = t.delete(x.next[c], key, d+1)
	}

	// Remove subtree if empty
	if x.val != nil {
		return x
	}
	for _, n := range x.next {
		if n != nil {
			return x
		}
	}
	return nil
}

// helper method for KeysWithPrefix.
func (t *Trie[T]) collect(x *node[T], prefix strings.Builder, base string, results *[]string) {
	if x == nil {
		return
	}
	if x.val != nil {
		*results = append(*results, base+prefix.String())
	}
	for c := 0; c < 256; c++ {
		if x.next[c] != nil { // Only recurse if child exists
			cur_prefix := prefix.String()
			prefix.WriteByte(byte(c))
			t.collect(x.next[c], prefix, base, results)
			// Undo the last added character by resetting the length
			prefix.Reset()
			prefix.WriteString(cur_prefix) // Restore previous content
		}
	}
}
