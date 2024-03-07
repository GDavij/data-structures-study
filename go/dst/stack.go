package dst

type stackNode[T any] struct {
	value    T
	previous *stackNode[T]
}

type Stack[T any] struct {
	top   *stackNode[T]
	Count int
}

func StackCreate[T any]() Stack[T] {
	return Stack[T]{top: nil, Count: 0}
}

func StackPush[T any](stack *Stack[T], value T) {
	node := new(stackNode[T])
	*node = stackNode[T]{value: value, previous: nil}

	if stack.Count == 0 {
		stack.top = node
		stack.Count++
		return
	}

	node.previous = stack.top
	stack.top = node
	stack.Count++
}

func StackPop[T any](stack *Stack[T]) *T {
	if stack.Count == 0 {
		return nil
	}

	value := new(T)
	*value = stack.top.value

	stack.top = stack.top.previous
	stack.Count--

	return value
}
