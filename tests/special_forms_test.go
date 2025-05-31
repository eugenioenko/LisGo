package testing

import (
	"testing"

	"github.com/eugenioenko/lisgo/pkg/lisgo"
)

func TestCondWithNoTruthyCondition(t *testing.T) {
	source := `
		(debug (cond
			(false 1)
			(false 2)
		))
	`
	result := lisgo.Eval(source)
	if result.GetType() != lisgo.LisgoTypeNull {
		t.Errorf("Expected null, got %s", result.ToString())
	}
}

func TestIfWithMissingBranches(t *testing.T) {
	source := `
		(debug (if true 1))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 1 {
		t.Errorf("Expected 1, got %d", result.ToInteger())
	}
}

func TestCondWithMixedConditions(t *testing.T) {
	source := `
		(debug (cond
			(false 1)
			(true 2)
			(false 3)
		))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 2 {
		t.Errorf("Expected 2, got %d", result.ToInteger())
	}
}

func TestIfWithInvalidCondition(t *testing.T) {
	source := `
		(debug (if "invalid" 1 2))
	`
	result := lisgo.Eval(source)
	if result.ToInteger() != 1 {
		t.Errorf("Expected 1, got %d", result.ToInteger())
	}
}
