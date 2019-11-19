package random

type randomNumber struct {
	min int
	max int
}

func (r *randomNumber) Generate() int {
	return GenerateMinMax(r.min, r.max)
}
