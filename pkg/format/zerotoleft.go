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

	switch num.(type) {
	case int:
		n = fmt.Sprint(num.(int))
	case uint64:
		n = fmt.Sprint(num.(uint64))
	case string:
		n = num.(string)
	default:
		err = errors.New("Number must be int, uint64 or string")
		return
	}

	str := "00000000000000000000000000000000000"
	str += fmt.Sprint(n)
	result = str[len(str)-length:]

	if err != nil {
		return
	}

	return
}
