package main

import (
	dst "data-structures/dst"
	"fmt"
)

func testQueue() {
	queue := dst.CreateQueue[int32]()
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

func main() {
	//testQueue()
}
