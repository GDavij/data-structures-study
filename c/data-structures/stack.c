#include "./stack.h"
#include "stdlib.h"
#include "stdio.h"

int stack_count(void* self)
{
  stack_t* this = (stack_t*)self;
  if (this->head == 0)
  {
    return 0;
  }
  stack_node_t* node = (stack_node_t*)this->head;
  int count = 0;
  while (node != 0)
  {
    count++;
    node = node->prev;
  }

  return count;
}


stack_t stack_c()
{
  stack_t stack = {
    .head = 0,
    .count = &stack_count
  };

  return stack;
}

void stack_put(void* self, void* data)
{
  stack_t* this = (stack_t*)self; 
  stack_node_t* node = malloc(sizeof(stack_node_t));
  stack_node_t node_data = {.data = malloc(sizeof(data)), .prev = 0};
  node_data.data = data;
  *node = node_data;

  if (this->head == 0)
  {
    this->head = node;
  } else {
    node->prev = this->head;
    this->head = node;
  }
}

evaluation_result stack_take(void* self)
{
  stack_t* this = (stack_t*)self;
  if (this->head == 0)
  {
    evaluation_result error_result = {0, 0};
    return error_result;
  }

  stack_node_t* node = (stack_node_t*)this->head;
  void* node_eval = node->data;
    
  evaluation_result result = {1, node_eval};
  if (node->prev == 0)
  {
    this->head = 0;
    return result;
  }

  stack_node_t* prev = node->prev;
  this->head = prev;

  return result;
}
