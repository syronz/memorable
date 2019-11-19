package random

type randomNumberBuilder struct {
	min int
	max int
}

func (r *randomNumberBuilder) SetMin(min int) RandomNumberBuilder {
	r.min = min
	return r
}

func (r *randomNumberBuilder) SetMax(max int) RandomNumberBuilder {
	r.max = max
	return r
}

func (r *randomNumberBuilder) Build() RandomNumber {
	return &randomNumber{
		min: r.min,
		max: r.max,
	}
}
