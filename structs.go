package structs

type (
	Stack struct {
		head   *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Len() int {
	return this.length
}

func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.head.value
}

func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.head
	this.head = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value interface{}) {
	n := &node{value, this.head}
	this.head = n
	this.length++
}
