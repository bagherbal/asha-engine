// Package bsector implements the first finite variational gate of the Asha
// engine: the Boolean B-field vacuum action.
//
// The geometric input is the Boolean--Octonionic contact construction inside
// Λ⁴R⁸. The dynamical object is the positive semidefinite operator
//
//	O_B = Wᵀ (I - P_G) W
//
// on Boolean coordinates, where W : C³ -> Λ⁴R⁸ is the normalized Boolean
// incidence isometry and P_G is the octonionic G₂ calibration projector.
//
// For a Boolean field B = W b, the finite action is
//
//	S_B[b] = ||(I - P_G) W b||² = bᵀ O_B b.
//
// The theorem checked by this package is that ker(O_B) is exactly the contact
// vacuum K = Im(P_B) ∩ Im(P_G), expressed in Boolean coordinates.
package bsector

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Vacuum struct {
	Contact           contact.Space
	Operator          linear.Matrix // O_B on Boolean coordinates, dimension 56x56.
	Eigenvalues       []float64
	Eigenvectors      linear.Matrix // columns correspond to Eigenvalues, sorted ascending.
	ZeroModeFrame     linear.Matrix // Boolean-coordinate frame for ker(O_B).
	ZeroModeProjector linear.Projector
	ContactInBoolean  linear.Matrix    // Wᵀ Q_K.
	ContactProjector  linear.Projector // Boolean-coordinate projector onto K.
}

func BuildDefault() (Vacuum, error) {
	space, err := contact.BuildDefault()
	if err != nil {
		return Vacuum{}, err
	}
	return Build(space, 1e-8)
}

func Build(space contact.Space, eps float64) (Vacuum, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	w := space.BooleanSupport.Normalized
	pg := space.G2Support.Support.Matrix
	ambient := pg.Rows()
	if ambient != pg.Cols() || ambient != w.Rows() {
		return Vacuum{}, fmt.Errorf("ambient mismatch: P_G is %dx%d, W is %dx%d", pg.Rows(), pg.Cols(), w.Rows(), w.Cols())
	}

	identityMinusPG, err := linear.Identity(ambient).Sub(pg)
	if err != nil {
		return Vacuum{}, err
	}
	left, err := w.Transpose().Mul(identityMinusPG)
	if err != nil {
		return Vacuum{}, err
	}
	operator, err := left.Mul(w)
	if err != nil {
		return Vacuum{}, err
	}
	if !operator.IsSymmetric(1000 * eps) {
		return Vacuum{}, fmt.Errorf("B-sector operator is not symmetric")
	}

	eig, err := linear.SymmetricEigenJacobi(operator, 1e-13, 0)
	if err != nil {
		return Vacuum{}, err
	}
	values, vectors, err := sortEigenAscending(eig.Values, eig.Vectors)
	if err != nil {
		return Vacuum{}, err
	}

	zeroCols := make([]int, 0)
	for i, value := range values {
		if math.Abs(value) < eps {
			zeroCols = append(zeroCols, i)
		}
	}
	if len(zeroCols) == 0 {
		return Vacuum{}, fmt.Errorf("B-sector operator has no zero modes")
	}

	zeroFrame := linear.NewMatrix(vectors.Rows(), len(zeroCols))
	for newCol, oldCol := range zeroCols {
		for row := 0; row < vectors.Rows(); row++ {
			zeroFrame.Set(row, newCol, vectors.At(row, oldCol))
		}
	}
	zeroProjectorMatrix, err := zeroFrame.Mul(zeroFrame.Transpose())
	if err != nil {
		return Vacuum{}, err
	}
	zeroProjector, err := linear.NewProjector("ker(O_B)", zeroProjectorMatrix)
	if err != nil {
		return Vacuum{}, err
	}

	contactInBoolean, err := w.Transpose().Mul(space.ContactFrame)
	if err != nil {
		return Vacuum{}, err
	}
	contactProjectorMatrix, err := contactInBoolean.Mul(contactInBoolean.Transpose())
	if err != nil {
		return Vacuum{}, err
	}
	contactProjector, err := linear.NewProjector("WᵀP_KW", contactProjectorMatrix)
	if err != nil {
		return Vacuum{}, err
	}

	return Vacuum{
		Contact:           space,
		Operator:          operator,
		Eigenvalues:       values,
		Eigenvectors:      vectors,
		ZeroModeFrame:     zeroFrame,
		ZeroModeProjector: zeroProjector,
		ContactInBoolean:  contactInBoolean,
		ContactProjector:  contactProjector,
	}, nil
}

func (v Vacuum) BooleanDimension() int { return v.Operator.Rows() }

func (v Vacuum) ZeroModeDimension(eps float64) int {
	count := 0
	for _, value := range v.Eigenvalues {
		if math.Abs(value) < eps {
			count++
		}
	}
	return count
}

func (v Vacuum) FirstPositiveEigenvalue(eps float64) float64 {
	for _, value := range v.Eigenvalues {
		if value > eps {
			return value
		}
	}
	return math.NaN()
}

func (v Vacuum) OperatorSymmetryResidual() (float64, error) {
	diff, err := v.Operator.Sub(v.Operator.Transpose())
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}

func (v Vacuum) ZeroModeActionResidual() (float64, error) {
	image, err := v.Operator.Mul(v.ZeroModeFrame)
	if err != nil {
		return 0, err
	}
	return image.FrobeniusNorm(), nil
}

func (v Vacuum) ContactActionResidual() (float64, error) {
	image, err := v.Operator.Mul(v.ContactInBoolean)
	if err != nil {
		return 0, err
	}
	return image.FrobeniusNorm(), nil
}

func (v Vacuum) KernelEqualsContactResidual() (float64, error) {
	diff, err := v.ZeroModeProjector.Matrix.Sub(v.ContactProjector.Matrix)
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}

func (v Vacuum) ContactFrameIsometryResidual() (float64, error) {
	qtq, err := v.ContactInBoolean.Transpose().Mul(v.ContactInBoolean)
	if err != nil {
		return 0, err
	}
	diff, err := qtq.Sub(linear.Identity(v.ContactInBoolean.Cols()))
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}

func (v Vacuum) NegativeEigenvalueCount(eps float64) int {
	count := 0
	for _, value := range v.Eigenvalues {
		if value < -eps {
			count++
		}
	}
	return count
}

func sortEigenAscending(values []float64, vectors linear.Matrix) ([]float64, linear.Matrix, error) {
	if len(values) != vectors.Cols() || vectors.Rows() != vectors.Cols() {
		return nil, linear.Matrix{}, fmt.Errorf("eigenvectors must be square with one column per eigenvalue")
	}
	idx := make([]int, len(values))
	for i := range idx {
		idx[i] = i
	}
	for i := 0; i < len(idx); i++ {
		min := i
		for j := i + 1; j < len(idx); j++ {
			if values[idx[j]] < values[idx[min]] {
				min = j
			}
		}
		idx[i], idx[min] = idx[min], idx[i]
	}
	sortedValues := make([]float64, len(values))
	sortedVectors := linear.NewMatrix(vectors.Rows(), vectors.Cols())
	for newCol, oldCol := range idx {
		sortedValues[newCol] = values[oldCol]
		for row := 0; row < vectors.Rows(); row++ {
			sortedVectors.Set(row, newCol, vectors.At(row, oldCol))
		}
	}
	return sortedValues, sortedVectors, nil
}
