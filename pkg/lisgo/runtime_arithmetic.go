package lisgo

import (
	"strings"
)

func RuntimeAssignment(interpreter *Interpreter, expressions []Expression) LisgoData {
	value := interpreter.Evaluate(expressions[1])
	token := expressions[0].(*ExpressionAtom).Atom.Literal
	interpreter.Scope.Set(token, value)
	return value
}

func RuntimeAddition(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)

	hasStringParam := ParamsHaveLisgoType(params, LisgoTypeString)
	if hasStringParam {
		return NewLisgoString(Concatenate(params))
	}

	hasFloatType := ParamsHaveLisgoType(params, LisgoTypeFloat)
	if hasFloatType {
		result := Reduce(params, func(total float64, item LisgoData, _ int) float64 {
			return total + item.ToFloat()
		}, 0.0)
		return NewLisgoFloat(result)
	}

	result := Reduce(params, func(total int64, item LisgoData, _ int) int64 {
		return total + item.ToInteger()
	}, 0)
	return NewLisgoInteger(result)
}

func RuntimeSubtraction(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	hasFloatType := ParamsHaveLisgoType(params, LisgoTypeFloat)
	hasIntType := ParamsHaveLisgoType(params, LisgoTypeInteger)
	paramsCount := len(params)
	switch {
	case hasFloatType && paramsCount == 1:
		return NewLisgoFloat(-params[0].ToFloat())
	case hasFloatType && paramsCount > 1:
		return NewLisgoFloat(params[0].ToFloat() - params[1].ToFloat())
	case hasIntType && paramsCount == 1:
		return NewLisgoInteger(-params[0].ToInteger())
	case hasIntType && paramsCount > 1:
		return NewLisgoInteger(params[0].ToInteger() - params[1].ToInteger())
	}
	return NewLisgoNull()
}

func RuntimeMultiplication(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	hasFloatType := ParamsHaveLisgoType(params, LisgoTypeFloat)
	if hasFloatType {
		result := Reduce(params, func(total float64, item LisgoData, _ int) float64 {
			return total * item.ToFloat()
		}, 1.0)
		return NewLisgoFloat(result)
	}

	result := Reduce(params, func(total int64, item LisgoData, _ int) int64 {
		return total * item.ToInteger()
	}, 1)
	return NewLisgoInteger(result)
}

func RuntimeDivision(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	return NewLisgoFloat(params[0].ToFloat() / params[1].ToFloat())
}

func Concatenate(items []LisgoData) string {
	strs := Map(items, func(item LisgoData) string {
		return item.ToString()
	})

	totalLen := 0
	for _, s := range strs {
		totalLen += len(s)
	}

	var builder strings.Builder
	builder.Grow(totalLen)

	for _, s := range strs {
		builder.WriteString(s)
	}

	return builder.String()
}
