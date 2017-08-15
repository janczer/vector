package vector

import "testing"

func TestPrint(t *testing.T) {
	v := New([]float64{1, 2, 3})
	//v2 := New([]float64{1, 2, 3})

	if v.Print() != "Vector: 1.00, 2.00, 3.00" {
		t.Errorf("Expect %s. Got: %s", "Vector: 1, 2, 3", v.Print())
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
			t.Errorf("%s and %s. Expected equals: %t, got %t", v1.Print(), v2.Print(), a.expected, v1.Eq(v2))
		}
	}
}
