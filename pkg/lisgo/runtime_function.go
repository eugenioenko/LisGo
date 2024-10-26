package lisgo

func RuntimeFunc(interpreter *Interpreter, expressions []Expression) LisgoData {
	name := expressions[0].(*ExpressionAtom).Atom.Literal
	args := make([]string, len(expressions[1].(*ExpressionList).List))
	for index, token := range expressions[1].(*ExpressionList).List {
		args[index] = token.(*ExpressionAtom).Atom.Literal
	}
	function := NewLisgoFunction(name, args, expressions[2:])
	interpreter.Scope.Set(name, function)
	return function
}

func RuntimeReturnFrom(interpreter *Interpreter, expressions []Expression) LisgoData {
	result := interpreter.Evaluate(expressions[1])
	from := expressions[0].(*ExpressionAtom).Atom.Literal

	panic(NewLisgoReturn(from, result))
}

func RuntimeReturn(interpreter *Interpreter, expressions []Expression) LisgoData {
	result := interpreter.Evaluate(expressions[0])

	panic(NewLisgoReturn("", result))
}
