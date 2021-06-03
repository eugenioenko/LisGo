package main

func main() {

	tokenizer := MakeTokenizer()
	tokenizer.LoadFromFile("demo.wok")
	tokenizer.Tokenize()

	parser := MakeParser()
	parser.Parse(tokenizer.tokens)

	interpreter := MakeInterpreter()
	interpreter.Interpret(parser.statements)

}
