# Advent of Code 2025
Using this as an opportunity to explore ocaml. Below will be notes intended for personal use

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

# Building and Compiling
## Dune
Dune is a build system commonly used for ocaml. We can also use Dune to generate the initial project scaffolding. To do this for a project called hello, you would use the following command: `opam exec -- dune init proj hello`.