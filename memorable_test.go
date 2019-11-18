package memorable

import (
	"testing"
)

func TestPing(t *testing.T) {
	sample := struct {
		out string
	}{
		out: "pong",
	}

	r := Ping()
	if r != sample.out {
		t.Errorf("Output is not equal with Pong!")
	}
	t.Log("good job")
}
