import { Stack } from "./dsts/stack";
import { Queue } from "./dsts/queue";
import { LinkedList } from "./dsts/linkedList";

const stack = new Stack<number>();

console.log("Stack: ")
stack.Add(123);
stack.Add(234);

console.log(`count: ${stack.count}`);


let element = stack.Remove();
console.log("element: " + element);

element = stack.Remove()
console.log("element: " + element);

console.log(`count: ${stack.count}`);


console.log("\nqueue: ")

const queue = new Queue<number>();

queue.enqueue(123);
queue.enqueue(234);
console.log("count: " + queue.count);

let elementQueue = queue.dequeue();
console.log("element: " + element);

element = queue.dequeue()
console.log("element: " + element);

console.log(`count: ${queue.count}`);



const ll = new LinkedList<number>();
console.log("Linked List: ");

console.log(`count: ${ll.count}`);

ll.append(10);
ll.append(50);
ll.append(90);
ll.append(61);

console.log("Find: ");

console.log("Last Element: " + ll.findFirst(61));

console.log("Medium Element: " + ll.findFirst(50));

console.log("First Element: " + ll.findFirst(10));


console.log("Count: " + ll.count);

console.log("Remove Node: ");

console.log("Last Element: " + ll.remove(61));

console.log("Medium Element: " + ll.remove(50));

console.log("First Element: " + ll.remove(10));

console.log("count: " + ll.count);