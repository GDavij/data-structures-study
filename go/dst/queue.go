package dst

type queueNode[T any] struct {
	value    T
	previous *queueNode[T]
}

type Queue[T any] struct {
	start *queueNode[T]
	end   *queueNode[T]
	Count int32
}

func QueueCreate[T any]() Queue[T] {
	return Queue[T]{start: nil, end: nil, Count: 0}
}

func QueueEnqueue[T any](queue *Queue[T], value T) {
	node := new(queueNode[T])
	*node = queueNode[T]{value: value, previous: nil}

	if queue.Count == 0 {
		queue.start = node
		queue.end = node
	} else {
		start := queue.start
		start.previous = node
		queue.start = node
	}

	queue.Count++
}

func QueueDequeue[T any](queue *Queue[T]) *T {
	if queue.Count == 0 {
		return nil
	}

	value := new(T)

	if queue.Count == 1 {
		node := queue.end
		*value = node.value
		queue.start = nil
		queue.end = nil
		queue.Count--

		return value
	}

	end := queue.end
	*value = end.value
	queue.end = end.previous
	queue.Count--

	return value
}
