# Operators

## Assignmment operator (=)

Always takes place from right to left. 

Assignment is an expression that can be evaluated <=> the assignment itself has a value

For example, in this snippet:

```cpp
int x, y;
y = 2 + (x = 5)
```

`y` is assigned the result of adding 2 and x, which in turn has the value 5 assigned to it. Considering assignment takes place from right to left, `y` evaluates to 7.

The following expression in also valid:

```cpp
int x, y, z;
x = y = z = 5;
```

with `x`, `y`, and `z` evaluating to 5.

## Arithmetic operators

* `+` - addition
* `-` - subtraction
* `*` - multiplication
* `/` - division
* `%` - modulo

## Iteration operators

* `++` - increment
* `--` - decrement

Can be used as a prefix or a suffix. If used as a prefix, evaluation happens after the variable with the prefix is incremented. If used as a suffix, evaluation takes place first, with incremention afterwards.

For example:
```cpp
int x = 3;
int y = ++x; // y = 4, x = 4
int z = x++; // z = 5, x = 4
```

## Relational and comparison operators

* `==` - equal to
* `!=` - not equal to
* `<` - less than
* `>` - greater than
* `<=` - less or equal to
* `>=` - greater or equal to

Evaluates to a Boolean value (`true` or `false`).

## Logical operators

* `!` - NOT
* `&&` - AND
* `||` - OR

Evaluates to a Boolean value.

Can 'short-circuit' - when using `&&` and the left-hand expression is false, evaluate entire expression to false and stop, or, when using `||` and the left hand expression is true, evaluate entire expression to true and stop.

## Conditional ternary operator (?)

```cpp
condition ? result1 : result2
```

If `condition` is true, evaluate to `result1`, otherwise evaluate to `result2`

## Comma operator

Allows to 'chain' expressions, decreasing the amount of lines%
```cpp
a = (b = 3, b + 2) // a = 5
```

## Bitwise operators

* `&` - Bitwise AND
* `|` - Bitwise inclusive OR
* `^` - Bitwise exclusive OR
* `~` - Bit inversion
* `<<` - Left shift
* `>>` - Right shift

## Explicit type casting operator

Converts a value of one type to another type.

```cpp
int i, j;
float f = 3.14;
i = (int) f
// OR
j = int(f)
// both ways are valid.
```

## Compount assignment operators

Modifies the value of a variable by performing an operation on it. Same as assigning the result of the operation to the first operand.

Available operators:
```cpp
+=, -=, *-, /=, %=, >>=, <<=, &=, ^=, |=
```
