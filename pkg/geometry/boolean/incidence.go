package boolean

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type IncidenceSupport struct {
	VectorDimension int
	LowerGrade      int
	UpperGrade      int
	LowerBasis      []combinatorics.Subset
	UpperBasis      []combinatorics.Subset
	Incidence       linear.Matrix
	Gram            linear.Matrix
	Normalized      linear.Matrix
	Support         linear.Projector
	GramEigenvalues []float64
}

func BuildIncidenceSupport(vectorDimension, lowerGrade, upperGrade int) (IncidenceSupport, error) {
	if upperGrade != lowerGrade+1 {
		return IncidenceSupport{}, fmt.Errorf("Boolean incidence currently expects adjacent grades: got %d -> %d", lowerGrade, upperGrade)
	}
	lower, err := combinatorics.Subsets(vectorDimension, lowerGrade)
	if err != nil {
		return IncidenceSupport{}, err
	}
	upper, err := combinatorics.Subsets(vectorDimension, upperGrade)
	if err != nil {
		return IncidenceSupport{}, err
	}

	m := linear.NewMatrix(len(upper), len(lower))
	for r, upperSubset := range upper {
		for c, lowerSubset := range lower {
			if upperSubset.ContainsAll(lowerSubset) {
				m.Set(r, c, 1)
			}
		}
	}

	gram, err := m.Transpose().Mul(m)
	if err != nil {
		return IncidenceSupport{}, err
	}

	eig, err := linear.SymmetricEigenJacobi(gram, 1e-13, 0)
	if err != nil {
		return IncidenceSupport{}, err
	}
	values, vectors, err := linear.SortEigenDescending(eig.Values, eig.Vectors)
	if err != nil {
		return IncidenceSupport{}, err
	}
	invSqrt := linear.NewMatrix(len(values), len(values))
	for i, value := range values {
		if value <= 0 {
			return IncidenceSupport{}, fmt.Errorf("incidence Gram matrix has non-positive eigenvalue at %d: %.16g", i, value)
		}
		invSqrt.Set(i, i, 1/math.Sqrt(value))
	}

	// W = M U Λ^{-1/2} U^T.
	mu, err := m.Mul(vectors)
	if err != nil {
		return IncidenceSupport{}, err
	}
	muis, err := mu.Mul(invSqrt)
	if err != nil {
		return IncidenceSupport{}, err
	}
	w, err := muis.Mul(vectors.Transpose())
	if err != nil {
		return IncidenceSupport{}, err
	}

	supportMatrix, err := w.Mul(w.Transpose())
	if err != nil {
		return IncidenceSupport{}, err
	}
	support, err := linear.NewProjector("P_B", supportMatrix)
	if err != nil {
		return IncidenceSupport{}, err
	}

	return IncidenceSupport{
		VectorDimension: vectorDimension,
		LowerGrade:      lowerGrade,
		UpperGrade:      upperGrade,
		LowerBasis:      lower,
		UpperBasis:      upper,
		Incidence:       m,
		Gram:            gram,
		Normalized:      w,
		Support:         support,
		GramEigenvalues: values,
	}, nil
}

func (s IncidenceSupport) LowerDimension() int { return len(s.LowerBasis) }
func (s IncidenceSupport) UpperDimension() int { return len(s.UpperBasis) }

func (s IncidenceSupport) HarmonicResidueDimension() int {
	return s.UpperDimension() - s.LowerDimension()
}

func (s IncidenceSupport) IsometryResidual() (float64, error) {
	wtw, err := s.Normalized.Transpose().Mul(s.Normalized)
	if err != nil {
		return 0, err
	}
	diff, err := wtw.Sub(linear.Identity(s.LowerDimension()))
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}

func (s IncidenceSupport) RankFromGram(eps float64) int {
	return linear.RankFromEigenvalues(s.GramEigenvalues, eps)
}
