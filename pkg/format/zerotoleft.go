package format

import (
	"errors"
	"fmt"
	"strconv"
)

// AddZeroToLeft fill left side of number with 0
func AddZeroToLeft(num interface{}, length int) (result string, err error) {
	if length > 35 {
		err = errors.New("Maximum length is 35")
		return
	}

	var n uint64

	switch num.(type) {
	case int:
		n = uint64(num.(int))
	case uint64:
		n = num.(uint64)
	case string:
		n, err = strconv.ParseUint(num.(string), 10, 64)
	}

	if err != nil {
		return
	}

	fmt.Printf("%T - %v", n, n)

	result = "OK"
	return
}
