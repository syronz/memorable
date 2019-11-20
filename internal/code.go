package internal

// Code is an interface for creating memorable code
type Code interface {
	Pong() string
	ShowChars() string
	Return() Code
}

type code struct {
	characters string
	length     uint8
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
