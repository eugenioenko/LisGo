# 🦜 LisGo

LisGo is a minimalist programming language built with Go. It serves as a foundation to explore and experiment with the complexities of functional programming language interpretation. Initially inspired by Lisp, LisGo will gradually diverge as it evolves.

## 🌐 [Try it out in the Web Assembly Playground!](https://eugenioenko.github.io/LisGo/live)

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
