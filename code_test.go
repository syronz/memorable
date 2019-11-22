package memorable

import (
	"testing"
)

func TestVariation(t *testing.T) {
	samples := []struct {
		chars        string
		length       int
		expectLength int
	}{
		{"AB", 2, 4},
		{"ABC", 2, 9},
	}

	for _, v := range samples {
		mem, _ := New().Chars(v.chars).Length(v.length).Build()
		variation := mem.Variation()
		if len(variation) != v.expectLength {
			t.Errorf("For chars = %q and length = %v the expected length should be %d, but it is %d",
				v.chars, v.length, v.expectLength, len(variation))
		}
	}

}

func TestGenerate(t *testing.T) {
	samples := []struct {
		chars        string
		length       int
		expectLength int
	}{
		{"AB", 2, 4},
		{"ABC", 2, 9},
	}

	for _, v := range samples {
		mem, _ := New().Chars(v.chars).Length(v.length).Build()
		codes := mem.Generate()
		if len(codes) != v.expectLength {
			t.Errorf("For chars = %q and length = %v the expected length should be %d, but it is %d",
				v.chars, v.length, v.expectLength, len(codes))
		}
	}

}
