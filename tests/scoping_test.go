package testing

import (
	"lisgo/pkg/lisgo"
	"testing"
)

func TestVariableShadowingAcrossScopes(t *testing.T) {
	source := `
		(:= x 10)
		(func test ()
			(:= x 20)
			(debug x)
		)
		(debug (test))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 20 {
		t.Errorf("Expected 20, got %d", result.ToInteger())
	}
}

func TestAccessParentScopeVariable(t *testing.T) {
	source := `
		(:= x 10)
		(func test ()
			(debug x)
		)
		(debug (test))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 10 {
		t.Errorf("Expected 10, got %d", result.ToInteger())
	}
}
