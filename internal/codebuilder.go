package internal

// CodeBuilder build the code
type CodeBuilder interface {
	Chars(string) CodeBuilder
	Length(uint8) CodeBuilder
	Build() Code
}

type codeBuilder struct {
	characters string
	length     uint8
}

func (c codeBuilder) Chars(chars string) CodeBuilder {
	c.characters = chars
	return c
}

func (c codeBuilder) Length(length uint8) CodeBuilder {
	c.length = length
	return c
}

func (c codeBuilder) Build() Code {
	return &code{
		characters: c.characters,
		length:     c.length,
	}
}

// New for create an instance of the builder pattern
func New() CodeBuilder {
	return &codeBuilder{}
}
