package helpers

import (
	"testing"
)

func TestApplyVar(t *testing.T) {
	vars := map[string]interface{}{
		"%TEST%":  "fox",
		"%TEST2%": "dog",
	}

	tS := "quick brown %TEST% jumps over the lazy %TEST2%"

	if ApplyVar(tS, vars) != "quick brown fox jumps over the lazy dog" {
		t.Error("most equal")
	}
}

func TestApplyMap(t *testing.T) {
	vars := map[string]interface{}{
		"%TEST%":  "fox",
		"%TEST2%": "dog",
	}

	tM := map[string]string{"one": "quick brown %TEST% jumps over the lazy %TEST2%"}
	eR := map[string]string{"one": "quick brown fox jumps over the lazy dog"}

	newMap := ApplyMap(tM, vars)

	for k, v := range newMap {
		if eR[k] != v {
			t.Error("error")
		}
	}
}
