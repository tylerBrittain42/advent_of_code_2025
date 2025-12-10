# Advent of Code 2025

Using this as an opportunity to explore ocaml. Below will be notes intended for personal use

https://ocaml.org/docs/tour-of-ocaml

# utop

_utop(Universal Top Level)_ is a REPL to evaluate ocaml expressions. It can be initialized by the command `utop` and closed by either `ctrl+D` or the command `#quit;;`. When `#` is visible, it means that the repl is still waiting for input. To signify the end of an expression and begin evaluation, end the line with `;;`.

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

The extension used by ocaml projects is `*.ml`. There is no dedicated main function and a main.ml file is not required.

## lib, bin, and \_bin

`lib` and `bin` do not contain executables. Instead they contain the source code for library code and program code. Instead, binaries(alongside source code) is stored in `_bin`.

# Modules and the standard library
Some quick points:
- a module is a collection of named values
- identical names from diff modules don't clash
- the standard lib is a collection of diff modules.

Each OCaml file is compiled into a module. 