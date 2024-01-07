# monkey-interpreter

Following Thorsten Ball's book [Writing an Interpreter in Go](https://interpreterbook.com/). This repository contains an interpreter for the Monkey langauge written in Go.

![Monkey Logo](https://interpreterbook.com/img/monkey_logo-d5171d15.png)

### Core concepts this codebase covers

_lexer_ - converts fragments of the monkey programming language into tokens

_parser_ - syntactically analyses and convert tokens of the monkey programming language into an abstract syntax tree (AST) data structure

_evaluator_ - executes monkey source code in AST structure using a tree walking interpreter

_repl_ - a runtime command line read evaluate print loop for monkey source code

### Supported Language Features

|Feature|Lexer Support Implemented|Parser Support Implemented|Evaluator Support Implemented|
|-------|-------------------------|--------------------------|-----------------------------|
|Identifiers (ascii) |✅|✅|✅|
|Identifiers (unicode) | | | |
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
