package lisgo

func RuntimeCond(interpreter *Interpreter, expressions []Expression) LisgoData {
	for _, expression := range expressions {
		condition := expression.(*ExpressionList).List
		if interpreter.Evaluate(condition[0]).ToBoolean() {
			return interpreter.Evaluate(condition[1])
		}
	}
	return NewLisgoNull()
}

func RuntimeIf(interpreter *Interpreter, expressions []Expression) LisgoData {
	expressionsCount := len(expressions)
	if expressionsCount > 0 {
		if interpreter.Evaluate(expressions[0]).ToBoolean() {
			if expressionsCount > 1 {
				return interpreter.Evaluate(expressions[1])
			}
		} else {
			if expressionsCount > 2 {
				return interpreter.Evaluate(expressions[2])
			}
		}
	}
	return NewLisgoNull()
}
