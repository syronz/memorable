package internal

import (
	"fmt"
	"strconv"

	"github.com/syronz/memorable/pkg/format"
)

// Code is an interface for creating memorable code
type Code interface {
	Pong() string
	ShowChars() string
	Return() Code
	Generate() []memCode
}

type code struct {
	characters string
	length     uint8
	max        uint64
	base       int
}

type memCode struct {
	Code string
	Mark uint64
}

func (c code) Pong() string {
	return "ping"
}

func (c code) ShowChars() string {
	return c.characters
}

func (c code) Return() Code {
	return c
}

func (c code) Variation() (mems []memCode) {

	for i := uint64(0); i <= c.max; i++ {
		filledCode, _ := format.AddZeroToLeft(strconv.FormatUint(uint64(i), c.base), int(c.length))
		_ = replaceChars(filledCode, c.characters)
		mem := memCode{
			Code: filledCode,
			Mark: i + 1,
		}
		mems = append(mems, mem)

	}

	return
}

func (c code) Generate() (mems []memCode) {
	mems = c.Variation()

	return
}

func replaceChars(filledCode string, chars string) (replacedCode string) {
	replacedCode = "kkkk"
	charMap := make(map[string]string)
	for i:=0; i < 
	fmt.Println("###########$$$$$$$$$$$ ", chars)
	return
}
