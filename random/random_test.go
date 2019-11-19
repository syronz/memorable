package random

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	for i := 0; i < 20; i++ {
		result := Generate(40)
		// t.Errorf("Result: %v", result)
		t.Log(result)
	}
}
