package random

import (
	"testing"
)

func TestRandomNumberBuilder(t *testing.T) {
	t.Log("INSIDE THE TEST")

	randEngine := New().SetMin(10).SetMax(100).Build()
	randNum := randEngine.Generate()

	t.Log(">>>>> ", randNum)

}
