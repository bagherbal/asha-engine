package linear

import (
	"fmt"
	"math"
)

type Matrix struct {
	rows int
	cols int
	data []float64
}

func NewMatrix(rows, cols int) Matrix {
	return Matrix{rows: rows, cols: cols, data: make([]float64, rows*cols)}
}

func Identity(n int) Matrix {
	m := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		m.Set(i, i, 1)
	}
	return m
}

func (m Matrix) Rows() int { return m.rows }
func (m Matrix) Cols() int { return m.cols }

func (m Matrix) At(r, c int) float64 {
	return m.data[r*m.cols+c]
}

func (m Matrix) Set(r, c int, v float64) {
	m.data[r*m.cols+c] = v
}

func (m Matrix) Transpose() Matrix {
	out := NewMatrix(m.cols, m.rows)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			out.Set(c, r, m.At(r, c))
		}
	}
	return out
}

func (m Matrix) Add(n Matrix) (Matrix, error) {
	if m.rows != n.rows || m.cols != n.cols {
		return Matrix{}, fmt.Errorf("dimension mismatch: %dx%d + %dx%d", m.rows, m.cols, n.rows, n.cols)
	}
	out := NewMatrix(m.rows, m.cols)
	for i := range out.data {
		out.data[i] = m.data[i] + n.data[i]
	}
	return out, nil
}

func (m Matrix) Sub(n Matrix) (Matrix, error) {
	if m.rows != n.rows || m.cols != n.cols {
		return Matrix{}, fmt.Errorf("dimension mismatch: %dx%d - %dx%d", m.rows, m.cols, n.rows, n.cols)
	}
	out := NewMatrix(m.rows, m.cols)
	for i := range out.data {
		out.data[i] = m.data[i] - n.data[i]
	}
	return out, nil
}

func (m Matrix) Mul(n Matrix) (Matrix, error) {
	if m.cols != n.rows {
		return Matrix{}, fmt.Errorf("dimension mismatch: %dx%d * %dx%d", m.rows, m.cols, n.rows, n.cols)
	}
	out := NewMatrix(m.rows, n.cols)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < n.cols; c++ {
			sum := 0.0
			for k := 0; k < m.cols; k++ {
				sum += m.At(r, k) * n.At(k, c)
			}
			out.Set(r, c, sum)
		}
	}
	return out, nil
}

func (m Matrix) FrobeniusNorm() float64 {
	sum := 0.0
	for _, v := range m.data {
		sum += v * v
	}
	return math.Sqrt(sum)
}

func (m Matrix) AlmostEqual(n Matrix, eps float64) bool {
	if m.rows != n.rows || m.cols != n.cols {
		return false
	}
	for i := range m.data {
		if math.Abs(m.data[i]-n.data[i]) > eps {
			return false
		}
	}
	return true
}
