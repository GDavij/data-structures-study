package main

import (
	dst "data-structures/dst"
	"data-structures/models"
	"fmt"
)

func testQueue() {
	queue := dst.QueueCreate[int32]()
	dst.QueueEnqueue(&queue, 1)
	dst.QueueEnqueue(&queue, 2)
	dst.QueueEnqueue(&queue, 3)
	dst.QueueEnqueue(&queue, 4)
	dst.QueueEnqueue(&queue, 5)
	dst.QueueEnqueue(&queue, 6)
	dst.QueueEnqueue(&queue, 7)
	dst.QueueEnqueue(&queue, 8)
	dst.QueueEnqueue(&queue, 9)
	dst.QueueEnqueue(&queue, 10)
	dst.QueueEnqueue(&queue, 11)

	fmt.Println(fmt.Sprintln("queue count: ", queue.Count))

	lastOnQueue := dst.QueueDequeue(&queue)
	for lastOnQueue != nil {
		fmt.Println(fmt.Sprint("value: ", *lastOnQueue))
		fmt.Println(fmt.Sprintln("count: ", queue.Count))
		lastOnQueue = dst.QueueDequeue(&queue)
	}
}

func testStack() {
	stack := dst.StackCreate[int32]()

	dst.StackPush(&stack, 1)
	dst.StackPush(&stack, 2)
	dst.StackPush(&stack, 3)
	dst.StackPush(&stack, 4)
	dst.StackPush(&stack, 5)
	dst.StackPush(&stack, 6)
	dst.StackPush(&stack, 7)
	dst.StackPush(&stack, 8)
	dst.StackPush(&stack, 9)
	dst.StackPush(&stack, 10)
	dst.StackPush(&stack, 11)
	dst.StackPush(&stack, 12)

	fmt.Println(fmt.Sprintln("count: ", stack.Count))

	top := dst.StackPop(&stack)
	for top != nil {
		fmt.Println(fmt.Sprint("value: ", *top))
		fmt.Println(fmt.Sprintln("count: ", stack.Count))
		top = dst.StackPop(&stack)
	}
}

func testLinkedList() {
	ll := dst.LinkedListCreate[int32]()

	fmt.Println(fmt.Sprintf("count: %d", ll.Count))

	dst.LinkedListAppend(&ll, 123) //1
	dst.LinkedListAppend(&ll, 456) //2
	dst.LinkedListAppend(&ll, 789) //3
	dst.LinkedListAppend(&ll, 101) //4
	dst.LinkedListAppend(&ll, 112) //5
	dst.LinkedListAppend(&ll, 131) //6
	dst.LinkedListAppend(&ll, 415) //7
	dst.LinkedListAppend(&ll, 161) //8
	dst.LinkedListAppend(&ll, 718) //9

	fmt.Println(fmt.Sprintf("count: %d", ll.Count))
	firstValue := dst.LinkedListRemove(&ll, 123)

	fmt.Println(fmt.Sprintf("first value: %d", *firstValue))

	midValue := dst.LinkedListRemove(&ll, 112)
	fmt.Println(fmt.Sprintf("mid value: %d", *midValue))

	lastValue := dst.LinkedListRemove(&ll, 718)
	fmt.Println(fmt.Sprintf("last value value: %d", *lastValue))
	fmt.Println(fmt.Sprintf("count: %d", ll.Count))

	fmt.Println(fmt.Sprintf("count: %d", ll.Count))

	firstValueThatIsContained := dst.LinkedListContains(&ll, 456)
	fmt.Println(fmt.Sprintf("First Value Is Contained: %t", firstValueThatIsContained))

	midValueThatIsContained := dst.LinkedListContains(&ll, 101)
	fmt.Printf("Mid Value Is Contained: %t\n", midValueThatIsContained)

	lastValueIsContained := dst.LinkedListContains(&ll, 161)
	fmt.Printf("Last Value That Is Contained: %t\n", lastValueIsContained)

	valueThatIsNotContained := dst.LinkedListContains(&ll, 123512512)
	fmt.Printf("Value that is not contained (should print false): %t\n", valueThatIsNotContained)
}

func testLinkedList2() {
	ll := dst.LinkedListCreate[models.User]()

	dst.LinkedListAppend(&ll, models.User{
		Id:     1,
		Name:   "John Doe",
		Stacks: dst.LinkedListCreate[string](),
	})

	dst.LinkedListAppend(&ll, models.User{
		Id:     2,
		Name:   "Marie Curie",
		Stacks: dst.LinkedListCreate[string](),
	})
	dst.LinkedListAppend(&ll, models.User{
		Id:     3,
		Name:   "Grace Hopper",
		Stacks: dst.LinkedListCreate[string](),
	})

	fmt.Println("Subject: must find the 3 entities of the list by id")
	johnDoe := dst.LinkedListFirstOrDefault(&ll, func(u models.User) bool {
		return u.Id == 1
	})

	fmt.Printf("First User: %s\n\n", johnDoe.Name)

	marieCurie := dst.LinkedListFirstOrDefault(&ll, func(u models.User) bool {
		return u.Id == 2
	})

	fmt.Printf("Mid User: %s\n\n", marieCurie.Name)

	graceHoper := dst.LinkedListFirstOrDefault(&ll, func(u models.User) bool {
		return u.Id == 3
	})

	fmt.Printf("Last User: %s\n\n", graceHoper.Name)

	nilUser := dst.LinkedListFirstOrDefault(&ll, func(u models.User) bool {
		return u.Id == 100
	})

	fmt.Printf("Nil User (must print a 0x0 pointer): %p\n\n", nilUser)

	fmt.Println("Subject: Must get only entities between 2 and 3 as a linked List")

	entitiesWithId2And3LinkedList := dst.LinkedListWhere(&ll, func(u models.User) bool {
		return u.Id >= 2 && u.Id <= 3
	})

	currentIteratorValue := dst.LinkedListFirstOrDefault(&entitiesWithId2And3LinkedList, func(u models.User) bool { return true })
	for currentIteratorValue != nil {
		fmt.Printf("Entity Found. Id: %d, Name: %s\n", currentIteratorValue.Id, currentIteratorValue.Name)
		dst.LinkedListRemove(&entitiesWithId2And3LinkedList, *currentIteratorValue)
		currentIteratorValue = dst.LinkedListFirstOrDefault(&entitiesWithId2And3LinkedList, func(u models.User) bool { return true })
	}
}

func main() {
	//testQueue()
	// testStack()
	// testLinkedList()
	testLinkedList2()
}
