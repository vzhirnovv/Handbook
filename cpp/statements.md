# Statements

## Selection statements (if, else)

```cpp
if (condition) {
  // code block
} else if (condition)
  // one line of code
else 
  // one line of code
```

## Switch statement

Check for a value among a number of possible constant expressions.

```cpp
switch(expression) {
  case const1:
    ...
    break;
  case const2:
    ...
    break;
  default:
    ...
}
```

## Iteration statements

### while 

Repeats `statement` until `condition` is met. Evaluates `condition` before executing statement. Reevaluates `condition` after executing `statement`.

```cpp
int main() {
  int n = 10;

  while (n > 0) {
    cout << n;
    --n;
  }
}
```

### do while

Repeats `statement` until `condition` is met. Evaluates `condition` after executing statement.

```cpp
  do {
    // statement
  } while (condition);

```

Usually preferred over a while loop when the statement needs to be executed at least once.

### for

```cpp
for (init; condition; change) statement;
```

### Range-based for

```cpp
for (declaration: range) statement;
```

## Jump statements

### break

Forcefully exits the loop.

### continue

Skips the rest of the loop statement, jumps back to loop start.

### goto

Makes an absolute jump to another point in the prorgam. Ignores nesting levels and does not cause stack unwinding. Careful use required.

