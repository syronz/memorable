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

func TestHello(t *testing.T) {
	samples := []struct {
		in  string
		out string
	}{
		{"world", "Hello world"},
		{"Diako", "Hello Diako"},
		{"Ako", "Hello Ako"},
		{"", "Hello "},
	}

	for _, v := range samples {
		r := Hello(v.in)
		if r != v.out {
			t.Errorf("Hello(%v) should return %q, but returns %q", v.in, v.out, r)
		}
	}
}

func TestInternalPackageRandom(t *testing.T) {
	result := InternalPackageRandom()
	t.Log("Result is: ", result)
}

func TestGetList(t *testing.T) {
	result := GetList()
	t.Log("List is: ", result)
}
