package format

import (
	"testing"
)

func TestAddZeroToLeft(t *testing.T) {
	samples := []struct {
		num    interface{}
		length int
		expect string
		err    interface{}
	}{
		{15, 4, "0015", nil},
		{15, 40, "", "Length should be in this range 0-35"},
		{uint8(7), 4, "", "Number must be int, uint64 or string"},
		{nil, 10, "", "Number must be int, uint64 or string"},
		{15, -6, "", "Length should be in this range 0-35"},
		{0, 10, "0000000000", nil},
		{127, 9, "000000127", nil},
		{"1234567890123456789012345678901234567890", 5, "67890", nil},
		{"", 12, "000000000000", nil},
		{uint64(7), 9, "000000007", nil},
	}

	for _, v := range samples {

		r, err := AddZeroToLeft(v.num, v.length)
		if err != nil {
			errMsg := v.err.(string)
			if err.Error() != errMsg {
				t.Errorf(
					"AddZeroToLeft(%v, %v) should return error, returned error is %q which should be %q",
					v.num, v.length, err.Error(), errMsg)
			}
		} else {
			if r != v.expect {
				t.Errorf("AddZeroToLeft(%v, %v) = %v, which is should be %v", v.num, v.length, r, v.expect)
			}
		}
	}
}

func TestConvertorZeroFill(t *testing.T) {
	samples := []struct {
		num    uint64
		base   int
		length int
		result string
	}{
		{3, 2, 5, "00011"},
	}

	for _, v := range samples {
		r := ConvertorZeroFill(v.num, v.base, v.length)
		if r != v.result {
			t.Errorf("ConvertorZeroFill(%v, %v, %v) = %q, which is should be %q",
				v.num, v.base, v.length, r, v.result)
		}
	}

}
