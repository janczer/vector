package vector

import (
	"fmt"
	"math"
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

func (v Vector) Add(w Vector) Vector {
	if v.dimension != w.dimension {
		panic("Vector should have the same dimension")
	}

	r := []float64{}

	for i := 0; i < v.dimension; i++ {
		r = append(r, v.coordinates[i]+w.coordinates[i])
	}

	return New(r)
}

func (v Vector) Sub(w Vector) Vector {
	if v.dimension != w.dimension {
		panic("Vector should have the same dimension")
	}

	r := []float64{}

	for i := 0; i < v.dimension; i++ {
		r = append(r, v.coordinates[i]-w.coordinates[i])
	}

	return New(r)
}

func (v Vector) Dot(w Vector) float64 {
	if v.dimension != w.dimension {
		panic("Vector should have the same dimension")
	}

	var r float64

	for i := 0; i < v.dimension; i++ {
		r += v.coordinates[i] * w.coordinates[i]
	}

	return r
}

func (v Vector) Scalar(num float64) Vector {
	r := []float64{}

	for i := 0; i < v.dimension; i++ {
		r = append(r, v.coordinates[i]*num)
	}

	return New(r)
}

func (v Vector) Magnitude() float64 {
	var s float64

	for i := 0; i < v.dimension; i++ {
		//fmt.Printf("pow %f = %f\n", v.coordinates[i], math.Pow(v.coordinates[i], 2))
		s += math.Pow(v.coordinates[i], 2.0)
	}

	return math.Sqrt(s)
}

func (v Vector) Normalize() Vector {
	m := v.Magnitude()
	r := []float64{}

	for i := 0; i < v.dimension; i++ {
		r = append(r, v.coordinates[i]/m)
	}

	return New(r)
}

func (v Vector) Angle(w Vector, inDegree bool) float64 {
	d := v.Dot(w)
	mag := v.Magnitude() * w.Magnitude()

	if inDegree {
		return math.Acos(d/mag) * 180.0 / math.Pi
	}

	return math.Acos(d / mag)
}

func (v Vector) Paraller(w Vector) bool {
	a := v.Angle(w, true)
	if a == 0 || a == 180 {
		return true
	}

	return false
}

func (v Vector) Orthogonal(w Vector) bool {
	if v.Angle(w, true) == 90 {
		return true
	}

	return false
}

func (v Vector) Proj(w Vector) Vector {
	return v.Normalize().Scalar(v.Normalize().Dot(w))
}
