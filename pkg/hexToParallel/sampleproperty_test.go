package hexToParallel

import (
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"math"
	"testing"
)

func TestHexToParallelOneByte(t *testing.T) {

}

func TestHexParse(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("hex to int", prop.ForAll(
		HexToParallelSanitizeWord,
		//gen.AnyString(),
	))

	properties.TestingRun(t)
}

func TestSqrt(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("greater one of all greater one", prop.ForAll(
		func(v float64) bool {
			return math.Sqrt(v) >= 1
		},
		gen.Float64Range(1, math.MaxFloat64),
	))

	properties.Property("squared is equal to value", prop.ForAll(
		func(v float64) bool {
			r := math.Sqrt(v)
			return math.Abs(r*r-v) < 1e-10*v
		},
		gen.Float64Range(0, math.MaxFloat64),
	))

	properties.TestingRun(t)
}
