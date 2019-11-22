package random

// Number is an interface for generating new number
type Number interface {
	Generate() int
}

// NumberBuilder is used for building RandomNumber
type NumberBuilder interface {
	SetMin(int) NumberBuilder
	SetMax(int) NumberBuilder
	Build() Number
}

// New return back an instance of RandomNumberBuilder
func New() NumberBuilder {
	return &randomNumberBuilder{}
}
