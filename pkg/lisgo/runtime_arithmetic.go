package lisgo

func RuntimeAssignment(interpreter *Interpreter, expressions []Expression) LisgoData {
	value := interpreter.Evaluate(expressions[1])
	token := expressions[0].(*ExpressionAtom).Atom.Literal
	interpreter.Scope.Set(token, value)
	return value
}

func RuntimeAddition(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	count := MathReduce(params, func(total int64, item LisgoData, index int) int64 {
		total += item.ToInteger()
		return total
	}, 0)
	return NewLisgoInteger(count)
}

func RuntimeMultiplication(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	count := MathReduce(params, func(total int64, item LisgoData, index int) int64 {
		total *= item.ToInteger()
		return total
	}, 0)
	return NewLisgoInteger(count)
}

func RuntimeSubstraction(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	if len(params) == 1 {
		return NewLisgoInteger(-params[0].ToInteger())
	}
	count := MathReduce(params[1:], func(total int64, item LisgoData, index int) int64 {
		total -= item.ToInteger()
		return total
	}, params[0].ToInteger())
	return NewLisgoInteger(count)
}
