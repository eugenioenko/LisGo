package testing

import (
	"lisgo/pkg/lisgo"
	"testing"
)

func TestDeeplyNestedExpressions(t *testing.T) {
	source := "(debug " + repeat("(+ ", 1000) + "1" + repeat(")", 1000) + ")"
	result := lisgo.Eval(source)
	if result.ToInteger() != 1 {
		t.Errorf("Expected 1, got %d", result.ToInteger())
	}
}

func TestLargeNumberHandling(t *testing.T) {
	source := "(debug 9223372036854775807)" // Max int64
	result := lisgo.Eval(source)
	if result.ToInteger() != 9223372036854775807 {
		t.Errorf("Expected 9223372036854775807, got %d", result.ToInteger())
	}
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
