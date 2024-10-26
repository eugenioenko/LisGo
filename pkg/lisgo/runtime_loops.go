package lisgo

func RuntimeWhile(interpreter *Interpreter, expressions []Expression) LisgoData {
	var result LisgoData = NewLisgoNull()
	for interpreter.Evaluate(expressions[0]).ToBoolean() {
		result = interpreter.Evaluate(expressions[1])
	}
	return result
}
