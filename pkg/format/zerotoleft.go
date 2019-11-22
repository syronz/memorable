package format

import (
	"errors"
	"fmt"
)

// AddZeroToLeft fill left side of number with 0
func AddZeroToLeft(num interface{}, length int) (result string, err error) {
	if length > 35 || length < 0 {
		err = errors.New("Length should be in this range 0-35")
		return
	}

	var n string

	switch v := num.(type) {
	case int, uint64, string:
		n = fmt.Sprint(v)
	default:
		err = errors.New("Number must be int, uint64 or string")
		return
	}

	str := "00000000000000000000000000000000000"
	str += fmt.Sprint(n)
	result = str[len(str)-length:]

	return
}
