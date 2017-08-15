package vector

import (
	//"math"
	//"fmt"
	"math/big"
	"testing"
)

func TestPrint(t *testing.T) {
	v := New([]*big.Float{
		big.NewFloat(1),
		big.NewFloat(2),
		big.NewFloat(3),
	})

	if v.Print() != "Vector: 1.00000, 2.00000, 3.00000" {
		t.Errorf("Expect %s. Got: %s", "Vector: 1, 2, 3", v.Print())
	}
}

func TestEq(t *testing.T) {
	tests := []struct {
		vector1  []*big.Float
		vector2  []*big.Float
		expected bool
	}{
		{
			[]*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3)},
			[]*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3)},
			true,
		},
		{
			[]*big.Float{big.NewFloat(2), big.NewFloat(2), big.NewFloat(3)},
			[]*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3)},
			false,
		},
		{
			[]*big.Float{big.NewFloat(0)},
			[]*big.Float{big.NewFloat(0)},
			true,
		},
		{
			[]*big.Float{big.NewFloat(1)},
			[]*big.Float{big.NewFloat(2)},
			false,
		},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		if v1.Eq(v2) != a.expected {
			t.Errorf("%s and %s. Expected equals: %t, got %t", v1.Print(), v2.Print(), a.expected, v1.Eq(v2))
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		vector1  []*big.Float
		vector2  []*big.Float
		expected Vector
	}{
		{
			[]*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3)},
			[]*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3)},
			New([]*big.Float{big.NewFloat(2), big.NewFloat(4), big.NewFloat(6)}),
		},
		{
			[]*big.Float{big.NewFloat(0), big.NewFloat(1), big.NewFloat(3)},
			[]*big.Float{big.NewFloat(5), big.NewFloat(1), big.NewFloat(3)},
			New([]*big.Float{big.NewFloat(5), big.NewFloat(2), big.NewFloat(6)}),
		},
		{
			[]*big.Float{big.NewFloat(0.111111), big.NewFloat(1222.11983), big.NewFloat(0.223)},
			[]*big.Float{big.NewFloat(5.111111), big.NewFloat(1.11100001), big.NewFloat(30.123311)},
			New([]*big.Float{big.NewFloat(5.222222), big.NewFloat(1223.23083001), big.NewFloat(30.346311)}),
		},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		s := v1.Add(v2)

		if !s.Eq(a.expected) {
			t.Errorf("%s + %s should be %s, got %s", v1.Print(), v2.Print(), a.expected.Print(), s.Print())
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		vector1  []*big.Float
		vector2  []*big.Float
		expected Vector
	}{
		{
			[]*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3)},
			[]*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3)},
			New([]*big.Float{big.NewFloat(0), big.NewFloat(0), big.NewFloat(0)}),
		},
		{
			[]*big.Float{big.NewFloat(2), big.NewFloat(2), big.NewFloat(2)},
			[]*big.Float{big.NewFloat(1), big.NewFloat(1), big.NewFloat(1)},
			New([]*big.Float{big.NewFloat(1), big.NewFloat(1), big.NewFloat(1)}),
		},
		{
			[]*big.Float{big.NewFloat(0.111111), big.NewFloat(1222.11983), big.NewFloat(0.223)},
			[]*big.Float{big.NewFloat(5.111111), big.NewFloat(1.11100001), big.NewFloat(30.123311)},
			New([]*big.Float{big.NewFloat(-5.0), big.NewFloat(1221.00882999), big.NewFloat(-29.9003110)}),
		},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		s := v1.Sub(v2)

		if !s.Eq(a.expected) {
			t.Errorf("%s - %s should be %s, got %s", v1.Print(), v2.Print(), a.expected.Print(), s.Print())
		}
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		vector1  []*big.Float
		vector2  []*big.Float
		expected *big.Float
	}{
		{
			[]*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3)},
			[]*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3)},
			big.NewFloat(14),
		},
		{
			[]*big.Float{big.NewFloat(2), big.NewFloat(2), big.NewFloat(2)},
			[]*big.Float{big.NewFloat(1), big.NewFloat(1), big.NewFloat(1)},
			big.NewFloat(6),
		},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		s := v1.Dot(v2)

		if a.expected.Cmp(s) != 0 {
			t.Errorf("%s . %s should be %f, got %.50f", v1.Print(), v2.Print(), a.expected, s)
		}
	}
}

//func TestScalar(t *testing.T) {
//tests := []struct {
//vector1  []float64
//num      float64
//expected Vector
//}{
//{[]float64{1, 2, 3}, 2, New([]float64{2, 4, 6})},
//{[]float64{2, 2, 2}, 1, New([]float64{2, 2, 2})},
//{[]float64{2, 2, 2}, 8, New([]float64{16, 16, 16})},
//{[]float64{1, 3, 4}, 0, New([]float64{0, 0, 0})},
//}

//for _, a := range tests {
//v1 := New(a.vector1)

//s := v1.Scalar(a.num)

//if !s.Eq(a.expected) {
//t.Errorf("%s * %.2f should be %s, got %s", v1.Print(), a.num, a.expected.Print(), s.Print())
//}
//}
//}

//func TestMagnitude(t *testing.T) {
//tests := []struct {
//vector1 []float64
//m       float64
//}{
//{[]float64{1, 2, 3}, math.Sqrt(14)},
//{[]float64{2, 2, 2}, math.Sqrt(12)},
//{[]float64{1, 3, 4}, math.Sqrt(26)},
//{[]float64{-1, 0, -8}, math.Sqrt(65)},
//}

//for _, a := range tests {
//v1 := New(a.vector1)

//if v1.Magnitude() != a.m {
//t.Errorf("Magnitude of %s shoulde be %f, got %f", v1.Print(), a.m, v1.Magnitude())
//}
//}
//}

//func TestNormalize(t *testing.T) {
//e1 := math.Sqrt(14)
//e2 := math.Sqrt(12)
//e3 := math.Sqrt(65)

//tests := []struct {
//vector1 []float64
//vector2 []float64
//}{
//{[]float64{1, 2, 3}, []float64{1 / e1, 2 / e1, 3 / e1}},
//{[]float64{2, 2, 2}, []float64{2 / e2, 2 / e2, 2 / e2}},
//{[]float64{-1, 0, -8}, []float64{-1 / e3, 0, -8 / e3}},
//}

//for _, a := range tests {
//v1 := New(a.vector1)
//v2 := New(a.vector2)

//if !v2.Eq(v1.Normalize()) {
//t.Errorf("%s normalized to %s, got %s", v1.Print(), v2.Print(), v1.Normalize().Print())
//}
//}
//}

//func TestAngle(t *testing.T) {
//tests := []struct {
//vector1  []float64
//vector2  []float64
//expected float64
//}{
//{[]float64{3, 0}, []float64{3, 3}, 45},
////{[]float64{3, 0}, []float64{0, 3}, 90},
////{[]float64{1, 1}, []float64{1, 1}, 0},
//}

//for _, a := range tests {
//v1 := New(a.vector1)
//v2 := New(a.vector2)

//if a.expected != v1.Angle(v2, true) {
//t.Errorf("Angle between %s and %s is %.16f, got %.16f", v1.Print(), v2.Print(), a.expected, v1.Angle(v2, true))
//}
//}
//}

//func TestParaller(t *testing.T) {
//tests := []struct {
//vector1  []float64
//vector2  []float64
//expected bool
//}{
//{[]float64{3, 0}, []float64{3, 3}, false},
////{[]float64{3, 3}, []float64{3, 3}, true},
//}

//for _, a := range tests {
//v1 := New(a.vector1)
//v2 := New(a.vector2)

//if a.expected != v1.Paraller(v2) {
//t.Errorf("Parraler %s and %s is %t, got %t", v1.Print(), v2.Print(), a.expected, v1.Paraller(v2))
//}
//}
//}
