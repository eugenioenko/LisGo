package lisgo

func WF(name string, function Callable) *LisgoCallable {
	return &LisgoCallable{DType: LisgoTypeCallable, name: name, function: function}
}

func EvalParams(interpreter *Interpreter, expressions []Expression) []LisgoData {
	params := make([]LisgoData, len(expressions))
	for index, expression := range expressions {
		params[index] = interpreter.Evaluate(expression)
	}
	return params
}

func RuntimeDebug(interpreter *Interpreter, expressions []Expression) LisgoData {
	return interpreter.Evaluate(expressions[0])
}
