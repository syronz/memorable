package random

type randomNumberBuilder struct {
	min int
	max int
}

func (r *randomNumberBuilder) SetMin(min int) NumberBuilder {
	r.min = min
	return r
}

func (r *randomNumberBuilder) SetMax(max int) NumberBuilder {
	r.max = max
	return r
}

func (r *randomNumberBuilder) Build() Number {
	return &randomNumber{
		min: r.min,
		max: r.max,
	}
}
