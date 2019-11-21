package internal

import (
	"fmt"
	"strconv"
)

// CodeBuilder build the code
type CodeBuilder interface {
	Chars(string) CodeBuilder
	Length(uint8) CodeBuilder
	Build() Code
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

func (c codeBuilder) Build() Code {

	baseStr := strconv.FormatUint(uint64(c.base-1), c.base)
	for i := uint8(0); i < c.length; i++ {
		c.maxInBase += baseStr
	}
	fmt.Println("!!!!!!!!!!!!!!!!!!>>>!!", baseStr, c.maxInBase)

	c.maxInDec, c.Error = strconv.ParseUint(c.maxInBase, c.base, 64)

	return &code{
		characters: c.characters,
		length:     c.length,
		max:        c.maxInDec,
		base:       c.base,
	}
}

// New for create an instance of the builder pattern
func New() CodeBuilder {
	return &codeBuilder{}
}
