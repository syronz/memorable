package random

type RandomNumber interface {
	Generate() int
}

type RandomNumberBuilder interface {
	SetMin(int) RandomNumberBuilder
	SetMax(int) RandomNumberBuilder
	Build() RandomNumber
}

func New() RandomNumberBuilder {
	return &randomNumberBuilder{}
}
