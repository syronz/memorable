package memorable

import (
	"errors"
	"strconv"
)

// CodeBuilder build the code
type CodeBuilder interface {
	Chars(string) CodeBuilder
	Length(uint8) CodeBuilder
	Build() (Code, error)
}

type codeBuilder struct {
	characters string
	length     uint8
	base       int
	maxInBase  string
	maxInDec   uint64
	Error      error
}

func (c codeBuilder) Chars(chars string) CodeBuilder {
	c.characters = chars
	c.base = len(chars)
	return c
}

func (c codeBuilder) Length(length uint8) CodeBuilder {
	c.length = length
	return c
}

func (c codeBuilder) Build() (createdCode Code, err error) {
	if c.length > 35 || c.length < 0 {
		err = errors.New("Length should be in this range 0-35")
		return
	}

	if len(c.characters) < 2 || len(c.characters) > 36 {
		err = errors.New("Number of chars should be in this range 2-36")
		return
	}

	baseStr := strconv.FormatUint(uint64(c.base-1), c.base)
	for i := uint8(0); i < c.length; i++ {
		c.maxInBase += baseStr
	}

	c.maxInDec, c.Error = strconv.ParseUint(c.maxInBase, c.base, 64)

	createdCode = &code{
		characters: c.characters,
		length:     c.length,
		max:        c.maxInDec,
		base:       c.base,
	}

	return
}
