### Exported names

In Go, a name is exported if it begins with a **capital** letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the math package.

_pizza_ and _pi_ do not start with a capital letter, so they are not exported.

When importing a package, you can refer **only** to its exported names. Any "unexported" names are not accessible from outside the package.
