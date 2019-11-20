package memorable

import (
	"fmt"
	"github.com/syronz/memorable/internal"
	"github.com/syronz/memorable/random"
)

// Ping used for test
func Ping() string {
	return "pong"
}

// Hello is another test
func Hello(name string) string {
	return "Hello " + name
}

// InternalPackageRandom is for testing random numbers, in the test cache is active
func InternalPackageRandom() int {
	return 888
	// rnd := random.Generate(1000)
	// return rnd
}

// RandomInRangeBuilder use Builder design pattern for generating random number
func RandomInRangeBuilder(min, max int) int {
	randEngine := random.New().SetMin(min).SetMax(max).Build()
	return randEngine.Generate()
}

// GetList is for testing builder pattern for internal of application
func GetList() string {
	mem := internal.New().Chars("ABC").Length(15).Build()
	obj := mem.Return()
	fmt.Printf(">>>>>>>> %+v", obj)

	return mem.ShowChars()

}
