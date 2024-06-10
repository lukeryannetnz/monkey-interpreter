# monkey-interpreter

Following Thorsten Ball's book [Writing an Interpreter in Go](https://interpreterbook.com/). This repository contains an interpreter for the Monkey langauge written in Go.

![coverage](https://raw.githubusercontent.com/lukeryannetnz/monkey-interpreter/badges/.badges/main/coverage.svg)

![Monkey Logo](https://interpreterbook.com/img/monkey_logo-d5171d15.png)

### Money language
Monkey is a programming language designed by Thorsten Ball for the purpose of writing an interpreter for it in his book [Writing an Interpreter in Go](https://interpreterbook.com/). The language is a simple programming language with a C-like syntax. The language is designed to be easy to write an interpreter for, but also to be powerful enough to write interesting programs in.

There are a number of builtin functions in the monkey language
puts - prints the arguments to the console
```monkey
puts("Hello World!")
>Hello World!
>null
```

len - returns the length of the argument
```monkey
let a = ["a", "b", "c"]
len(a)
> 3
```

first - returns the first element of the argument
```monkey
let a = ["a", "b", "c"]
first(a)
> a
```

last - returns the last element of the argument
```monkey
let a = ["a", "b", "c"]
last(a)
> c
```

rest - returns all but the first element of the argument
```monkey
let a = ["a", "b", "c"]
rest(a)
> [b, c]
```

push - returns a new array with the argument appended to the end
```monkey
let a = ["a", "b", "c"]
push(a, "d")
> [a, b, c, d]
```

### Getting Started
You can start the repl with the command `go run main.go`. This will start the monkey repl where you can enter monkey code and see the output.

You can run the tests with the command `go test ./...`. This will run all the tests in the project.

### Core concepts this codebase covers

_lexer_ - converts fragments of the monkey programming language into tokens

_parser_ - syntactically analyses and convert tokens of the monkey programming language into an abstract syntax tree (AST) data structure

_evaluator_ - executes monkey source code in AST structure using a tree walking interpreter

_repl_ - a runtime command line read evaluate print loop for monkey source code

### Supported Language Features

|Feature|Lexer Support Implemented|Parser Support Implemented|Evaluator Support Implemented|
|-------|-------------------------|--------------------------|-----------------------------|
|Identifiers (ascii) |✅|✅|✅|
|Integer literals |✅|✅|✅|
|Assignment operator |✅|✅|✅|
|Addition operator |✅|✅|✅|
|Subtraction operator |✅|✅|✅|
|Multiplication operator |✅|✅|✅|
|Division operator |✅|✅|✅|
|Bang operator |✅|✅|✅|
|Less than operator |✅|✅|✅|
|Greater than operator |✅|✅|✅|
|Equals operator |✅|✅|✅|
|Not equals operator |✅|✅|✅|
|Comma delimiter |✅|✅|✅|
|Semicolon delimiter |✅|✅|✅|
|Left parenthesis delimiter |✅|✅|✅|
|Right parenthesis delimiter |✅|✅|✅|
|Left brace delimiter |✅|✅|✅|
|Right brace delimiter |✅|✅|✅|
|Function literals |✅|✅| |
|Let keyword |✅|✅|✅|
|True keyword |✅|✅|✅|
|False keyword |✅|✅|✅|
|If keyword |✅|✅|✅|
|Else keyword |✅|✅|✅|
|Return keyword |✅|✅|✅|
|String literals |✅|✅|✅|
