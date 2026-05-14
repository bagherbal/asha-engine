// Package higgs extracts the finite vacuum-mixing sector from the projected
// Boolean gauge connection.
//
// Gate 8 showed that strict compression PAP is not a Lie representation. Gate
// 10 explained why: the off-diagonal blocks PAQ and QAP were discarded. This
// package keeps those blocks and treats their sum
//
//	Φ_A = P_C A P_K + P_K A P_C
//
// as the finite second-fundamental/Higgs-like field associated with a Boolean-
// compressed connection generator A. No physical Higgs mass is inferred here;
// the gate only computes the canonical finite mixing operators and spectra.
package higgs

import (
	"sync"

	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/gauge/connection"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Analysis struct {
	Connection connection.Analysis

	HiggsFields          []linear.Matrix // Φ_i = P_C A_i P_K + P_K A_i P_C on Boolean coordinates.
	HiggsSpanRank        int
	HiggsSpanEigenvalues []float64

	VacuumMixingOperator     linear.Matrix // Σ (P_C A_i P_K)^T(P_C A_i P_K), supported on K.
	ComplementMixingOperator linear.Matrix // Σ (P_K A_i P_C)^T(P_K A_i P_C), supported on K⊥.

	ContactDimension         int
	VacuumMixingRank         int
	VacuumUnmixedDimension   int
	ComplementMixingRank     int
	VacuumMixingTrace        float64
	ComplementMixingTrace    float64
	VacuumMixingSpectrum     []float64
	ComplementMixingSpectrum []float64

	MaxOffDiagonalBlockResidual float64 // max ||PΦP||+||QΦQ||.
	MaxSkewResidual             float64 // max ||Φ+Φᵀ||.
	MaxPositiveResidual         float64 // most negative eigenvalue violation across positive operators.
	TotalMixingNormSquared      float64 // Σ ||Φ_i||².
}

var (
	higgsDefaultOnce  sync.Once
	higgsDefaultValue Analysis
	higgsDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	higgsDefaultOnce.Do(func() {
		higgsDefaultValue, higgsDefaultErr = buildHiggsDefaultUncached()
	})
	return higgsDefaultValue, higgsDefaultErr
}

func buildHiggsDefaultUncached() (Analysis, error) {
	a, err := connection.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(a, 1e-9)
}

func Build(a connection.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-9
	}
	p := a.Compression.BooleanComplementProjector
	q := a.Compression.BooleanContactProjector
	dim := p.Rows()

	vacuumMix := linear.NewMatrix(dim, dim)
	complementMix := linear.NewMatrix(dim, dim)
	higgsFields := make([]linear.Matrix, 0, len(a.Compression.BooleanGenerators))

	maxDiag := 0.0
	maxSkew := 0.0
	totalNormSq := 0.0

	for _, gen := range a.Compression.BooleanGenerators {
		paq, err := sandwich(p, gen, q)
		if err != nil {
			return Analysis{}, err
		}
		qap, err := sandwich(q, gen, p)
		if err != nil {
			return Analysis{}, err
		}
		phi, err := paq.Add(qap)
		if err != nil {
			return Analysis{}, err
		}
		higgsFields = append(higgsFields, phi)
		n := phi.FrobeniusNorm()
		totalNormSq += n * n

		pPhiP, err := sandwich(p, phi, p)
		if err != nil {
			return Analysis{}, err
		}
		qPhiQ, err := sandwich(q, phi, q)
		if err != nil {
			return Analysis{}, err
		}
		diag := pPhiP.FrobeniusNorm() + qPhiQ.FrobeniusNorm()
		if diag > maxDiag {
			maxDiag = diag
		}

		skew, err := phi.Add(phi.Transpose())
		if err != nil {
			return Analysis{}, err
		}
		if s := skew.FrobeniusNorm(); s > maxSkew {
			maxSkew = s
		}

		kt, err := paq.Transpose().Mul(paq)
		if err != nil {
			return Analysis{}, err
		}
		vacuumMix, err = vacuumMix.Add(kt)
		if err != nil {
			return Analysis{}, err
		}

		ct, err := qap.Transpose().Mul(qap)
		if err != nil {
			return Analysis{}, err
		}
		complementMix, err = complementMix.Add(ct)
		if err != nil {
			return Analysis{}, err
		}
	}

	spanRank, spanValues, err := operatorSpanRank(higgsFields, eps)
	if err != nil {
		return Analysis{}, err
	}

	vacValues, vacRank, vacTrace, vacNeg, err := spectrumSummary(vacuumMix, eps)
	if err != nil {
		return Analysis{}, err
	}
	compValues, compRank, compTrace, compNeg, err := spectrumSummary(complementMix, eps)
	if err != nil {
		return Analysis{}, err
	}

	maxPositiveResidual := math.Max(vacNeg, compNeg)
	qTrace, err := q.Trace()
	if err != nil {
		return Analysis{}, err
	}
	contactDim := int(math.Round(qTrace))

	return Analysis{
		Connection:                  a,
		HiggsFields:                 higgsFields,
		HiggsSpanRank:               spanRank,
		HiggsSpanEigenvalues:        spanValues,
		VacuumMixingOperator:        vacuumMix,
		ComplementMixingOperator:    complementMix,
		ContactDimension:            contactDim,
		VacuumMixingRank:            vacRank,
		VacuumUnmixedDimension:      contactDim - vacRank,
		ComplementMixingRank:        compRank,
		VacuumMixingTrace:           vacTrace,
		ComplementMixingTrace:       compTrace,
		VacuumMixingSpectrum:        vacValues,
		ComplementMixingSpectrum:    compValues,
		MaxOffDiagonalBlockResidual: maxDiag,
		MaxSkewResidual:             maxSkew,
		MaxPositiveResidual:         maxPositiveResidual,
		TotalMixingNormSquared:      totalNormSq,
	}, nil
}

func sandwich(left, middle, right linear.Matrix) (linear.Matrix, error) {
	lm, err := left.Mul(middle)
	if err != nil {
		return linear.Matrix{}, err
	}
	return lm.Mul(right)
}

func operatorSpanRank(fields []linear.Matrix, eps float64) (int, []float64, error) {
	if len(fields) == 0 {
		return 0, nil, fmt.Errorf("empty Higgs field list")
	}
	rows := fields[0].Rows()
	cols := fields[0].Cols()
	raw := linear.NewMatrix(rows*cols, len(fields))
	for c, f := range fields {
		if f.Rows() != rows || f.Cols() != cols {
			return 0, nil, fmt.Errorf("Higgs field shape mismatch")
		}
		idx := 0
		for r := 0; r < rows; r++ {
			for k := 0; k < cols; k++ {
				raw.Set(idx, c, f.At(r, k))
				idx++
			}
		}
	}
	gram, err := raw.Transpose().Mul(raw)
	if err != nil {
		return 0, nil, err
	}
	eig, err := linear.SymmetricEigenJacobi(gram, 1e-13, 0)
	if err != nil {
		return 0, nil, err
	}
	values, _, err := linear.SortEigenDescending(eig.Values, eig.Vectors)
	if err != nil {
		return 0, nil, err
	}
	rank := 0
	for _, value := range values {
		if value > eps {
			rank++
		}
	}
	return rank, values, nil
}

func spectrumSummary(m linear.Matrix, eps float64) ([]float64, int, float64, float64, error) {
	if !m.IsSymmetric(1e-8) {
		sym, err := m.Add(m.Transpose())
		if err != nil {
			return nil, 0, 0, 0, err
		}
		m = sym.Scale(0.5)
	}
	eig, err := linear.SymmetricEigenJacobi(m, 1e-12, 0)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	values, _, err := linear.SortEigenDescending(eig.Values, eig.Vectors)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	rank := 0
	trace := 0.0
	mostNegativeViolation := 0.0
	for _, value := range values {
		trace += value
		if value > eps {
			rank++
		}
		if value < -eps && -value > mostNegativeViolation {
			mostNegativeViolation = -value
		}
	}
	return values, rank, trace, mostNegativeViolation, nil
}

func TopSpectrum(values []float64, n int) []float64 {
	if n > len(values) {
		n = len(values)
	}
	out := make([]float64, n)
	copy(out, values[:n])
	return out
}
