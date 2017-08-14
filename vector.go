package vector

import (
	"fmt"
	//"strconv"
)

type Vector struct {
	coordinates []float64
	dimension   int
}

func New(c []float64) Vector {
	return Vector{c, len(c)}
}

func (v Vector) Print() string {
	s := ""
	for i := 0; i < v.dimension-1; i++ {
		s += fmt.Sprintf("%.2f, ", v.coordinates[i])
	}

	return "Vector: " + s + fmt.Sprintf("%.2f", v.coordinates[v.dimension-1])
}

func (v Vector) Eq(w Vector) bool {
	if v.dimension != w.dimension {
		panic("Vector should have the same dimension")
	}

	for i := 0; i < v.dimension; i++ {
		if v.coordinates[i] != w.coordinates[i] {
			return false
		}
	}

	return true
}
