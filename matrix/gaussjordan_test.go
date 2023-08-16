package matrix

import (
	"testing"
)

func TestNewMatrix(t *testing.T) {
	a := NewMatrix(25)
	if a == nil {
		t.Error("a == nil")
	}
	if Len(a) != 25 {
		t.Error("Len(a) != 25")
	}
}

func TestGaussJ1(t *testing.T) {
	a := NewMatrixFrom(3, [][]float64{{2, 1, -1}, {-3, -1, 2}, {-2, 1, 2}})
	b := NewVecFrom(3, []float64{8, -11, -3})
	GaussJ(a, b)
	if !a.isDiagonal() {
		t.Error("a is not diagonal")
	}
	if !a.Equal(NewMatrixFrom(3, [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})) {
		t.Error("a != I")
	}
	if !b.Equal(NewVecFrom(3, []float64{2, 3, -1})) {
		t.Error("b != [2, 3, -1]")
	}	
}