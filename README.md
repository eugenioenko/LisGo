# 🦜 LisGo

[![Go Version](https://img.shields.io/badge/Go-1.22-blue)](https://golang.org)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/eugenioenko/lisgo/actions)
[![License](https://img.shields.io/badge/license-MIT-blue)](/home/enko/Documents/lisgo/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/eugenioenko/lisgo)](https://goreportcard.com/report/github.com/eugenioenko/lisgo)

LisGo is a minimalist programming language built with Go. It serves as a foundation to explore and experiment with the complexities of functional programming language interpretation. Initially inspired by Lisp, LisGo will gradually diverge as it evolves.

## 🌐 [Try it out in the Web Assembly Playground!](https://eugenioenko.github.io/LisGo/live)

---

## 📖 Table of Contents

- [Building the Project](#-building-the-project)
- [Building the WebAssembly Binary](#-building-the-webassembly-binary)
- [Running the Project](#-running-the-project)
- [Running the Playground Locally](#-running-the-playground-locally)
- [Command-Line Interface (CLI)](#-command-line-interface-cli)
- [Testing](#-testing)
- [Continuous Integration Pipeline](#-continuous-integration-pipeline)
- [Design Tradeoffs & Notes](#-design-tradeoffs--notes)
- [LisGo Language Documentation](#-lisgo-language-documentation)
- [Updates](#-updates)
- [License](#-license)

---

## 📦 Building the Project

To build the project, simply run:

```bash
go build
```

---

## 🌐 Building the WebAssembly Binary

To build the WebAssembly binary, use:

```bash
make build
```

This will execute:

```bash
GOOS=js GOARCH=wasm go build -o live/lisgo.wasm wasm/lisgo.go
```

---

## ▶️ Running the Project

- **Execute a script**:

  ```bash
  lisgo exec [filename]
  ```

  This will execute the script specified by `[filename]`.

- **Evaluate code directly**:
  ```bash
  lisgo eval [code]
  ```
  This will execute the code passed as an argument.

---

## 🛠️ Running the Playground Locally

WebAssembly's `instantiateStreaming` method requires CORS to be enabled. To run the playground locally, set up a local server:

```bash
npm install http-server
cd live
http-server
```

---

## 🖥️ Command-Line Interface (CLI)

LisGo provides a command-line interface for running LisGo scripts and evaluating code directly from your terminal.

### Usage

```bash
go run lisgo.go
```

Or, if you have built the binary:

```bash
./lisgo
```

### Commands

- **exec [filename]**

  Executes the LisGo script specified by `[filename]`.

  **Example:**

  ```bash
  lisgo exec demo.lisp
  ```

- **eval [code]**

  Evaluates the LisGo code passed as an argument.

  **Example:**

  ```bash
  lisgo eval "(print (+ 1 2 3))"
  ```

- **help**

  Prints the help message with usage instructions.

  **Example:**

  ```bash
  lisgo help
  ```

---

## ✅ Testing

Run all tests using:

```bash
make test
```

This command will execute all the tests located in the `tests` folder by running:

```bash
go test ./...
```

---

## 🚀 Continuous Integration Pipeline

LisGo's CI pipeline is defined in the `.github/workflows/go.yml` file and includes the following jobs:

1. **Lint**:

   - Runs `golangci-lint` to ensure code quality.
   - Checks for linting issues across the codebase.

2. **Format**:

   - Verifies code formatting using `gofmt`.
   - Ensures consistent formatting across all Go files.

3. **Build**:

   - Builds the Go binary for the project.
   - Compiles the WebAssembly binary using:
     ```bash
     GOOS=js GOARCH=wasm go build -o live/lisgo.wasm wasm/lisgo.go
     ```

4. **Test**:

   - Runs all tests in the `tests` folder using:
     ```bash
     go test ./tests/... -v
     ```

5. **Static Analysis**:

   - Executes `go vet` to perform static code analysis and catch potential issues.

6. **Tidy Check**:
   - Ensures the `go.mod` and `go.sum` files are up-to-date.
   - Verifies no uncommitted changes exist after running `go mod tidy`.

The pipeline is triggered on:

- Pushes to the `main` branch.
- Pull requests targeting the `main` branch.

---

## ⚖️ Design Tradeoffs & Notes

LisGo is intentionally minimalist and experimental. Some notable tradeoffs and design decisions include:

- **Simplicity over Completeness:** The language omits many features found in full Lisp implementations to keep the codebase approachable and easy to modify.
- **Performance vs. Clarity:** The interpreter prioritizes code clarity and educational value over raw execution speed or memory efficiency.
- **Error Handling:** Error messages are basic and may not always provide detailed diagnostics, in favor of a simpler implementation.
- **Type System:** LisGo uses dynamic typing with minimal type checking, which can lead to runtime errors but simplifies the interpreter.
- **Standard Library:** The built-in functions are limited; users are encouraged to extend the language as needed.
- **Deviation from Lisp:** While inspired by Lisp, LisGo intentionally diverges in syntax and semantics as new features are added or simplified.
- **Concurrency:** No built-in concurrency primitives are provided, reflecting a focus on core language features first.

These tradeoffs are made to keep LisGo accessible for learning, experimentation, and rapid prototyping.

---

## 📖 LisGo Language Documentation

This section provides an overview of LisGo's syntax and expressions. Each expression is explained with examples.

### 🖨️ Printing Expressions

The `print` expression outputs a value or result to the console.

**Syntax**:

```lisp
(print expression)
```

**Examples**:

```lisp
(print (+ 1 2 3)) ; Outputs the sum of 1, 2, and 3
(print "hello world") ; Outputs "hello world"
(print (print "yes")) ; Prints "yes" and returns nil
```

---

### ➕ Math Expressions

LisGo supports basic arithmetic operations.

- **Addition (`+`)**: Adds multiple numbers.
- **Subtraction (`-`)**: Subtracts one or more numbers.
- **Multiplication (`*`)**: Multiplies multiple numbers.
- **Division (`/`)**: Divides multiple numbers.
- **Modulo (`%`)**: Computes the remainder of division.

**Syntax**:

```lisp
(+ number1 number2 ...)
(- number1 number2 ...)
(* number1 number2 ...)
(/ number1 number2 ...)
(% number1 number2 ...)
```

**Examples**:

```lisp
(+ 1 2 3) ; Returns 6
(- 8 4) ; Returns 4
(* 2 3) ; Returns 6
(/ 8 2) ; Returns 4
(% 10 3) ; Returns 1
```

---

### 🔀 Logical Expressions

LisGo supports logical operations.

- **Equality (`==`)**: Checks if all arguments are equal.
- **Inequality (`!=`)**: Checks if any arguments are not equal.
- **Negation (`!`)**: Negates a boolean value.

**Syntax**:

```lisp
(== value1 value2 ...)
(!= value1 value2 ...)
(! value)
```

**Examples**:

```lisp
(== 1 1 1) ; Returns true
(!= 1 2) ; Returns true
(! true) ; Returns false
```

---

### 🔀 Conditional Expressions

#### `cond`

The `cond` expression executes a specific block of code based on multiple conditions. It checks each condition in sequence and runs the associated code block for the first truthy condition.

**Syntax**:

```lisp
(cond
  (condition1 expression1)
  (condition2 expression2)
  ...)
```

**Example**:

```lisp
(cond
  (0 (print "0")) ; This will not execute as 0 is falsey
  (0 (print "1")) ; This will not execute as 0 is falsey
  (2 (print "2")) ; This will execute as 2 is truthy
)
```

#### `if`

The `if` expression evaluates a condition and executes one of two possible blocks based on the result.

**Syntax**:

```lisp
(if condition
    true_expression
    false_expression)
```

**Example**:

```lisp
(if 1
    (print "yes") ; Executes if the condition is truthy
    (print "no") ; Executes if the condition is falsey
)
```

---

### 📝 Variable Assignment

The `:=` expression assigns a value to a variable. Once assigned, the variable can be used in subsequent expressions.

**Syntax**:

```lisp
(:= variable_name value)
```

**Example**:

```lisp
(:= index 0) ; Assigns 0 to index
(print index) ; Prints the value of index
```

---

### 🔁 Loop Expressions

#### `while`

The `while` expression repeats a block of code as long as the condition remains true.

**Syntax**:

```lisp
(while condition expression)
```

**Example**:

```lisp
(:= index 0) ; Initialize index to 0
(while (!= index 10) ; Loop until index equals 10
    (print
        (:= index (+ index 1))
    )
)
```

---

### 🔧 Function Definition Expressions

#### `func`

The `func` expression defines a new function with a specified name and parameters. You can then call the function and optionally return a value using `return`.

**Syntax**:

```lisp
(func function_name (param1 param2 ...)
    expression
    (return return_value))
```

**Examples**:

```lisp
(func function (alpha beta)
    (print alpha)
    (return "the_return_value")
)
```

---

### 📞 Function Call Expression

To call a function, use its name followed by its arguments.

**Example**:

```lisp
(func addition (a b)
    (return (+ a b))
)

(print (addition 100 1)) ; Outputs 101 by calling the addition function
```

---

### 🔄 Return Expressions

#### `return`

The `return` expression exits a function and optionally returns a value.

**Syntax**:

```lisp
(return value)
```

**Example**:

```lisp
(func example ()
    (return 42)
)
```

#### `return-from`

The `return-from` expression exits a specific function and optionally returns a value.

**Syntax**:

```lisp
(return-from function_name value)
```

**Example**:

```lisp
(func outer ()
    (func inner ()
        (return-from outer 42)
    )
    (inner)
)
```

---

## 🆕 Updates

- ✅ Implemented tokenizer (scanner)
- ✅ Implemented parser (s-expression parser)
- ✅ Added rough runtime interpretation schema
- ✅ Added conditional expressions
- ✅ Added loop expressions
- ✅ Added math expressions
- ✅ Added logical expressions
- ✅ Added function definition expression
- ✅ Added return expressions

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
