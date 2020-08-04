# Constants

## Literals

`a = 5`, 5 is a *literal constant*

### Integer Numerals

```
75 - decimal
0113 - octal
0x4b - hexadecimal
```

```
75 - int
75u - unsigned int
75l - long
75ul - unsigned long
75lu - unsigned long
```

### Floating point numerals

``` 
3.14159
6.02e23 - same as 6.02 * 10^23
1.6e-19 - same as 1.6 * 10^-19
3.0
```

```
3.14159L - long double
6.02e23f - float
```

### Character and string literals

```
'z'
'p'
"Hello, World!"
```

`\` is an character escape: use it to represent special characters or force some symbols

Several string literals can be concatenated:

```
"this is" "a very long string" " split into a few"   " smaller strings"
```
is the same as:
```
"this isa very long string split into a few smaller strings"
```

Strings can also be spread across multiple lines using a *line-continuation* character

```
x = "string expressed in \
two lines"
```
is the same as:
```
x = "string expressed in two lines"
```

Type can be specified using a prefix (personally, I don't consider that a good practice, a better solution would be to explicitly declare the type of a variable)

```
u"string" - char16_t
U"string" - char32_t
L"string" - wchar_t
```

## Typed constant expressions

```c++
const double pi = 3.14159;
const char tab = '\t';
```

Afterwards, these variables can be used as literals.

## Preprocessor definitions

```c++
#define PI 3.14159
#define NEWLINE '\n'
```

Afterwards, any occurence of a defined indentifier in the code is interpreted as a replacement (any sequence of characters until the end of the line). This happens before compilation via preprocessing, so there is not way to check the validity or type of the syntax.
