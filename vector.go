package vector

import (
	"fmt"
	//"math"
	"math/big"
	//"strconv"
)

const prec = 50

type Vector struct {
	coordinates []*big.Float
	dimension   int
}

func New(c []*big.Float) Vector {
	for i := 0; i < len(c); i++ {
		c[i].SetPrec(prec)
	}
	return Vector{c, len(c)}
}

func (v Vector) Print() string {
	s := ""
	for i := 0; i < v.dimension-1; i++ {
		s += fmt.Sprintf("%.5f, ", v.coordinates[i])
	}

	return "Vector: " + s + fmt.Sprintf("%.5f", v.coordinates[v.dimension-1])
}

func (v Vector) Eq(w Vector) bool {
	if v.dimension != w.dimension {
		panic("Vector should have the same dimension")
	}

	for i := 0; i < v.dimension; i++ {
		if v.coordinates[i].Cmp(w.coordinates[i]) != 0 {
			return false
		}
	}

	return true
}

func (v Vector) Add(w Vector) Vector {
	if v.dimension != w.dimension {
		panic("Vector should have the same dimension")
	}

	r := []*big.Float{}

	for i := 0; i < v.dimension; i++ {
		z := big.NewFloat(0)
		z.Add(v.coordinates[i], w.coordinates[i])
		r = append(r, z)
	}

	return New(r)
}

func (v Vector) Sub(w Vector) Vector {
	if v.dimension != w.dimension {
		panic("Vector should have the same dimension")
	}

	r := []*big.Float{}

	for i := 0; i < v.dimension; i++ {
		r = append(r, big.NewFloat(0).Sub(v.coordinates[i], w.coordinates[i]))
	}

	return New(r)
}

func (v Vector) Dot(w Vector) *big.Float {
	if v.dimension != w.dimension {
		panic("Vector should have the same dimension")
	}

	r := big.NewFloat(0).SetPrec(prec)

	for i := 0; i < v.dimension; i++ {
		z := big.NewFloat(0).SetPrec(prec).Mul(v.coordinates[i], w.coordinates[i])
		r = z.Add(z, r)
	}

	return r
}

func (v Vector) Scalar(num float64) Vector {
	r := []*big.Float{}

	for i := 0; i < v.dimension; i++ {
		r = append(r, big.NewFloat(0).Mul(v.coordinates[i], big.NewFloat(num)))
	}

	return New(r)
}

func (v Vector) Magnitude() float64 {
	var s *big.Float

	for i := 0; i < v.dimension; i++ {
		z := big.NewFloat(0).Mul(v.coordinates[i], v.coordinates[i])
		s = z.Add(z, s)
	}

	//return math.Sqrt(s)
	return 0.0
}

//func (v Vector) Normalize() Vector {
//m := v.Magnitude()
//r := []float64{}

//for i := 0; i < v.dimension; i++ {
//r = append(r, v.coordinates[i]/m)
//}

//return New(r)
//}

//func (v Vector) Angle(w Vector, inDegree bool) float64 {
//d := v.Dot(w)
//mag := v.Magnitude() * w.Magnitude()

//if inDegree {
//return math.Acos(d/mag) * 180.0 / math.Pi
//}

//return math.Acos(d / mag)
//}

//func (v Vector) Paraller(w Vector) bool {
//a := v.Angle(w, true)
//if a == 0 || a == 180 {
//return true
//}

//return false
//}

//func (v Vector) Orthogonal(w Vector) bool {
//if v.Angle(w, true) == 90 {
//return true
//}

//return false
//}

//func (v Vector) Proj(w Vector) Vector {
//return v.Normalize().Scalar(v.Normalize().Dot(w))
//}
