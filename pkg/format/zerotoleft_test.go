package format

import (
	"errors"
	"fmt"
	"testing"
)

func TestAddZeroToLeft(t *testing.T) {
	samples := []struct {
		num    interface{}
		length int
		expect string
		err    error
	}{
		{15, 4, "0015", nil},
		{15, 40, "", nil},
	}

	for _, v := range samples {

		r, err := AddZeroToLeft(v.num, v.length)
		fmt.Printf(">>>%T ++++ %+v\n", err, err)
		if err != nil {
			fmt.Printf("######## %T ///// %[1]v\n", err.Error() == "Maximum length is 35")
		}
		fmt.Println("~~~~~~~~~~~~~", err == errors.New("Maximum length is 35"))
		t.Log(r, err)
	}
}
