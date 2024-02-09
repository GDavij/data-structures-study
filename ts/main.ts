import { Stack } from "./dsts/stack";
import { Queue } from "./dsts/queue";

var stack = new Stack<number>();

stack.Add(123);
console.log(stack.count);