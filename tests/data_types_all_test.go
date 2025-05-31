package testing

import (
	"testing"

	w "github.com/eugenioenko/lisgo/pkg/lisgo"
)

func TestLisgoInteger(t *testing.T) {
	integer := w.NewLisgoInteger(42)
	if integer.GetType() != w.LisgoTypeInteger {
		t.Errorf("Expected type %d, got %d", w.LisgoTypeInteger, integer.GetType())
	}
	if integer.ToString() != "42" {
		t.Errorf("Expected string '42', got '%s'", integer.ToString())
	}
	if integer.ToBoolean() != true {
		t.Errorf("Expected boolean true, got false")
	}
	if integer.ToInteger() != 42 {
		t.Errorf("Expected integer 42, got %d", integer.ToInteger())
	}
	if integer.ToFloat() != 42.0 {
		t.Errorf("Expected float 42.0, got %f", integer.ToFloat())
	}
}

func TestLisgoString(t *testing.T) {
	str := w.NewLisgoString("hello")
	if str.GetType() != w.LisgoTypeString {
		t.Errorf("Expected type %d, got %d", w.LisgoTypeString, str.GetType())
	}
	if str.ToString() != "hello" {
		t.Errorf("Expected string 'hello', got '%s'", str.ToString())
	}
	if str.ToBoolean() != true {
		t.Errorf("Expected boolean true, got false")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when converting non-numeric string to integer")
		}
	}()
	str.ToInteger()
}

func TestLisgoBoolean(t *testing.T) {
	boolean := w.NewLisgoBoolean(true)
	if boolean.GetType() != w.LisgoTypeBoolean {
		t.Errorf("Expected type %d, got %d", w.LisgoTypeBoolean, boolean.GetType())
	}
	if boolean.ToString() != "true" {
		t.Errorf("Expected string 'true', got '%s'", boolean.ToString())
	}
	if boolean.ToBoolean() != true {
		t.Errorf("Expected boolean true, got false")
	}
	if boolean.ToInteger() != 1 {
		t.Errorf("Expected integer 1, got %d", boolean.ToInteger())
	}
	if boolean.ToFloat() != 1.0 {
		t.Errorf("Expected float 1.0, got %f", boolean.ToFloat())
	}
}

func TestLisgoFloat(t *testing.T) {
	float := w.NewLisgoFloat(3.14)
	if float.GetType() != w.LisgoTypeFloat {
		t.Errorf("Expected type %d, got %d", w.LisgoTypeFloat, float.GetType())
	}
	if float.ToString() != "3.14" {
		t.Errorf("Expected string '3.14', got '%s'", float.ToString())
	}
	if float.ToBoolean() != true {
		t.Errorf("Expected boolean true, got false")
	}
	if float.ToInteger() != 3 {
		t.Errorf("Expected integer 3, got %d", float.ToInteger())
	}
	if float.ToFloat() != 3.14 {
		t.Errorf("Expected float 3.14, got %f", float.ToFloat())
	}
}

func TestLisgoNull(t *testing.T) {
	null := w.NewLisgoNull()
	if null.GetType() != w.LisgoTypeNull {
		t.Errorf("Expected type %d, got %d", w.LisgoTypeNull, null.GetType())
	}
	if null.ToString() != "null" {
		t.Errorf("Expected string 'null', got '%s'", null.ToString())
	}
	if null.ToBoolean() != false {
		t.Errorf("Expected boolean false, got true")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when calling GetValue on null")
		}
	}()
	null.GetValue()
}

func TestLisgoException(t *testing.T) {
	exception := w.NewLisgoException("error")
	if exception.GetType() != w.LisgoTypeException {
		t.Errorf("Expected type %d, got %d", w.LisgoTypeException, exception.GetType())
	}
	if exception.ToString() != "error" {
		t.Errorf("Expected string 'error', got '%s'", exception.ToString())
	}
	if exception.ToBoolean() != true {
		t.Errorf("Expected boolean true, got false")
	}
}

func TestLisgoFunction(t *testing.T) {
	function := w.NewLisgoFunction("testFunc", []string{"arg1"}, nil)
	if function.GetType() != w.LisgoTypeFunction {
		t.Errorf("Expected type %d, got %d", w.LisgoTypeFunction, function.GetType())
	}
	if function.ToString() != "testFunc" {
		t.Errorf("Expected string 'testFunc', got '%s'", function.ToString())
	}
	if function.ToBoolean() != true {
		t.Errorf("Expected boolean true, got false")
	}
}

func TestLisgoCallable(t *testing.T) {
	callable := w.NewLisgoCallable("testCallable", nil)
	if callable.GetType() != w.LisgoTypeCallable {
		t.Errorf("Expected type %d, got %d", w.LisgoTypeCallable, callable.GetType())
	}
	if callable.ToString() != "testCallable" {
		t.Errorf("Expected string 'testCallable', got '%s'", callable.ToString())
	}
	if callable.ToBoolean() != true {
		t.Errorf("Expected boolean true, got false")
	}
}

func TestLisgoReturn(t *testing.T) {
	ret := w.NewLisgoReturn("testReturn", w.NewLisgoInteger(42))
	if ret.GetType() != w.LisgoTypeReturn {
		t.Errorf("Expected type %d, got %d", w.LisgoTypeReturn, ret.GetType())
	}
	if ret.ToString() != "42" {
		t.Errorf("Expected string '42', got '%s'", ret.ToString())
	}
	if ret.ToBoolean() != true {
		t.Errorf("Expected boolean true, got false")
	}
}
