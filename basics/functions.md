# Functions

```
type name (params) {statements}
```

## Void functions

`void` is a special type that represents the absence of a value.

`void` can also be used instead of `params` to explicitly specify that the function does not have parameters.

```cpp
void func (void) {} // necessary in C, use empty params in C++
```

## Return value of main()

Execution of `main` can end normally without an explicit return statement. In this case, the compiler automatically ends the program with an *implicit return statement*:

```cpp
return 0;
```

When `main` returns 0, it means that the program ended successfully. There are other values that `main` can return, but this behaviour is usually not portable between platforms. Only the following types are unniversal:

* `0` - success
* `EXIT_SUCCESS` - success
* `EXIT_FAILURE` - fail


## Arguments by value vs by reference

```cpp
void func(int foo, char bar) // by value
void func(int& foo, char& bar) // by reference
```

When arguments are passed by value, the values of these arguments at the moment of function call are copied into the variables represented by function parameters.

When arguments are passed by reference, instead of a copy, the variable itself is passed, which means that any modification on the local variables of a function will change the variable that was assigned to the argument during function call.

### Efficiency and const references

Calling a function with parameters taken by value causes copies to be made, which can produce overhead when a large compound type is passed. This copy can be avoided by passing the parameters as references, but functions with reference parameters are percievedas functions that modify the arguments passed. To solve this problem, **const references** can be used:

```cpp
type func (const type& foo, const &type bar)
```

This notation forbids the function from being used 

