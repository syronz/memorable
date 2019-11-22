package memorable

import (
	"github.com/syronz/memorable/pkg/format"
	"strconv"
)

// Code is an interface for creating memorable code
type Code interface {
	Generate() []memCode
	Variation() []memCode
}

type code struct {
	characters string
	length     int
	max        uint64
	base       int
}

type memCode struct {
	Code string
	Mark uint64
}

func (c code) Variation() (mems []memCode) {
	swapper := c.replaceChars()
	mems = make([]memCode, c.max+1)

	for i := uint64(0); i <= c.max; i++ {
		filledCode, _ := format.AddZeroToLeft(strconv.FormatUint(uint64(i), c.base), int(c.length))
		swappedCode := swapper(filledCode)

		mem := memCode{
			Code: swappedCode,
			Mark: i + 1,
		}
		mems[i] = mem
	}

	return
}

func (c code) Generate() (mems []memCode) {
	mems = c.Variation()

	return
}

func (c code) replaceChars() func(string) string {
	charMap := make(map[string]rune)
	for i := 0; i < c.base; i++ {
		charBase := strconv.FormatUint(uint64(i), c.base)
		charMap[charBase] = rune(c.characters[i])
	}

	return func(filledCode string) string {
		str := []rune(filledCode)

		for i, v := range filledCode {
			str[i] = charMap[string(v)]
		}

		return string(str)
	}
}
