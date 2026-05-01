# Monkey (Go)

An educational Monkey language implementation in Go, with working **tokenizer**, **lexer**, and a partially implemented **parser/AST**.

## Requirements

- Go `1.26.1` (as declared in `go.mod`)

## Run

```bash
go run .
```

The CLI opens a small REPL-style menu:

1. `Tokenizeing` (token classification for a single word)
2. `Lexing` (token stream for a statement/input line)
0. Exit

## Test

```bash
go test ./...
```

## What is implemented

- **Token system (`token/`)**
  - Token specs for keywords, identifiers, numeric literals, booleans, operators, delimiters, and braces/parentheses
  - Regex-based token matching with ordered priority
- **Lexer (`lexer/`)**
  - Character scanner with support for identifiers, integers, floats, punctuation, operators, and multi-char comparators (`==`, `!=`, `<=`, `>=`)
  - Iterator API (`All()`) for streaming tokens
- **AST (`ast/`)**
  - Program, statement, and expression node hierarchy
  - Nodes for `let`, `return`, identifier, integer, boolean, and prefix expressions
- **Parser (`parser/`)**
  - Parsing for `let` statements, `return` statements, identifiers, booleans, integers, and prefix expressions (`!`, `-`)
  - Basic parser error collection

## Current limitations

- Parser is not yet wired into the top-level CLI (`main.go` currently exposes tokenizing + lexing flows).
- Infix expression parsing and precedence handling are not finished.
- String literals are defined as a TODO and not yet supported.

## Repository layout

- `main.go`: entrypoint for the interactive menu
- `repl/`: interactive CLI prompts and input handling
- `token/`: token types/specs and token creation
- `lexer/`: lexical scanner
- `ast/`: AST node definitions
- `parser/`: parser implementation and parser tests

## Notes

- A pre-commit hook exists at `.githooks/pre-commit` to run `gofmt` on staged Go files.
