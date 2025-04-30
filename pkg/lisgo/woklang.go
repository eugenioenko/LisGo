package lisgo

func Eval(source string) (result LisgoData) {
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
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile(filename)
	tokens := tokenizer.Tokenize()

	parser := MakeParser()
	expressions := parser.Parse(tokens)

	interpreter := MakeInterpreter()
	return interpreter.Interpret(expressions)
}
