package linear

import "fmt"

type Projector struct {
	Name   string
	Matrix Matrix
}

func NewProjector(name string, matrix Matrix) (Projector, error) {
	if matrix.Rows() != matrix.Cols() {
		return Projector{}, fmt.Errorf("projector %s must be square: %dx%d", name, matrix.Rows(), matrix.Cols())
	}
	return Projector{Name: name, Matrix: matrix}, nil
}

func (p Projector) Trace() (float64, error) {
	return p.Matrix.Trace()
}

func (p Projector) IdempotenceResidual() (float64, error) {
	sq, err := p.Matrix.Mul(p.Matrix)
	if err != nil {
		return 0, err
	}
	diff, err := sq.Sub(p.Matrix)
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}

func (p Projector) SymmetryResidual() (float64, error) {
	diff, err := p.Matrix.Sub(p.Matrix.Transpose())
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}
