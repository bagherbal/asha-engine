package contact

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/geometry/boolean"
	"github.com/bagherbal/asha-engine/pkg/geometry/g2"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

// Space is the finite Boolean--Octonionic contact geometry inside Λ⁴R⁸.
//
// It is constructed from two independently defined projectors:
//
//	P_B: Boolean incidence support, rank 56.
//	P_G: octonionic G₂ calibration support, rank 14.
//
// The contact space is the exact intersection
//
//	K = Im(P_B) ∩ Im(P_G).
//
// Computationally, because Q_G is an orthonormal frame for Im(P_G), the
// intersection is the eigenspace with eigenvalue 1 of
//
//	Ω = Q_Gᵀ P_B Q_G.
type Space struct {
	BooleanSupport     boolean.IncidenceSupport
	G2Support          g2.CalibrationSupport
	Overlap            linear.Matrix
	OverlapEigenvalues []float64
	ContactFrame       linear.Matrix
	ContactProjector   linear.Projector
	ContactIndices     []int
	BareLeakage        linear.Matrix
}

func BuildDefault() (Space, error) {
	b, err := boolean.BuildIncidenceSupport(8, 3, 4)
	if err != nil {
		return Space{}, err
	}
	g, err := g2.BuildCalibrationSupport()
	if err != nil {
		return Space{}, err
	}
	return Build(b, g, 1e-8)
}

func Build(b boolean.IncidenceSupport, g g2.CalibrationSupport, eps float64) (Space, error) {
	if b.UpperDimension() != g.MiddleDimension() {
		return Space{}, fmt.Errorf("ambient chamber mismatch: Boolean upper dim %d, G₂ middle dim %d", b.UpperDimension(), g.MiddleDimension())
	}
	if eps <= 0 {
		eps = 1e-8
	}

	pbg, err := b.Support.Matrix.Mul(g.Orthonormal)
	if err != nil {
		return Space{}, err
	}
	overlap, err := g.Orthonormal.Transpose().Mul(pbg)
	if err != nil {
		return Space{}, err
	}
	if !overlap.IsSymmetric(100 * eps) {
		return Space{}, fmt.Errorf("contact overlap Q_GᵀP_BQ_G is not symmetric")
	}

	eig, err := linear.SymmetricEigenJacobi(overlap, 1e-13, 0)
	if err != nil {
		return Space{}, err
	}
	values, vectors, err := linear.SortEigenDescending(eig.Values, eig.Vectors)
	if err != nil {
		return Space{}, err
	}

	indices := make([]int, 0)
	for i, value := range values {
		if math.Abs(value-1) < eps {
			indices = append(indices, i)
		}
	}
	if len(indices) == 0 {
		return Space{}, fmt.Errorf("contact eigenspace is empty: no overlap eigenvalue near 1")
	}

	contactCoordinates := linear.NewMatrix(vectors.Rows(), len(indices))
	for newCol, oldCol := range indices {
		for row := 0; row < vectors.Rows(); row++ {
			contactCoordinates.Set(row, newCol, vectors.At(row, oldCol))
		}
	}
	contactFrame, err := g.Orthonormal.Mul(contactCoordinates)
	if err != nil {
		return Space{}, err
	}
	projectorMatrix, err := contactFrame.Mul(contactFrame.Transpose())
	if err != nil {
		return Space{}, err
	}
	projector, err := linear.NewProjector("P_K", projectorMatrix)
	if err != nil {
		return Space{}, err
	}

	pbpg, err := b.Support.Matrix.Mul(g.Support.Matrix)
	if err != nil {
		return Space{}, err
	}
	leakage, err := pbpg.Sub(projector.Matrix)
	if err != nil {
		return Space{}, err
	}

	return Space{
		BooleanSupport:     b,
		G2Support:          g,
		Overlap:            overlap,
		OverlapEigenvalues: values,
		ContactFrame:       contactFrame,
		ContactProjector:   projector,
		ContactIndices:     indices,
		BareLeakage:        leakage,
	}, nil
}

func (s Space) Dimension() int { return s.ContactFrame.Cols() }

func (s Space) AmbientDimension() int { return s.ContactFrame.Rows() }

func (s Space) ExpectedContactDenominator() int {
	// The G₂ calibration sector is 7_t ⊕ 7_s, so one contact copy has half
	// the sector dimension. This is computed from the constructed sector, not
	// from observed physics.
	return s.G2Support.SectorDimension() / 2
}

func (s Space) ContactIndex() float64 {
	denom := s.ExpectedContactDenominator()
	if denom == 0 {
		return math.NaN()
	}
	return float64(s.Dimension()) / float64(denom)
}

func (s Space) FrameIsometryResidual() (float64, error) {
	qtq, err := s.ContactFrame.Transpose().Mul(s.ContactFrame)
	if err != nil {
		return 0, err
	}
	diff, err := qtq.Sub(linear.Identity(s.Dimension()))
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}

func (s Space) BooleanContainmentResidual() (float64, error) {
	pbpk, err := s.BooleanSupport.Support.Matrix.Mul(s.ContactProjector.Matrix)
	if err != nil {
		return 0, err
	}
	diff, err := pbpk.Sub(s.ContactProjector.Matrix)
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}

func (s Space) G2ContainmentResidual() (float64, error) {
	pgpk, err := s.G2Support.Support.Matrix.Mul(s.ContactProjector.Matrix)
	if err != nil {
		return 0, err
	}
	diff, err := pgpk.Sub(s.ContactProjector.Matrix)
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}

func (s Space) BareLeakageNorm() float64 {
	return s.BareLeakage.FrobeniusNorm()
}

func (s Space) BareLeakageNormSquared() float64 {
	n := s.BareLeakage.FrobeniusNorm()
	return n * n
}

func (s Space) OverlapMultiplicityNear(target, eps float64) int {
	count := 0
	for _, value := range s.OverlapEigenvalues {
		if math.Abs(value-target) < eps {
			count++
		}
	}
	return count
}
