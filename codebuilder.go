package memorable

import (
	"errors"
	"strconv"
)

// CodeBuilder build the code
type CodeBuilder interface {
	Chars(string) CodeBuilder
	Length(int) CodeBuilder
	Build() (Code, error)
}

type codeBuilder struct {
	characters string
	length     int
	base       int
	Error      error
}

func (c codeBuilder) Chars(chars string) CodeBuilder {
	if len(chars) < 2 || len(chars) > 36 {
		c.Error = errors.New("Number of chars should be in the range 2 to 36")
	}
	c.characters = chars
	c.base = len(chars)
	return c
}

func (c codeBuilder) Length(length int) CodeBuilder {
	if length > 35 || length < 0 {
		c.Error = errors.New("Length should be in the range 0 to 35")
	}
	c.length = length
	return c
}

func (c codeBuilder) Build() (createdCode Code, err error) {
	err = c.Error
	if err != nil {
		return
	}

	baseStr := strconv.FormatUint(uint64(c.base-1), c.base)
	var maxInBase string
	for i := int(0); i < c.length; i++ {
		maxInBase += baseStr
	}

	var maxInDec uint64
	maxInDec, err = strconv.ParseUint(maxInBase, c.base, 64)

	createdCode = &code{
		characters: c.characters,
		length:     c.length,
		max:        maxInDec,
		base:       c.base,
	}

	return
}
