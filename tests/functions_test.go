package testing

import (
	"lisgo/pkg/lisgo"
	"testing"
)

func TestDefineAFunction(t *testing.T) {
	source := `
		(func function (a b c)
			(return c)
		)
		(debug function)
	`
	result := lisgo.Eval(source)
	if result.GetType() != lisgo.LisgoTypeFunction {
		t.Fail()
	}
}

func TestFunctionShouldReturnValue(t *testing.T) {
	source := `
		(func function (a b c)
			(return c)
		)
		(debug (function 1 2 777))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 777 {
		t.Fail()
	}
}

func TestFunctionShouldNullUndefinedParams(t *testing.T) {
	source := `
		(func function (a b c)
			(return c)
		)
		(debug (function 1 2))
	`
	result := lisgo.Eval(source)
	if result.GetType() != lisgo.LisgoTypeNull {
		t.Fail()
	}
}

func TestFunctionShouldReturnFromInner(t *testing.T) {
	source := `
		(func function (a b c)
			(func inner (x y z)
				(return-from function 777)
			)
			(inner 1 2 3)
			(return c)
		)
		(debug (function 1 2 3))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 777 {
		t.Fail()
	}
}

func TestFunctionShouldThrow(t *testing.T) {
	source := `
		(func function (a b c)
			(func inner (x y z)
				(return-from function_does_not_exist 777)
			)
			(inner 1 2 3)
			(return c)
		)
		(debug (function 1 2 3))
	`
	result := lisgo.Eval(source)
	if result.GetType() != lisgo.LisgoTypeException {
		t.Fail()
	}
}
