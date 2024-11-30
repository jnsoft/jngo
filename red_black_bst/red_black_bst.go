package red_black_dst

import (
	"errors"

	"github.com/jnsoft/jngo/misc"
)

type Color bool

const (
	RED   Color = false
	BLACK Color = true
)

// BST helper node data type
type Node[K misc.Ordered, V any] struct {
	key    K
	value  V
	color  Color
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
	n      int
}

type RedBlackTree[K misc.Ordered, V any] struct {
	root *Node[K, V]
}

func NewRedBlackTree[K misc.Ordered, V any]() *RedBlackTree[K, V] {
	return &RedBlackTree[K, V]{}
}

// return number of key-value pairs in this symbol table
func (t *RedBlackTree[K, V]) Size() int {
	return size(t.root)
}

// height of tree (1-node tree has height 0)
func (t *RedBlackTree[K, V]) Height(int) int {
	return height(t.root)
}

// is this symbol table empty?
func (t *RedBlackTree[K, V]) IsEmpty() bool {
	return t.root == nil
}

// value associated with the given key; null if no such key
func (t *RedBlackTree[K, V]) Get(key K) (V, error) {
	return get(t.root, key)
}

// is there a key-value pair with the given key?
func (t *RedBlackTree[K, V]) Contains(key K) bool {
	return contains(t.root, key)
}

// insert the key-value pair; overwrite the old value with the new value if the key is already present
func (t *RedBlackTree[K, V]) Put(key K, value V) {
	t.root = put(t.root, key, value)
	t.root.color = BLACK
}

// delete the key-value pair with the minimum key
func (t *RedBlackTree[K, V]) DeleteMin() error {

	if t.IsEmpty() {
		return errors.New("BST underflow")
	}

	// if both children of root are black, set root to red
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = RED
	}

	t.root = deleteMin(t.root)
	if !t.IsEmpty() {
		t.root.color = BLACK
	}
	return nil
}

// delete the key-value pair with the maximum key
func (t *RedBlackTree[K, V]) DeleteMax() error {

	if t.IsEmpty() {
		return errors.New("BST underflow")
	}

	// if both children of root are black, set root to red
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = RED
	}

	t.root = deleteMax(t.root)
	if !t.IsEmpty() {
		t.root.color = BLACK
	}
	return nil
}

// delete the key-value pair with the given key
func (t *RedBlackTree[K, V]) Delete(key K) {
	if !t.Contains(key) {
		// return errors.New("symbol table does not contain $s", s)
		return
	}

	// if both children of root are black, set root to red
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = RED
	}

	t.root = delete(t.root, key)
	if !t.IsEmpty() {
		t.root.color = BLACK
	}
}

// the smallest key; nil if no such key
func (t *RedBlackTree[K, V]) Min() (K, error) {
	if t.IsEmpty() {
		var zero K
		return zero, errors.New("tree is empty")
	}
	return min(t.root).key, nil
}

// the largest key; null if no such key
func (t *RedBlackTree[K, V]) Max() (K, error) {
	if t.IsEmpty() {
		var zero K
		return zero, errors.New("tree is empty")
	}
	return max(t.root).key, nil
}

func NewNode[K misc.Ordered, V any](key K, value V, color Color, n int) *Node[K, V] {
	return &Node[K, V]{key: key, value: value, color: color, n: n}
}

// ------------ Node helper functions ----------------------------

// isRed = false if x is null
func isRed[K misc.Ordered, V any](x *Node[K, V]) bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

// number of node in subtree rooted at x, 0 if x is null
func size[K misc.Ordered, V any](x *Node[K, V]) int {
	if x == nil {
		return 0
	}
	return x.n
}

func size2[K misc.Ordered, V any](x *Node[K, V]) int {
	if x == nil {
		return 0
	}
	return 1 + size(x.left) + size(x.right)
}

func height[K misc.Ordered, V any](x *Node[K, V]) int {
	if x == nil {
		return -1
	}
	return 1 + misc.Max(height[K, V](x.left), height[K, V](x.right))
}

// value associated with the given key in subtree rooted at x; null if no such key
func get[K misc.Ordered, V any](x *Node[K, V], key K) (V, error) {
	for x != nil {
		if key < (x.key) {
			x = x.left
		} else if key > x.key {
			x = x.right
		} else {
			return x.value, nil
		}
	}
	var zero V
	return zero, errors.New("key not found")
}

// is there a key-value pair with the given key in the subtree rooted at x?
func contains[K misc.Ordered, V any](x *Node[K, V], key K) bool {
	for x != nil {
		if key < x.key {
			x = x.left
		} else if key > x.key {
			x = x.right
		} else {
			return true
		}
	}
	return false
}

// insert the key-value pair in the subtree rooted at h
func put[K misc.Ordered, V any](x *Node[K, V], key K, value V) *Node[K, V] {
	if x == nil {
		return NewNode(key, value, RED, 1)
	}

	if key < x.key {
		x.left = put(x.left, key, value)
	} else if key > x.key {
		x.right = put(x.right, key, value)
	} else {
		x.value = value
	}

	// fix-up any right-leaning links
	if isRed(x.right) && !isRed(x.left) {
		x = rotateLeft(x)
	}
	if isRed(x.left) && isRed(x.left.left) {
		x = rotateRight(x)
	}
	if isRed(x.left) && isRed(x.right) {
		flipColors(x)
	}

	x.n = size(x.left) + size(x.right) + 1
	return x
}

// delete the key-value pair with the minimum key rooted at h
func deleteMin[K misc.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	if x.left == nil {
		return nil
	}

	if !isRed(x.left) && !isRed(x.left.left) {
		x = moveRedLeft(x)
	}

	x.left = deleteMin(x.left)
	return balance(x)
}

// delete the key-value pair with the maximum key rooted at x
func deleteMax[K misc.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	if isRed(x.left) {
		x = rotateRight(x)
	}

	if x.right == nil {
		return nil
	}

	if !isRed(x.right) && !isRed(x.right.left) {
		x = moveRedRight(x)
	}

	x.right = deleteMax(x.right)

	return balance(x)
}

// delete the key-value pair with the given key rooted at x
func delete[K misc.Ordered, V any](x *Node[K, V], key K) *Node[K, V] {
	if x == nil {
		return nil
	}
	if key < x.key {
		if !isRed(x.left) && !isRed(x.left.left) {
			x = moveRedLeft(x)
		}
		x.left = delete(x.left, key)
	} else {
		if isRed(x.left) {
			x = rotateRight(x)
		}
		if key == x.key && x.right == nil {
			return nil
		}
		if !isRed(x.right) && !isRed(x.right.left) {
			x = moveRedRight(x)
		}
		if key == x.key {
			h := min(x.right)
			x.key = h.key
			x.value = h.value
			x.right = deleteMin(x.right)
		} else {
			x.right = delete(x.right, key)
		}
	}
	return balance(x)
}

// ------------ Red black helpers ----------------------------

func rotateLeft[K misc.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	h := x.right
	x.right = h.left
	h.left = x
	h.color = x.color
	x.color = RED
	h.n = x.n
	x.n = 1 + size(x.left) + size(x.right)
	return h
}

func rotateRight[K misc.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	h := x.left
	x.left = h.right
	h.right = x
	h.color = x.color
	x.color = RED
	h.n = x.n
	x.n = 1 + size(x.left) + size(x.right)
	return h
}

func flipColors[K misc.Ordered, V any](x *Node[K, V]) {
	x.color = RED
	x.left.color = BLACK
	x.right.color = BLACK
}

// Assuming that h is red and both h.left and h.left.left
// are black, make h.left or one of its children red.
func moveRedLeft[K misc.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	flipColors(x)
	if isRed(x.right.left) {
		x.right = rotateRight(x.right)
		x = rotateLeft(x)
	}
	return x
}

// Assuming that h is red and both h.right and h.right.left
// are black, make h.right or one of its children red.
func moveRedRight[K misc.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	flipColors(x)
	if isRed(x.left.left) {
		x = rotateRight(x)
	}
	return x
}

// restore red-black tree invariant
func balance[K misc.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	if isRed(x.right) {
		x = rotateLeft(x)
	}
	if isRed(x.left) && isRed(x.left.left) {
		x = rotateRight(x)
	}
	if isRed(x.left) && isRed(x.right) {
		flipColors(x)
	}

	x.n = size(x.left) + size(x.right) + 1
	return x
}

// ------------ ordered symbol table methods ----------------------------

// the smallest key in subtree rooted at x; null if no such key
func min[K misc.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	if x.left == nil {
		return x
	} else {
		return min(x.left)
	}
}

// the largest key in the subtree rooted at x; null if no such key
func max[K misc.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	if x.right == nil {
		return x
	} else {
		return max(x.right)
	}
}

/*





       public RedBlackBST<TKey, TValue> Copy()
       {
           RedBlackBST<TKey, TValue> copy = new RedBlackBST<TKey, TValue>();
           foreach (TKey key in this.Keys)
               copy.Put(key, this.Get(key));
           return copy;
       }

       #endregion

       #region Red black helpers















       // the largest key less than or equal to the given key
       public TKey Floor(TKey key)
       {
           Node x = floor(root, key);
           if (x == null)
               return default(TKey);
           else return x.Key;
       }

       // the largest key in the subtree rooted at x less than or equal to the given key
       private Node floor(Node x, TKey key)
       {
           if (x == null) return null;
           int cmp = key.CompareTo(x.Key);
           if (cmp == 0)
               return x;
           if (cmp < 0)
               return floor(x.left, key);
           Node t = floor(x.right, key);
           if (t != null)
               return t;
           else
               return x;
       }

       // the smallest key greater than or equal to the given key
       public TKey Ceiling(TKey key)
       {
           Node x = ceiling(root, key);
           if (x == null) return default(TKey);
           else return x.Key;
       }

       // the smallest key in the subtree rooted at x greater than or equal to the given key
       private Node ceiling(Node x, TKey key)
       {
           if (x == null) return null;
           int cmp = key.CompareTo(x.Key);
           if (cmp == 0)
               return x;
           if (cmp > 0)
               return ceiling(x.right, key);
           Node t = ceiling(x.left, key);
           if (t != null)
               return t;
           else
               return x;
       }

       // the key of rank k
       public TKey Select(int k)
       {
           if (k < 0 || k >= Size)
               return default(TKey);
           Node x = select(root, k);
           return x.Key;
       }

       // the key of rank k in the subtree rooted at x
       private Node select(Node x, int k)
       {
           int t = size(x.left);
           if (t > k)
               return select(x.left, k);
           else if (t < k)
               return select(x.right, k - t - 1);
           else
               return x;
       }

       // number of keys less than key
       public int Rank(TKey key) { return rank(key, root); }

       // number of keys less than key in the subtree rooted at x
       private int rank(TKey key, Node x)
       {
           if (x == null) return 0;
           int cmp = key.CompareTo(x.Key);
           if (cmp < 0)
               return rank(key, x.left);
           else if (cmp > 0)
               return 1 + size(x.left) + rank(key, x.right);
           else
               return size(x.left);
       }

       #endregion

       // all of the keys, as an Iterable
       public IEnumerable<TKey> Keys { get { return keys(Min, Max); } }

       // the keys between lo and hi, as an Iterable
       public IEnumerable<TKey> keys(TKey lo, TKey hi)
       {
           Queue<TKey> queue = new Queue<TKey>();
           // if (isEmpty() || lo.compareTo(hi) > 0) return queue;
           keys(root, queue, lo, hi);
           return queue;
       }

       // add the keys between lo and hi in the subtree rooted at x to the queue
       private void keys(Node x, Queue<TKey> queue, TKey lo, TKey hi)
       {
           if (x == null) return;
           int cmplo = lo.CompareTo(x.Key);
           int cmphi = hi.CompareTo(x.Key);
           if (cmplo < 0)
               keys(x.left, queue, lo, hi);
           if (cmplo <= 0 && cmphi >= 0)
               queue.Enqueue(x.Key);
           if (cmphi > 0)
               keys(x.right, queue, lo, hi);
       }

       // number keys between lo and hi
       public int size(TKey lo, TKey hi)
       {
           if (lo.CompareTo(hi) > 0)
               return 0;
           if (Contains(hi))
               return Rank(hi) - Rank(lo) + 1;
           else
               return Rank(hi) - Rank(lo);
       }

       #region integrety check of structure

       public bool Check()
       {
           StringBuilder sb = new StringBuilder();
           if (!isBST()) sb.Append("Not in symmetric order");
           if (!isSizeConsistent()) sb.Append("Subtree counts not consistent");
           if (!isRankConsistent()) sb.Append("Ranks not consistent");
           if (!is23()) sb.Append("Not a 2-3 tree");
           if (!isBalanced()) sb.Append("Not balanced");
           return isBST() && isSizeConsistent() && isRankConsistent() && is23() && isBalanced();
       }

       // does this binary tree satisfy symmetric order?
       // Note: this test also ensures that data structure is a binary tree since order is strict
       private bool isBST()
       {
           return isBST(root, default(TKey), default(TKey));
       }

       // is the tree rooted at x a BST with all keys strictly between min and max
       // (if min or max is null, treat as empty constraint)
       // Credit: Bob Dondero's elegant solution
       private bool isBST(Node x, TKey min, TKey max)
       {
           if (x == null) return true;
           if (min != null && x.Key.CompareTo(min) <= 0) return false;
           if (max != null && x.Key.CompareTo(max) >= 0) return false;
           return isBST(x.left, min, x.Key) && isBST(x.right, x.Key, max);
       }

       // are the size fields correct?
       private bool isSizeConsistent() { return isSizeConsistent(root); }
       private bool isSizeConsistent(Node x)
       {
           if (x == null) return true;
           if (x.N != size(x.left) + size(x.right) + 1) return false;
           return isSizeConsistent(x.left) && isSizeConsistent(x.right);
       }

       // check that ranks are consistent
       private bool isRankConsistent()
       {
           for (int i = 0; i < Size; i++)
               if (i != Rank(Select(i))) return false;
           foreach (TKey key in Keys)
               if (key.CompareTo(Select(Rank(key))) != 0) return false;
           return true;
       }

       // Does the tree have no red right links, and at most one (left)
       // red links in a row on any path?
       private bool is23() { return is23(root); }
       private bool is23(Node x)
       {
           if (x == null) return true;
           if (isRed(x.right)) return false;
           if (x != root && isRed(x) && isRed(x.left))
               return false;
           return is23(x.left) && is23(x.right);
       }

       // do all paths from root to leaf have same number of black edges?
       private bool isBalanced()
       {
           int black = 0;     // number of black links on path from root to min
           Node x = root;
           while (x != null)
           {
               if (!isRed(x)) black++;
               x = x.left;
           }
           return isBalanced(root, black);
       }

       // does every path from the root to a leaf have the given number of black links?
       private bool isBalanced(Node x, int black)
       {
           if (x == null) return black == 0;
           if (!isRed(x)) black--;
           return isBalanced(x.left, black) && isBalanced(x.right, black);
       }

       #endregion

       public IEnumerator<TValue> GetEnumerator()
       {
           IEnumerable<TKey> keys = Keys;
           foreach (TKey key in Keys)
               yield return Get(key);
       }

       IEnumerator IEnumerable.GetEnumerator() { return GetEnumerator(); }
   }

*/
