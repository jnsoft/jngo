package ds

type Color bool

const (
	RED   Color = false
	BALCK Color = true
)

// BST helper node data type
type Node[K comparable, V any] struct {
	Key    K
	Value  V
	Color  Color
	Left   *Node[K, V]
	Right  *Node[K, V]
	Parent *Node[K, V]
	N int
}

type RedBlackTree[K comparable, V any] struct {
	Root *Node[K, V]
}

func NewNode[K comparable, V any](key K, value V, color Color) *Node[K, V] {
	return &Node[K, V]{Key: key, Value: value, Color: color}
}

// Node helper functions

// isRed = false if x is null
func (Node) IsRed() bool {
	if x == nil {
		return false
	}
	return x.Color == RED
}

// number of node in subtree rooted at x, 0 if x is null
func (Node) Size() int
{
	if (x == null)
		return 0;
	return x.N;
}
