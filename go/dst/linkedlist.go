package dst

type linkedNode[T comparable] struct {
	value T
	next  *linkedNode[T]
}

type LinkedList[T comparable] struct {
	head  *linkedNode[T]
	Count int32
}

func LinkedListCreate[T comparable]() LinkedList[T] {
	return LinkedList[T]{
		head:  nil,
		Count: 0,
	}
}

func LinkedListAppend[T comparable](self *LinkedList[T], value T) {
	node := new(linkedNode[T])
	*node = linkedNode[T]{
		value: value,
		next:  nil,
	}

	if self.head == nil {
		self.head = node
		self.Count++
		return
	}

	current := self.head
	if current.next == nil {
		current.next = node
		self.Count++
		return
	}

	for current.next != nil {
		current = current.next
	}

	current.next = node
	self.Count++
}

func LinkedListRemove[T comparable](self *LinkedList[T], value T) *T {
	current := self.head
	result := new(T)

	if current.value == value {
		*result = current.value
		self.head = current.next
		self.Count--
		return result
	}

	previous := current
	current = current.next

	for current.next != nil && current.value != value {
		previous = current
		current = current.next
	}

	// Sure this algorithm can be optimized and better
	if current.value == value {
		*result = current.value
		previous.next = current.next
		self.Count--
		return result
	}

	return nil
}

func LinkedListContains[T comparable](self *LinkedList[T], value T) bool {
	if self.head == nil {
		return false
	}

	current := self.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}

	return false
}

type LinkedListSearchFilter[T comparable] func(T) bool

func LinkedListFirstOrDefault[T comparable](self *LinkedList[T], filter LinkedListSearchFilter[T]) *T {
	if self.head == nil {
		return nil
	}

	current := self.head
	for current != nil {
		isValid := filter(current.value)
		if isValid {
			result := new(T)
			*result = current.value
			return result
		}
		current = current.next
	}

	return nil
}

func LinkedListWhere[T comparable](self *LinkedList[T], filter LinkedListSearchFilter[T]) LinkedList[T] {
	if self.head == nil {
		return *self
	}

	filteredLinkedList := LinkedListCreate[T]()
	current := self.head

	for current != nil {
		isValid := filter(current.value)
		if isValid {
			LinkedListAppend(&filteredLinkedList, current.value)
		}

		current = current.next
	}

	return filteredLinkedList
}
