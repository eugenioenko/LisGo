package testing

import (
	. "lisgo/pkg/lisgo"
	"testing"
)

func TestAdditionInt1(t *testing.T) {
	source := `
		(:= value (+ 1 2 3))
		(debug value)
	`
	result := Eval(source)
	if result.GetType() != LisgoTypeInteger || result.ToInteger() != 6 {
		t.Fail()
	}
}

func TestAdditionConcat1(t *testing.T) {
	source := `
		(:= value (+ "1" "2" 3 4.4))
		(debug value)
	`
	result := Eval(source)
	if result.GetType() != LisgoTypeString || result.ToString() != "1234.4" {
		t.Fail()
	}
}

func TestDivisionInt(t *testing.T) {
	source := `
		(:= value (/ 100 20 2))
		(debug value)
	`
	result := Eval(source)
	if result.GetType() != LisgoTypeFloat || result.ToFloat() != 2.5 {
		t.Fail()
	}
}

func TestDivisionIntInt(t *testing.T) {
	source := `
		(:= value (/ 100 2))
		(debug value)
	`
	result := Eval(source)
	if result.GetType() != LisgoTypeInteger || result.ToInteger() != 50 {
		t.Fail()
	}
}
