// Package higgspotential builds the first finite Higgs-potential candidate from
// the vacuum-mixing sector discovered by the projected Boolean connection.
//
// This package is deliberately conservative. It does not infer the observed
// Higgs boson mass, the electroweak scale, or a Standard Model quartic coupling.
// It extracts the finite invariants that any future bridge-layer Higgs potential
// must respect:
//
//   - four active real contact directions selected by the off-diagonal mixing;
//   - three protected unmixed contact directions;
//   - a positive order-parameter trace on the active contact sector;
//   - pair-degenerate contact mixing eigenvalues, indicating a natural complex
//     two-component organization candidate.
package higgspotential

import (
	"sync"

	"fmt"
	"math"

	gaugehiggs "github.com/bagherbal/asha-engine/pkg/gauge/higgs"
)

type DegeneracyCluster struct {
	Start        int
	End          int
	Multiplicity int
	Mean         float64
	MaxSpread    float64
}

type Analysis struct {
	Mixing gaugehiggs.Analysis

	ActiveContactDimension    int
	ProtectedContactDimension int
	ComplementResponseRank    int

	ActiveContactSpectrum  []float64
	DegeneracyClusters     []DegeneracyCluster
	PairDegeneracyResidual float64
	PairDegenerateSpectrum bool

	OrderParameterNormSquared float64 // τ = Tr(M_K) over the finite active contact sector.
	QuadraticTrace            float64 // Tr(M_K).
	QuarticTrace              float64 // Tr(M_K²).
	NormalizedQuarticShape    float64 // Tr(M_K²)/Tr(M_K)², scale-free.
	MeanActiveEigenvalue      float64
	SpectralAnisotropy        float64 // (λ_max-λ_min)/mean over active eigenvalues.

	MexicanHatKinematics bool // true when the finite spectrum has the correct 4+3 kinematic pattern.
}

var (
	higgsPotentialDefaultOnce  sync.Once
	higgsPotentialDefaultValue Analysis
	higgsPotentialDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	higgsPotentialDefaultOnce.Do(func() {
		higgsPotentialDefaultValue, higgsPotentialDefaultErr = buildHiggsPotentialDefaultUncached()
	})
	return higgsPotentialDefaultValue, higgsPotentialDefaultErr
}

func buildHiggsPotentialDefaultUncached() (Analysis, error) {
	m, err := gaugehiggs.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(m, 1e-8)
}

func Build(m gaugehiggs.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	active := positiveValues(m.VacuumMixingSpectrum, eps)
	if len(active) == 0 {
		return Analysis{}, fmt.Errorf("vacuum mixing spectrum has no positive active values")
	}

	quartic, err := traceSquare(m.VacuumMixingSpectrum)
	if err != nil {
		return Analysis{}, err
	}
	clusters := clusterDegeneracies(active, eps)
	pairResidual, pairOK := pairDegeneracy(active, eps)
	mean := mean(active)
	anisotropy := 0.0
	if mean > eps {
		anisotropy = (active[0] - active[len(active)-1]) / mean
	}
	normalizedQuartic := 0.0
	if math.Abs(m.VacuumMixingTrace) > eps {
		normalizedQuartic = quartic / (m.VacuumMixingTrace * m.VacuumMixingTrace)
	}
	mexicanHatKinematics := m.VacuumMixingRank == 4 && m.VacuumUnmixedDimension == 3 && pairOK && m.VacuumMixingTrace > eps

	return Analysis{
		Mixing:                    m,
		ActiveContactDimension:    m.VacuumMixingRank,
		ProtectedContactDimension: m.VacuumUnmixedDimension,
		ComplementResponseRank:    m.ComplementMixingRank,
		ActiveContactSpectrum:     active,
		DegeneracyClusters:        clusters,
		PairDegeneracyResidual:    pairResidual,
		PairDegenerateSpectrum:    pairOK,
		OrderParameterNormSquared: m.VacuumMixingTrace,
		QuadraticTrace:            m.VacuumMixingTrace,
		QuarticTrace:              quartic,
		NormalizedQuarticShape:    normalizedQuartic,
		MeanActiveEigenvalue:      mean,
		SpectralAnisotropy:        anisotropy,
		MexicanHatKinematics:      mexicanHatKinematics,
	}, nil
}

func positiveValues(values []float64, eps float64) []float64 {
	out := make([]float64, 0, len(values))
	for _, v := range values {
		if v > eps {
			out = append(out, v)
		}
	}
	return out
}

func traceSquare(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("cannot compute trace square of empty spectrum")
	}
	sum := 0.0
	for _, v := range values {
		sum += v * v
	}
	return sum, nil
}

func clusterDegeneracies(values []float64, eps float64) []DegeneracyCluster {
	if len(values) == 0 {
		return nil
	}
	clusters := make([]DegeneracyCluster, 0)
	start := 0
	for start < len(values) {
		end := start
		minV, maxV := values[start], values[start]
		for end+1 < len(values) && math.Abs(values[end+1]-values[start]) <= eps {
			end++
			if values[end] < minV {
				minV = values[end]
			}
			if values[end] > maxV {
				maxV = values[end]
			}
		}
		sum := 0.0
		for i := start; i <= end; i++ {
			sum += values[i]
		}
		clusters = append(clusters, DegeneracyCluster{
			Start:        start,
			End:          end,
			Multiplicity: end - start + 1,
			Mean:         sum / float64(end-start+1),
			MaxSpread:    maxV - minV,
		})
		start = end + 1
	}
	return clusters
}

func pairDegeneracy(values []float64, eps float64) (float64, bool) {
	if len(values) == 0 || len(values)%2 != 0 {
		return math.Inf(1), false
	}
	maxResidual := 0.0
	for i := 0; i < len(values); i += 2 {
		res := math.Abs(values[i] - values[i+1])
		if res > maxResidual {
			maxResidual = res
		}
	}
	return maxResidual, maxResidual < eps
}

func mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}
