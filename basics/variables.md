# Variables and Types

## Identifiers

The **identifier** of a variable is, in simple terms, the 'name' of that variable.

C++ indentifiers can be a sequence of letters, digits, and underscore characters. Spaces, punctuation marks and symbols are disallowed. An identifier **cannot start** with a number.
Keep in mind, C++ is case sensitive, so be careful when referencing a variable.

*Note: in C++, an underscore usually indicates a private member variable in a class.*

## Fundamental data types

Fundamental (implemented by the language) types can be classified into 4 categories:

* **Character types**: represent characters and symbols
  - `char`: one byte in size.
  - `char16_t`: not smaller than char, at least 16 bits
  - `char32_t`: not smaller than `char16_t`, at least 32 bits 
  - `wchar_t`: largest supported character set
* **Integer types**: store whole number values, can be signed (supports negatives) or unsigned
  - `signed char`: one byte in size, ranges from -128 to 127 (signed) 
  - `<signed?> short int`: at least 16 bits in size, ranges from -32,768 to 23,767 (signed) or 0 to 65,535 (unsigned)
  - `<signed?> int`: at least 16 bits in size, can be 2 or 4 bytes, ranges accordingly
  - `<signed?> long int`: at least 32 bits in size, ranges accordingly
  - `<signed?> long long int`: at least 64 bits in size, ranges accordingly
* **Floating point types**: can represent real values with different levels of precision
  - `float`: 32 bit IEEE 754 single precision Floating Point Number (7 digits of precision)
  - `double`: 64 bit IEEE 754 double precision Floating Point Number (15 digits of precision)
  - `long double`: has the precision of at least `double`, precision can be extended based on the compiler in use
* **Boolean type**
  - `bool`: *true* or *false*, what did you expect?
* **Void type**
  - `void`: no storage
* **Null pointer**: points to nothing :)

## Declaring and initializing variables

```
<type> <name> [= <value> | (<value>) - C++ | {<value>} - C++11 and up]
```

## Type deduction

When a new variable is initialized, the compiler can figure out what the type of the variable is automatically by the initializer. For this, we can use `auto` or `decltype()`:

```c++
int foo = 4;
auto bar = foo; // int bar = foo;
decltype(foo) bar; // int bar
```
