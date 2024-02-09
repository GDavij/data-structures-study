#include "data-structures/stack.h"
#include "stdio.h"

int main()
{
  stack_t stack = stack_c();
  
  stack_put(&stack, "Hello World");
  stack_put(&stack, "Hello World");

  stack_take(&stack);
  stack_take(&stack);

  printf("count: %d", stack.count(&stack));
}
