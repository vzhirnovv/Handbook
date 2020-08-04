### Arrays

The type [n]T is an array of n values of type T.

The expression

```
var a [10]int
```

declares a variable a as an array of ten integers.

An array's length is part of its type, so **arrays cannot be resized**.

### Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```
a[low : high]
```

This selects a half-open range which **includes the first element, but excludes the last one**.

The following expression creates a slice which includes elements 1 through 3 of a:

```
a[1:4]
```

A slice **does not store any data**, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

### Slice literals

A slice literal is like an array literal without the length.

This is an array literal:

```
[3]bool{true, true, false}
```

And this creates the same array as above, then builds a slice that references it:

```
[]bool{true, true, false}
```

### Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

```
var a [10]int
```

these slice expressions are equivalent:

```
a[0:10]
a[:10]
a[0:]
a[:]
```

### Slice length and capacity

A slice has both a _length_ and a _capacity_.

The _length_ of a slice is the number of elements it contains.

The _capacity_ of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The _length_ and _capacity_ of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

_Note: basically, when you slice a slice, the array that slice is referencing still exists in the memory, which is why slicing and extending a slice preserves data, and going out of slice bounds throws an error._

### Nil slices

The zero value of a slice is **nil**.

A nil slice has a length and capacity of 0 and has no underlying array.
