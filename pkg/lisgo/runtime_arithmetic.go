package lisgo

import (
	"math"
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
	return RuntimeBinary(params, "+")
}

func RuntimeSubtraction(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	return RuntimeBinary(params, "-")
}

func RuntimeMultiplication(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	return RuntimeBinary(params, "*")
}

func RuntimeDivision(interpreter *Interpreter, expressions []Expression) LisgoData {
	params := EvalParams(interpreter, expressions)
	return RuntimeBinary(params, "/")
}

func RuntimeBinary(params []LisgoData, operation string) LisgoData {
	hasStringType := ParamsSomeAreType(params, LisgoTypeString)
	if operation == "+" && hasStringType {
		return NewLisgoString(ConcatenateStringParams(params))
	}

	hasFloatType := ParamsSomeAreType(params, LisgoTypeFloat)
	if hasFloatType || operation == "/" {
		result := Reduce(params[1:], func(total float64, item LisgoData, _ int) float64 {
			value := item.ToFloat()
			switch operation {
			case "+":
				return total + value
			case "-":
				return total - value
			case "*":
				return total * value
			case "/":
				return total / value
			case "%":
				return math.Mod(total, value)
			}
			panic("Unknown binary operation " + operation)
		}, params[0].ToFloat())
		// cast result to integer when result has no decimals
		if result == float64(int64(result)) {
			return NewLisgoInteger(int64(result))
		} else {
			return NewLisgoFloat(result)
		}
	}

	result := Reduce(params[1:], func(total int64, item LisgoData, _ int) int64 {
		value := item.ToInteger()
		switch operation {
		case "+":
			return total + value
		case "-":
			return total - value
		case "*":
			return total * value
		case "%":
			return total % value
		}
		panic("Unknown binary operation " + operation)
	}, params[0].ToInteger())
	return NewLisgoInteger(result)
}

func ConcatenateStringParams(items []LisgoData) string {
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
