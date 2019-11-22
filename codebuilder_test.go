package memorable

import (
	"testing"
)

func TestBuild(t *testing.T) {

	// mem, err := New().Chars("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789").Length(2).Build()
	mem, err := New().Chars("AB").Length(2).Build()
	if err != nil {
		t.Error(err.Error())
	}

	variations := mem.Variation()

	for _, v := range variations {
		t.Logf("%4d - %v\n", v.Mark, v.Code)
	}

}
