package testing

import (
	"testing"

	"github.com/eugenioenko/lisgo/pkg/lisgo"
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

func TestVariableShadowingInNestedFunctions(t *testing.T) {
	source := `
		(:= x 10)
		(func outer ()
			(:= x 20)
			(func inner ()
				(debug x)
			)
			(debug (inner))
		)
		(debug (outer))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 20 {
		t.Errorf("Expected 20, got %d", result.ToInteger())
	}
}

func TestAccessUndefinedVariable(t *testing.T) {
	source := `(debug y)`
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for undefined variable")
		}
	}()
	lisgo.Eval(source)
}
