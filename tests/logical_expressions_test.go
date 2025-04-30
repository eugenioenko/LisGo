package testing

import (
	w "lisgo/pkg/lisgo"
	"testing"
)

func TestInequalityShouldBeBooleanFalse(t *testing.T) {
	v := w.Eval("(debug (!= 1 1))")
	if v.GetType() != w.LisgoTypeBoolean {
		t.Fail()
	}
	if v.ToBoolean() != false {
		t.Fail()
	}
}

func TestInequalityShouldBeBooleanTrue(t *testing.T) {
	v := w.Eval("(debug (!= 1 2))")
	if v.GetType() != w.LisgoTypeBoolean {
		t.Fail()
	}
	if v.ToBoolean() != true {
		t.Fail()
	}
}

func TestEqualityShouldBeBooleanTrue(t *testing.T) {
	v := w.Eval("(debug (== 1 1))")
	if v.GetType() != w.LisgoTypeBoolean {
		t.Fail()
	}
	if v.ToBoolean() != true {
		t.Fail()
	}
}

func TestEqualityMultipleShouldBeBooleanTrue(t *testing.T) {
	v := w.Eval("(debug (== 1 1 1 1 1 1 1 1))")
	if v.GetType() != w.LisgoTypeBoolean {
		t.Fail()
	}
	if v.ToBoolean() != true {
		t.Fail()
	}
}

func TestEqualityShouldBeBooleanFalse(t *testing.T) {
	v := w.Eval("(debug (== 1 2))")
	if v.GetType() != w.LisgoTypeBoolean {
		t.Fail()
	}
	if v.ToBoolean() != false {
		t.Fail()
	}
}

func TestEqualityMultipleShouldBeBooleanFalse(t *testing.T) {
	v := w.Eval("(debug (== 1 2 3 4 5 6))")
	if v.GetType() != w.LisgoTypeBoolean {
		t.Fail()
	}
	if v.ToBoolean() != false {
		t.Fail()
	}
}

func TestInvalidLogicalExpression(t *testing.T) {
	v := w.Eval("(debug (!= 1 \"string\"))")
	if v.GetType() != w.LisgoTypeBoolean {
		t.Fail()
	}
	if v.ToBoolean() != true {
		t.Fail()
	}
}

func TestMixedTypeEquality(t *testing.T) {
	v := w.Eval("(debug (== 1 \"1\"))")
	if v.GetType() != w.LisgoTypeBoolean || v.ToBoolean() != false {
		t.Fail()
	}
}
