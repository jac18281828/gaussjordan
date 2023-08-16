package matrix

import (
	"errors"
)

type matrix struct {
	data [][]float64
	dataLen uint
}

func (a *matrix) len() uint {
	return Len(a)
}

func (a *matrix) get(i uint, j uint) float64 {
	return Get(a, i, j)
}

func (a *matrix) set(i uint, j uint, val float64) {
	a.data[i][j] = val
}

func (a *matrix) isVector() bool {
	return a.dataLen == 0
}

func (a *matrix) isSquare() bool {
	return a.dataLen == a.len()
}

func (a *matrix) isDiagonal() bool {
	if !a.isSquare() {
		return false
	}
	for i := uint(0); i < a.len(); i++ {
		for j := uint(0); j < a.len(); j++ {
			if i != j && a.get(i, j) != 0 {
				return false
			}
		}
	}
	return true
}

func (a *matrix) isSymmetric() bool {
	if !a.isSquare() {
		return false
	}
	for i := uint(0); i < a.len(); i++ {
		for j := uint(0); j < a.len(); j++ {
			if a.get(i, j) != a.get(j, i) {
				return false
			}
		}
	}
	return true
}

func (a *matrix) Equal(b *matrix) bool {
	if a.len() != b.len() && a.isVector() != b.isVector() {
		return false
	}
	for i := uint(0); i < a.len(); i++ {
		var jlen uint = a.len();
		if a.isVector() {
			jlen = 1;
		}
		for j := uint(0); j < jlen; j++ {
			if a.get(i, j) != b.get(i, j) {
				return false
			}
		}
	}
	return true
}

func NewVec(n int) *matrix {
	mat := make([][]float64, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]float64, 1)
	}
	return &matrix{mat, 0}
}

func NewVecFrom(n int, init []float64) *matrix {
	mat := NewVec(n)
	for i := 0; i < n; i++ {
		mat.data[i][0] = init[i]
	}
	return mat
}

func NewMatrix(n uint) *matrix {
	mat := make([][]float64, n)
	for i := uint(0); i < n; i++ {
		mat[i] = make([]float64, n)
	}
	return &matrix{mat, n}
}

func NewMatrixFrom(n uint, init [][]float64) *matrix {
	mat := NewMatrix(n)
	for i := uint(0); i < n; i++ {
		for j := uint(0); j < n; j++ {
			mat.data[i][j] = init[i][j]
		}
	}
	return mat
}

func Len(a *matrix) uint {
	return uint(len((*a).data))
}

func Get(a *matrix, i uint, j uint) float64 {
	return (*a).data[i][j]
}

func GetMaxForCol(a *matrix, col uint) (max float64, maxIndex uint) {
	n := Len(a)
	max = Get(a, col, col)
	maxIndex = col
	for j := col + 1; j < n; j++ {
		if max < Get(a, j, col) {
			max = Get(a, j, col)
			maxIndex = j
		}
	}
	return max, maxIndex
}

// gauss jordan elimination method
func GaussJ(a *matrix, b *matrix) (error) {
	n := Len(a)
	if n != Len(b) {
		return errors.New("matrix size is not equal")
	}
	for i := uint(0); i < n; i++ {
		// find max value
		max, maxIndex := GetMaxForCol(a, i)
		// swap row
		for j := uint(0); j < n; j++ {
			tmp := Get(a, i, j)
			a.data[i][j] = Get(a, maxIndex, j)
			a.data[maxIndex][j] = tmp
		}
		tmp := Get(b, i, 0)
		b.data[i][0] = Get(b, maxIndex, 0)
		b.data[maxIndex][0] = tmp
		// divide row
		for j := uint(0); j < n; j++ {
			a.data[i][j] /= max
		}
		b.data[i][0] /= max
		// subtract row
		for j := uint(0); j < n; j++ {
			if i == j {
				continue
			}
			tmp := Get(a, j, i)
			for k := uint(0); k < n; k++ {
				a.data[j][k] -= tmp * Get(a, i, k)
			}
			b.data[j][0] -= tmp * Get(b, i, 0)
		}
	}
	return nil
}