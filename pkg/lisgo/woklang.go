package lisgo

import (
	"fmt"
)

func Eval(source string) (result LisgoData) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[Runtime Error] Oops! Unhandled Error")
			result = NewLisgoException("Unhandled Exception")
		}
	}()
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromString(source)
	tokens := tokenizer.Tokenize()

	parser := MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := MakeInterpreter()
	result = interpreter.Interpret(expressions)
	return result
}

func Exec(filename string) (result LisgoData) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[Runtime Error] Oops! Unhandled Error")
			result = NewLisgoException("Unhandled Exception")
		}
	}()
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile(filename)
	tokens := tokenizer.Tokenize()

	parser := MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := MakeInterpreter()
	return interpreter.Interpret(expressions)
}
