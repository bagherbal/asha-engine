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

func FromRows(values [][]float64) (Matrix, error) {
	if len(values) == 0 {
		return Matrix{}, fmt.Errorf("matrix must have at least one row")
	}
	cols := len(values[0])
	if cols == 0 {
		return Matrix{}, fmt.Errorf("matrix must have at least one column")
	}
	m := NewMatrix(len(values), cols)
	for r := range values {
		if len(values[r]) != cols {
			return Matrix{}, fmt.Errorf("ragged matrix at row %d: got %d columns, expected %d", r, len(values[r]), cols)
		}
		for c := range values[r] {
			m.Set(r, c, values[r][c])
		}
	}
	return m, nil
}

func Identity(n int) Matrix {
	m := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		m.Set(i, i, 1)
	}
	return m
}

func Diagonal(values []float64) Matrix {
	m := NewMatrix(len(values), len(values))
	for i, v := range values {
		m.Set(i, i, v)
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

func (m Matrix) Clone() Matrix {
	out := NewMatrix(m.rows, m.cols)
	copy(out.data, m.data)
	return out
}

func (m Matrix) DataCopy() []float64 {
	out := make([]float64, len(m.data))
	copy(out, m.data)
	return out
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

func (m Matrix) Scale(s float64) Matrix {
	out := NewMatrix(m.rows, m.cols)
	for i := range out.data {
		out.data[i] = s * m.data[i]
	}
	return out
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

func (m Matrix) Trace() (float64, error) {
	if m.rows != m.cols {
		return 0, fmt.Errorf("trace requires square matrix: %dx%d", m.rows, m.cols)
	}
	sum := 0.0
	for i := 0; i < m.rows; i++ {
		sum += m.At(i, i)
	}
	return sum, nil
}

func (m Matrix) FrobeniusNorm() float64 {
	sum := 0.0
	for _, v := range m.data {
		sum += v * v
	}
	return math.Sqrt(sum)
}

func (m Matrix) MaxAbs() float64 {
	max := 0.0
	for _, v := range m.data {
		if abs := math.Abs(v); abs > max {
			max = abs
		}
	}
	return max
}

func (m Matrix) MaxAbsDiff(n Matrix) (float64, error) {
	if m.rows != n.rows || m.cols != n.cols {
		return 0, fmt.Errorf("dimension mismatch: %dx%d vs %dx%d", m.rows, m.cols, n.rows, n.cols)
	}
	max := 0.0
	for i := range m.data {
		if d := math.Abs(m.data[i] - n.data[i]); d > max {
			max = d
		}
	}
	return max, nil
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

func (m Matrix) IsSymmetric(eps float64) bool {
	if m.rows != m.cols {
		return false
	}
	for r := 0; r < m.rows; r++ {
		for c := r + 1; c < m.cols; c++ {
			if math.Abs(m.At(r, c)-m.At(c, r)) > eps {
				return false
			}
		}
	}
	return true
}

func MustMul(a, b Matrix) Matrix {
	m, err := a.Mul(b)
	if err != nil {
		panic(err)
	}
	return m
}

func MustSub(a, b Matrix) Matrix {
	m, err := a.Sub(b)
	if err != nil {
		panic(err)
	}
	return m
}
