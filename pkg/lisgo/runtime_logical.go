package lisgo

func RuntimeEquality(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	result := Every(params, func(item LisgoData, index int) bool {
		return item.GetType() == params[0].GetType() && item.GetValue() == params[0].GetValue()
	})
	return NewLisgoBoolean(result)
}

func RuntimeInequality(interpreter *Interpreter, expressions []Expression) LisgoData {
	result := RuntimeEquality(interpreter, expressions)
	return NewLisgoBoolean(!result.ToBoolean())
}

func RuntimeNegation(interpreter *Interpreter, expressions []Expression) LisgoData {
	result := interpreter.Evaluate(expressions[0])
	return NewLisgoBoolean(!result.ToBoolean())
}
