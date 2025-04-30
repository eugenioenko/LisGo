package lisgo

import (
	"fmt"
)

type Interpreter struct {
	Root    *Scope
	Scope   *Scope
	Runtime *Scope
}

func MakeInterpreter() Interpreter {
	interpreter := Interpreter{}
	interpreter.Runtime = &Scope{parent: nil, values: RuntimeScope}
	interpreter.Root = NewScope(nil)
	interpreter.Scope = NewScope(interpreter.Root)
	return interpreter
}

func (interpreter *Interpreter) Interpret(statements []Expression) (result LisgoData) {
	for _, statement := range statements {
		result = interpreter.Evaluate(statement)
	}
	return result
}

func (interpreter *Interpreter) Evaluate(expr Expression) LisgoData {
	return expr.Accept(interpreter)
}

func (interpreter *Interpreter) Error(errorMessage string) {
	panic("[Runtime Error] " + errorMessage)
}

func (interpreter *Interpreter) FunctionCall(function *LisgoFunction, expressions []Expression) (result LisgoData) {
	params := EvalParams(interpreter, expressions)
	paramsMaxIndex := len(params) - 1
	scope := interpreter.Scope
	interpreter.Scope = NewScope(scope)
	for index := 0; index < len(function.args); index++ {
		if index <= paramsMaxIndex {
			interpreter.Scope.Set(function.args[index], params[index])
		} else {
			interpreter.Scope.Set(function.args[index], NewLisgoNull())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			ret := err.(*LisgoReturn)
			if ret.From == function.name || ret.From == "" {
				result = err.(*LisgoReturn).Value
				interpreter.Scope = scope
			} else {
				panic(err)
			}
		}
	}()
	result = interpreter.Interpret(function.body)
	interpreter.Scope = scope
	return result
}

func (interpreter *Interpreter) VisitExpressionList(expr *ExpressionList) LisgoData {
	if len(expr.List) == 0 {
		return NewLisgoNull()
	}
	callee := interpreter.Evaluate(expr.List[0])
	if callee.GetType() == LisgoTypeCallable {
		function := callee.GetValue().(Callable)
		return function(interpreter, expr.List[1:])
	}
	if callee.GetType() == LisgoTypeFunction {
		return interpreter.FunctionCall(callee.(*LisgoFunction), expr.List[1:])
	}
	return NewLisgoNull()
}

func (interpreter *Interpreter) VisitExpressionAtom(expr *ExpressionAtom) LisgoData {
	literal := expr.Atom.Literal

	switch expr.Atom.Type {
	case TokenTypeNull:
		return NewLisgoNull()
	case TokenTypeTrue:
		return NewLisgoBoolean(true)
	case TokenTypeFalse:
		return NewLisgoBoolean(false)
	case TokenTypeString:
		return NewLisgoString(literal)
	case TokenTypeInteger:
		return NewLisgoInteger(NewLisgoString(literal).ToInteger())
	case TokenTypeFloat:
		return NewLisgoFloat(NewLisgoString(literal).ToFloat())
	case TokenTypeBoolean:
		return NewLisgoBoolean(NewLisgoString(literal).ToBoolean())
	case TokenTypeIdentifier:
		scopeValue, ok := interpreter.Scope.Get(literal)
		if ok {
			return scopeValue
		}
		interpreter.Error(fmt.Sprintf("Undefined '%s'", literal))
		return NewLisgoNull()
	case TokenTypeReserved:
		runtimeValue, ok := interpreter.Runtime.Get(literal)
		if ok {
			return runtimeValue
		}
		interpreter.Error(fmt.Sprintf("Undefined predicate '%s'", literal))
		return NewLisgoNull()
	}
	return NewLisgoNull()
}
