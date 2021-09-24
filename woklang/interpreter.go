package woklang

import (
	"fmt"
	"os"
)

type Interpreter struct {
	root    *Scope
	scope   *Scope
	runtime *Scope
}

func MakeInterpreter() Interpreter {
	interpreter := Interpreter{}
	interpreter.runtime = &Scope{parent: nil, values: RuntimeScope}
	interpreter.root = NewScope(nil)
	interpreter.scope = NewScope(interpreter.root)
	return interpreter
}

func (interpreter *Interpreter) Interpret(statements []Expression) WokData {
	var result WokData = NewWokNull()
	for _, statement := range statements {
		result = interpreter.Evaluate(statement)
	}
	return result
}

func (interpreter *Interpreter) Evaluate(expr Expression) WokData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) Error(errorMessage string) {
	fmt.Println("[Runtime Error] " + errorMessage)
	os.Exit(1)
}

func (interpreter *Interpreter) VisitExpressionList(expr *ExpressionList) WokData {
	if len(expr.value) == 0 {
		return NewWokNull()
	}

	callee := interpreter.Evaluate(expr.value[1])

	if callee.GetType() == WokTypeFunction {
		function := callee.GetValue().(Function)

		return function(interpreter, expr.value[1:])
	}
	return NewWokNull()
}

func (interpreter *Interpreter) VisitExpressionAtom(expr *ExpressionAtom) WokData {
	literal := expr.value.literal

	switch expr.value.ttype {
	case TokenTypeNull:
		return NewWokNull()
	case TokenTypeTrue:
		return NewWokBoolean(true)
	case TokenTypeFalse:
		return NewWokBoolean(false)
	case TokenTypeString:
		return NewWokString(literal)
	case TokenTypeInteger:
		return NewWokInteger(NewWokString(literal).ToInteger())
	case TokenTypeFloat:

		return NewWokFloat(NewWokString(literal).ToFloat())
	case TokenTypeBoolean:
		return NewWokBoolean(NewWokString(literal).ToBoolean())
	case TokenTypeIdentifier:
		scopeValue, ok := interpreter.scope.Get(literal)

		if ok {
			return scopeValue
		}
		interpreter.Error(fmt.Sprintf("Undefined '%s'", literal))
		return NewWokNull()
	case TokenTypeReserved:
		runtimeValue, ok := interpreter.runtime.Get(literal)
		if ok {
			return runtimeValue
		}
		interpreter.Error(fmt.Sprintf("Undefined predicate '%s'", literal))
		return NewWokNull()
	}
	return NewWokNull()
}
