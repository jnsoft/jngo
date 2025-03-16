package lists

type Ordered interface {
	~int | ~float64 | ~string // Add other ordered types if needed
}

// singly-linked list
type SingleLinkedNode[T Ordered] struct {
	Val  T
	Next *SingleLinkedNode[T]
}

// merges ordered lists to ordered list
func MergeKLists[T Ordered](lists []*SingleLinkedNode[T]) *SingleLinkedNode[T] {
	if lists == nil {
		return nil
	}

	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	return merge(MergeKLists(lists[0:mid]), MergeKLists(lists[mid:]))
}

func merge[T Ordered](a *SingleLinkedNode[T], b *SingleLinkedNode[T]) *SingleLinkedNode[T] {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if a.Val <= b.Val {
		a.Next = merge(a.Next, b)
		return a
	} else {
		b.Next = merge(a, b.Next)
		return b
	}
}
