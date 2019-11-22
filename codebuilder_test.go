package memorable

import (
	"errors"
	"testing"
)

func TestBuild(t *testing.T) {
	samples := []struct {
		chars  string
		length int
		err    error
	}{
		{"AB", 2, nil},
		{"AB", 200, errors.New("Length should be in the range 0 to 35")},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789*", 2,
			errors.New("Number of chars should be in the range 2 to 36")},
	}

	for _, v := range samples {
		_, err := New().Chars(v.chars).Length(v.length).Build()

		if err == v.err {
			continue
		}

		if err.Error() != v.err.Error() {
			t.Errorf("For chars = %q and length = %v the error should be %q, but it is %q",
				v.chars, v.length, v.err.Error(), err.Error())

		}

	}
}
