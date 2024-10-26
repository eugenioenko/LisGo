# LisGo (WIP)

LisGo is a minimalist programming language built with Go. It serves as a foundation to explore and experiment with the complexities of functional programming language interpretation. Initially inspired by Lisp, LisGo will gradually diverge as it evolves.

## > [Try it out in Web Assembly playground!](https://eugenioenko.github.io/LisGo/live)

## Building the project

> go build

## Building web assembly binary

> make build

which will execute:

> GOOS=js GOARCH=wasm go build -o live/lisgo.wasm wasm/lisgo.go

## Running the project

> lisgo exec [filename] will execute the script

> lisgo eval [code] will execute the code passed as argument

## Running playground locally

WebAssembly's instantiateStreaming method requires CORS to be enabled, so a local server is necessary to run the playground. Set it up as follows:

> npm install http-server
> cd live
> http-server

## Testing

Testing can be done by running:

> make test

This command will execute all the tests located in the tests folder by running:

> go test ./tests

## LisGo Language Documentation

This document provides an overview of the expressions and syntax used in LisGo. Each expression is explained with examples to demonstrate its usage.

### Printing Expressions

> print

The print expression outputs a value or result to the console.

Syntax:

```
(print expression)
```

Examples:

```
(print (+ 1 2 3)) ; Outputs the sum of 1, 2, and 3
(print "hello world") ; Outputs "hello world"
(print (print "yes")) ; Prints "yes" and returns nil
```

### Math Expressions

> - and -

These operators perform basic arithmetic operations.

- (+) adds multiple numbers.
- (-) subtracts one or more numbers.

Syntax:

```
(+ number1 number2 ...)
(- number1 number2 ...)
```

Examples:

```
(+ 1 2 3) ; Returns 6
(- 8 4) ; Returns 4
(- 2) ; Returns -2 (unary operation)
```

### Conditional Expressions

> cond

The cond expression provides a way to execute a specific block of code based on multiple conditions. It checks each condition in sequence and runs the associated code block for the first truthy condition.

Syntax:

```
(cond
(condition1 expression1)
(condition2 expression2)
...)
```

Example:

```
(cond
    (0 (print "0")) ; This will not execute as 0 is falsey
    (0 (print "1")) ; This will not execute as 0 is falsey
    (2 (print "2")) ; This will execute as 2 is truthy
)
```

> if

The if expression evaluates a condition and executes one of two possible blocks based on the result.

Syntax:

```
(if condition
true_expression
false_expression)
```

Example:

```
(if 1
    (print "yes") ; Executes if the condition is truthy
    (print "no") ; Executes if the condition is falsey
)
```

### Variable Assignment

> :=

The := expression assigns a value to a variable. Once assigned, the variable can be used in subsequent expressions.

Syntax:

```
(:= variable_name value)
```

Example:

```
(:= index 0) ; Assigns 0 to index
(print index) ; Prints the value of index
```

### Loop Expressions

> while

The while expression repeats a block of code as long as the condition remains true.

Syntax:

```
(while condition expression)
```

Example:

```
(:= index 0) ; Initialize index to 0
(while (!= index 10) ; Loop until index equals 10
    (+ index 1) ;Increment index by 1
    (print index)
)
```

### Function Definition Expressions

> func

The func expression defines a new function with a specified name and parameters. You can then call the function and optionally return a value using return.

Syntax:

```
(func function_name (param1 param2 ...)
expression
(return return_value))
```

Examples:

```
(func function (alpha beta)
    (print alpha)
    (return  "the_return_value")
)
```

### Function Call Expression

To call a function, use its name followed by its arguments.

Example:

```
(func addition (a b)
    (return (+ a b))
)

(print (addition 100 1)) ; Outputs 101 by calling the addition function
```

## Updates

- Implemented tokenizer (scanner)
- Implemented parser (s-expression parser)
- Added rough runtime interpretation schema
- Added conditional expressions
- Added loop expressions
- Added math expressions
- Added func definition expression
