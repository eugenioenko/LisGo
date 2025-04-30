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
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected parser to panic")
		}
	}()
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
	lisgo.Eval(source)
}

func TestRecursiveFunction(t *testing.T) {
	source := `
		(func factorial (n)
			(if (== n 0)
				(return 1)
				(return (* n (factorial (- n 1))))
			)
		)
		(debug (factorial 5))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 120 {
		t.Fail()
	}
}

func TestInvalidFunctionCall(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected parser to panic on undefined function")
		}
	}()
	source := `
		(debug (undefined_function 1 2 3))
	`
	lisgo.Eval(source)
}

func TestVariableShadowing(t *testing.T) {
	source := `
		(:= x 10)
		(func test (x)
			(return x)
		)
		(debug (test 20))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 20 {
		t.Fail()
	}
}

func TestInvalidFunctionDefinition(t *testing.T) {
	source := `
		(func invalid_function)
	`
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid function definition")
		}
	}()
	lisgo.Eval(source)
}
