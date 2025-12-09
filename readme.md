# Advent of Code 2025

Using this as an opportunity to explore ocaml. Below will be notes intended for personal use

https://ocaml.org/docs/tour-of-ocaml

# Tooling

The following section covers tooling related to the following areas:

- switches
- dune
- utop

# Switches

A switch can be thought of as a venv for python. We use OPAM to create a switch.

A switch contains it's own:

- compiler
- library
- binaries

A global switch is accessible from anywhere in the system, while a local switch is tied to a specific directory.

## Useful Switch Commands

The following are useful commands.

- `opam switch list` lists all configured switches.
- `opam switch` lists all configured switches. Useful to confirm what the current switch is.
- `opam switch create <project name> <compiler version>` creates a switch with the specified compiler. The compiler version is required to be specified.
- `opam switch <name>` activates the switch called `<name>`

# Dune

Dune is a build system commonly used for ocaml. We can also use Dune to generate the initial project scaffolding and build the project.

## Useful Dune Commands

- `dune init proj <name>` creates the scaffolding for a new project
- `dune build` builds the project
- `dune exec <name>` executes the project we build(in this case `aoc`)
- `-w` is a flag for watch mode. This can be appended to both `build` and `exec`

# File Structure

The extension used by ocaml projects is `*.ml`

# lib, bin, and \_bin

`lib` and `bin` do not contain executables. Instead they contain the source code for library code and program code. Instead, binaries(alongside source code) is stored in `_bin`.

# utop

_utop(Universal Top Level)_ is a REPL to evaluate ocaml expressions. It can be initialized by the command `utop` and closed by either `ctrl+D` or the command `#quit;;`. When `#` is visible, it means that the repl is still waiting for input. To signify the end of an expression and begin evaluation, end the line with `;;`.

# Tour of ocaml

The following sections will contain notes taken from `a tour of ocaml`

## Comments

Both line and inline comments use the same notation, that being `(_ \<contents here\>_)

_example:_ `(* let () = print_endline "Hello, World!" *)`

# Values and Types

Everything in ocaml has a value and a corresponding type, even an expression that only contains literals. Anonymous expressions are depicted with `-`. These are anonymous because they are not bound to a variable. Below are some examples of values being evaluated in utop. Note that every value here is an anonymous expression because they are not bound to anything

```ocaml
# 6.28;;
- : float = 6.28

# "This is really disco!";;
- : string = "This is really disco!"

# 'a';; (* Note the single quotes *)
- : char = 'a'

# true;;
- : bool = true
```

ocaml has type inference allowing it to automatically determine the type of an each expression.

## Lists

- creation: to create a list use the `let` keyword
- to add a value to the front of a list, use the `cons` operator, indicated by `::`

```ocaml
utop # let u = ['a'];;
val u : char list = ['a']

utop # 9 :: u;;
- : int list = [9; 1; 2; 3; 4]
```

## Conditionals

`if-then-else` is used as an expression. It is the same as the ternary operator

```ocaml
utop # 2 * if 'h' = 'b' then 5 else 3;;
- : int = 6
```

## Bindings

The keyword `let` is used to assign a value to a name. This is considered _binding_ a value to a name. A bound value is considered a constant, not a variable. This is because the value is immutable.

**Note: there are mutable values called references.**

```ocaml
# let x = 50;;
val x : int = 50
```

Bindings can also be specified with a local scope by using the `let … = … in …` syntax. In the example below, `y` only lives within the expression. This can be verified by being unable to access y in the following line. Multiple locally-scoped variables can be chained in together.

```ocaml
# let y = 50 in y * y;;
- : int = 2500

# y;;
Error: Unbound value y

# let a = 1 in
  let b = 2 in
    a + b;;
- : int = 3
```

## Operators

Ocaml uses `=` for both assignment and comparison. The following are comparison operators:

- `=` structural equality(for immutable data, this is all that matters)
- `==` physical equality(for references, when we want to check if they point to the same memory)
- `<>` not structurally equal
- `!=` not physically equal

## Functions

Since everything, even functions, are values, we define functions using the `let` keyword. Reading the output below, we see that `doubleMe` is a function that takes in an int and returns an int. A function can also have labeled paramters

```ocaml
utop # let doubleMe x = x * 2;;
val doubleMe : int -> int = <fun>

utop # let doubled = doubleMe 2;;
val doubled : int = 4

(* labelled example *)
# String.ends_with;;
- : suffix:string -> string -> bool = <fun>

# String.ends_with ~suffix:"less" "stateless";;
- : bool = true
```

### Anonymous Functions

Anonmous functions use the `fun` keyword. These are typically used in lambdas or other cases where we want to call a function immediately

```ocaml
# (fun x -> x * x) 50;;
- : int = 2500
```

### Multiple Paramters

Multiple paramters are handled with spaces

```ocaml
# let cat a b = a ^ " " ^ b;;
val cat : string -> string -> string = <fun>
```

### Partial Application

Partial application is when we call a function that takes multiple values, but do not provide multiple values. Instead of providing a value, the function will provide another function that requires the missing values as input.

This is an example of partial application

```ocaml
utop # let animal_say animal say = animal ^ " says " ^ say;;
val animal_say : string -> string -> string = <fun>

utop # animal_say;;
- : string -> string -> string = <fun>

utop # animal_say "dog" "woof";;
- : string = "dog says woof"

utop # let cat_says = animal_say "cat";;
val cat_says : string -> string = <fun>

utop # cat_says;;
- : string -> string = <fun>

utop # cat_says "meow";;
- : string = "cat says meow"
```

### Higher Order Functions

A higher-order function is a function that expects another function as a parameter. This is commonly seen with map, filter, and reduce functions.

```ocaml
# List.map (fun x -> x * x) [0; 1; 2; 3; 4; 5];;
- : int list = [0; 1; 4; 9; 16; 25]
```

### Recursive functions

A recursive function is any function that calls itself from it's body. These functions must be declared using the let rec ... = ...`format where`rec`indicates`recursion`. In Ocaml, recursion is provided over loops whenever possible

```ocaml
# let rec range lo hi =
    if lo > hi then
      []
    else
      lo :: range (lo + 1) hi;;
val range : int -> int -> int list = <fun>

# range 2 5;;
- : int list = [2; 3; 4; 5]
```

## Side Effects

Operating system level input/output is done via functions, typically `read_line` and `print_endline`

- `read_line` reads standard input and returns it as a string when end of line is reached.
- `print_endline` prints a string to standard output followed by an end of line.

Passing or returning data is not a requirement of either of these functions. If nothing is used, a placehold called `unit` that is referred to as `()` will appear in the sigernature instead.The `unit` type often indicates side-effects, but this is not always true.

# Styling guidelines

- Do not use dashes. Use underscores instead.
- Prefer recursion over loops.
