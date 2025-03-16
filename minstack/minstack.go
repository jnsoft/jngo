package minstack

// MinStack represents a stack that supports retrieving the minimum value.
// init: ms := &MinStack{}
type MinStack struct {
	stack []pair
}

type pair struct {
	value      int
	currentMin int
}

func NewMinStack() MinStack {
	return MinStack{
		stack: []pair{},
	}
}

func (ms *MinStack) Push(val int) {
	currentMin := val
	if len(ms.stack) > 0 && ms.stack[len(ms.stack)-1].currentMin < val {
		currentMin = ms.stack[len(ms.stack)-1].currentMin
	}
	ms.stack = append(ms.stack, pair{value: val, currentMin: currentMin})
}

func (ms *MinStack) Pop() {
	if len(ms.stack) > 0 {
		ms.stack = ms.stack[:len(ms.stack)-1]
	}
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		panic("stack is empty")
	}
	return ms.stack[len(ms.stack)-1].value
}

func (ms *MinStack) GetMin() int {
	if len(ms.stack) == 0 {
		panic("stack is empty")
	}
	return ms.stack[len(ms.stack)-1].currentMin
}
