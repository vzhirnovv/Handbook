### Functions

A function can take zero or more arguments.

In this example, add takes two parameters of type `int`.

Notice that the type comes **after** the variable name.

(For more about why types look the way they do, see the article on Go's declaration syntax.)

### Functions with omitted types.=

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

Take the following example:

`x int, y int` can be converted to `x, y int`
