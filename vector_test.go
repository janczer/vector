package vector

import (
	"math"
	"testing"
)

func TestPrint(t *testing.T) {
	v := New([]float64{1, 2, 3})
	//v2 := New([]float64{1, 2, 3})

	if v.String() != "Vector: 1.00, 2.00, 3.00" {
		t.Errorf("Expect %s. Got: %s", "Vector: 1, 2, 3", v)
	}
}

func TestEq(t *testing.T) {
	tests := []struct {
		vector1  []float64
		vector2  []float64
		expected bool
	}{
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, true},
		{[]float64{0}, []float64{0}, true},
		{[]float64{2, 2, 3}, []float64{1, 2, 3}, false},
		{[]float64{1}, []float64{2}, false},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		if v1.Eq(v2) != a.expected {
			t.Errorf("%s and %s. Expected equals: %t, got %t", v1, v2, a.expected, v1.Eq(v2))
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		vector1  []float64
		vector2  []float64
		expected Vector
	}{
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, New([]float64{2, 4, 6})},
		{[]float64{0, 1, 3}, []float64{5, 1, 3}, New([]float64{5, 2, 6})},
		{[]float64{0}, []float64{0}, New([]float64{0})},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		s := v1.Add(v2)

		if !s.Eq(a.expected) {
			t.Errorf("%s + %s should be %s, got %s", v1, v2, a.expected, s)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		vector1  []float64
		vector2  []float64
		expected Vector
	}{
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, New([]float64{0, 0, 0})},
		{[]float64{2, 2, 2}, []float64{1, 1, 1}, New([]float64{1, 1, 1})},
		{[]float64{2, 2, 2}, []float64{5, 8, 2}, New([]float64{-3, -6, 0})},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		s := v1.Sub(v2)

		if !s.Eq(a.expected) {
			t.Errorf("%s - %s should be %s, got %s", v1, v2, a.expected, s)
		}
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		vector1  []float64
		vector2  []float64
		expected float64
	}{
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, 14},
		{[]float64{2, 2, 2}, []float64{1, 1, 1}, 6},
		{[]float64{2, 2, 2}, []float64{5, 8, 2}, 30},
		{[]float64{1, 3, 4}, []float64{0, 0, 0}, 0},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		s := v1.Dot(v2)

		if s != a.expected {
			t.Errorf("%s . %s should be %.2f, got %.2f", v1, v2, a.expected, s)
		}
	}
}

func TestScalar(t *testing.T) {
	tests := []struct {
		vector1  []float64
		num      float64
		expected Vector
	}{
		{[]float64{1, 2, 3}, 2, New([]float64{2, 4, 6})},
		{[]float64{2, 2, 2}, 1, New([]float64{2, 2, 2})},
		{[]float64{2, 2, 2}, 8, New([]float64{16, 16, 16})},
		{[]float64{1, 3, 4}, 0, New([]float64{0, 0, 0})},
	}

	for _, a := range tests {
		v1 := New(a.vector1)

		s := v1.Scalar(a.num)

		if !s.Eq(a.expected) {
			t.Errorf("%s * %.2f should be %s, got %s", v1, a.num, a.expected, s)
		}
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		vector1 []float64
		m       float64
	}{
		{[]float64{1, 2, 3}, math.Sqrt(14)},
		{[]float64{2, 2, 2}, math.Sqrt(12)},
		{[]float64{1, 3, 4}, math.Sqrt(26)},
		{[]float64{-1, 0, -8}, math.Sqrt(65)},
	}

	for _, a := range tests {
		v1 := New(a.vector1)

		if v1.Magnitude() != a.m {
			t.Errorf("Magnitude of %s shoulde be %f, got %f", v1, a.m, v1.Magnitude())
		}
	}
}

func TestNormalize(t *testing.T) {
	e1 := math.Sqrt(14)
	e2 := math.Sqrt(12)
	e3 := math.Sqrt(65)

	tests := []struct {
		vector1 []float64
		vector2 []float64
	}{
		{[]float64{1, 2, 3}, []float64{1 / e1, 2 / e1, 3 / e1}},
		{[]float64{2, 2, 2}, []float64{2 / e2, 2 / e2, 2 / e2}},
		{[]float64{-1, 0, -8}, []float64{-1 / e3, 0, -8 / e3}},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		if !v2.Eq(v1.Normalize()) {
			t.Errorf("%s normalized to %s, got %s", v1, v2, v1.Normalize())
		}
	}
}

func TestAngle(t *testing.T) {
	tests := []struct {
		vector1  []float64
		vector2  []float64
		expected float64
	}{
		{[]float64{3, 0}, []float64{3, 3}, 45},
		{[]float64{3, 0}, []float64{0, 3}, 90},
		{[]float64{1, 1}, []float64{-1, -1}, 180},
		{[]float64{1, 1}, []float64{1, 1}, 0},
		{[]float64{3, 3}, []float64{3, 3}, 0},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		if a.expected != v1.Angle(v2, true) {
			t.Errorf("Angle between %s and %s is %.16f, got %.16f", v1, v2, a.expected, v1.Angle(v2, true))
		}
	}
}

func TestParaller(t *testing.T) {
	tests := []struct {
		vector1  []float64
		vector2  []float64
		expected bool
	}{
		{[]float64{3, 0}, []float64{3, 3}, false},
		{[]float64{30, 30}, []float64{30, 30}, true},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		if a.expected != v1.Paraller(v2) {
			t.Errorf("Parraler %s and %s is %t, got %t", v1, v2, a.expected, v1.Paraller(v2))
		}
	}
}

func TestOrthogonal(t *testing.T) {
	tests := []struct {
		vector1  []float64
		vector2  []float64
		expected bool
	}{
		{[]float64{3, 0}, []float64{0, 3}, true},
		{[]float64{30, 30}, []float64{30, 30}, false},
	}

	for _, a := range tests {
		v1 := New(a.vector1)
		v2 := New(a.vector2)

		if a.expected != v1.Orthogonal(v2) {
			t.Errorf("Orthogonal %s and %s is %t, got %t", v1, v2, a.expected, v1.Paraller(v2))
		}
	}
}
