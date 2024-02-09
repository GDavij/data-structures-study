typedef struct 
{
  void* data;
  void* prev;
} stack_node_t;


typedef struct 
{
  stack_node_t* head;
  int (*count)(void* self);
} stack_t;

typedef struct
{
  int is_evaluated;
  void* eval;
} evaluation_result;

stack_t stack_c();

int stack_count(void* self);

void stack_put(void* self, void* data);

evaluation_result stack_take(void* self);
