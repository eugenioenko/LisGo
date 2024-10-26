package testing

import (
	"fmt"
	w "lisgo/pkg/lisgo"
	"math/rand"
	"testing"
)

func TestShouldBeInteger(t *testing.T) {
	for i := 0; i < 777; i++ {
		rnd := rand.Int63()
		v := w.Eval(fmt.Sprintf("(debug %d)", rnd))
		if v.GetType() != w.LisgoTypeInteger {
			t.Fail()
		}
	}
}

func TestShouldBeSameInteger(t *testing.T) {
	for i := 0; i < 777; i++ {
		rnd := rand.Int63()
		v := w.Eval(fmt.Sprintf("(debug %d)", rnd))
		if v.ToInteger() != rnd {
			t.Fail()
		}
	}
}

func TestShouldBeString(t *testing.T) {
	for i := 0; i < 777; i++ {
		rnd := rand.Int63()
		v := w.Eval(fmt.Sprintf(`(debug "%d")`, rnd))
		if v.GetType() != w.LisgoTypeString {
			t.Fail()
		}
	}
}
func TestShouldBeBooleanTrue(t *testing.T) {
	v := w.Eval("(debug true)")
	if v.GetType() != w.LisgoTypeBoolean {
		fmt.Print(v.ToString())
		t.Fail()
	}
	if v.ToBoolean() != true {
		t.Fail()
	}
}

func TestShouldBeBooleanFalse(t *testing.T) {
	v := w.Eval("(debug false)")
	if v.GetType() != w.LisgoTypeBoolean {
		t.Fail()
	}
	if v.ToBoolean() != false {
		t.Fail()
	}
}
