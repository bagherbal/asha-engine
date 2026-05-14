// Package action builds the first conservative representation-action bridge
// from the finite Higgs/contact mixing spectrum to the 16-state Witt/Fock basis.
//
// This package does not claim a Standard Model Yukawa matrix or physical fermion
// masses. It implements the standard finite operation that is justified once a
// four-dimensional one-particle response spectrum is available: second-quantize
// that one-particle operator over the exterior/Fock occupation basis.
//
// If λ_μ are the four active contact-Higgs eigenvalues, the induced Fock response
// on an occupation state n=(n_0,n_1,n_2,n_3) is
//
//	H_F |n⟩ = (Σ_μ λ_μ n_μ) |n⟩.
//
// The construction keeps the bridge honest: it verifies a representation action
// exists at the spectral/number-operator level, while leaving the canonical
// eigenvector map K₇ → Fock modes and the Yukawa texture as open theorems.
package action

import (
	"sync"

	"fmt"
	"math"
	"sort"

	"github.com/bagherbal/asha-engine/pkg/dynamics/higgspotential"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

type Analysis struct {
	Bridge matter.FockContactBridge

	OneParticleWeights []float64
	Operator           linear.Matrix // 16x16 diagonal second-quantized response operator.
	Spectrum           []float64

	VacuumIndex    int
	VacuumResponse float64

	OneParticleResponses []float64
	PairResidual         float64
	PairDegenerate       bool

	Trace               float64
	ExpectedTrace       float64
	TraceResidual       float64
	MaxResponse         float64
	ExpectedMaxResponse float64
	Rank                int

	EvenParityTrace     float64
	OddParityTrace      float64
	ParityTraceResidual float64

	CanonicalEigenvectorEmbeddingConstructed bool
	YukawaTextureConstructed                 bool
}

var (
	matterActionDefaultOnce  sync.Once
	matterActionDefaultValue Analysis
	matterActionDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	matterActionDefaultOnce.Do(func() {
		matterActionDefaultValue, matterActionDefaultErr = buildMatterActionDefaultUncached()
	})
	return matterActionDefaultValue, matterActionDefaultErr
}

func buildMatterActionDefaultUncached() (Analysis, error) {
	b, err := matter.BuildDefaultFockContactBridge()
	if err != nil {
		return Analysis{}, err
	}
	return Build(b, 1e-10)
}

func Build(b matter.FockContactBridge, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if b.FockModeCount != b.ActiveHiggsDirections {
		return Analysis{}, fmt.Errorf("cannot second-quantize: Fock modes=%d active Higgs directions=%d", b.FockModeCount, b.ActiveHiggsDirections)
	}
	weights := append([]float64(nil), b.Potential.ActiveContactSpectrum...)
	if len(weights) != b.Fock.ModeCount() {
		return Analysis{}, fmt.Errorf("active spectrum has %d values, expected %d", len(weights), b.Fock.ModeCount())
	}
	// Keep a stable descending spectral convention. This is a bridge convention,
	// not a physical hypercharge or flavor assignment.
	sort.Sort(sort.Reverse(sort.Float64Slice(weights)))

	responses := make([]float64, b.Fock.StateCount())
	op := linear.NewMatrix(b.Fock.StateCount(), b.Fock.StateCount())
	vacuumIndex := -1
	vacuumResponse := 0.0
	rank := 0
	trace := 0.0
	maxResponse := math.Inf(-1)
	evenTrace := 0.0
	oddTrace := 0.0

	for i, state := range b.Fock.States {
		response := occupationResponse(state, weights)
		responses[i] = response
		op.Set(i, i, response)
		trace += response
		if response > eps {
			rank++
		}
		if response > maxResponse {
			maxResponse = response
		}
		if state.ExcitationNumber()%2 == 0 {
			evenTrace += response
		} else {
			oddTrace += response
		}
		if state.IsVacuum() {
			vacuumIndex = i
			vacuumResponse = response
		}
	}
	if vacuumIndex < 0 {
		return Analysis{}, fmt.Errorf("Fock basis has no vacuum state")
	}

	oneParticle := make([]float64, 0, b.Fock.ModeCount())
	for _, state := range b.Fock.States {
		if state.ExcitationNumber() == 1 {
			oneParticle = append(oneParticle, occupationResponse(state, weights))
		}
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(oneParticle)))
	pairResidual, pairOK := pairResidual(oneParticle, eps)

	sumWeights := 0.0
	for _, w := range weights {
		sumWeights += w
	}
	expectedTrace := math.Pow(2, float64(len(weights)-1)) * sumWeights
	expectedMax := sumWeights

	return Analysis{
		Bridge:                                   b,
		OneParticleWeights:                       weights,
		Operator:                                 op,
		Spectrum:                                 responses,
		VacuumIndex:                              vacuumIndex,
		VacuumResponse:                           vacuumResponse,
		OneParticleResponses:                     oneParticle,
		PairResidual:                             pairResidual,
		PairDegenerate:                           pairOK,
		Trace:                                    trace,
		ExpectedTrace:                            expectedTrace,
		TraceResidual:                            math.Abs(trace - expectedTrace),
		MaxResponse:                              maxResponse,
		ExpectedMaxResponse:                      expectedMax,
		Rank:                                     rank,
		EvenParityTrace:                          evenTrace,
		OddParityTrace:                           oddTrace,
		ParityTraceResidual:                      math.Abs(evenTrace - oddTrace),
		CanonicalEigenvectorEmbeddingConstructed: false,
		YukawaTextureConstructed:                 false,
	}, nil
}

func occupationResponse(state spinor.FockState, weights []float64) float64 {
	sum := 0.0
	for i, occupied := range state.Occupation {
		if occupied {
			sum += weights[i]
		}
	}
	return sum
}

func pairResidual(values []float64, eps float64) (float64, bool) {
	if len(values) == 0 || len(values)%2 != 0 {
		return math.Inf(1), false
	}
	maxResidual := 0.0
	for i := 0; i < len(values); i += 2 {
		residual := math.Abs(values[i] - values[i+1])
		if residual > maxResidual {
			maxResidual = residual
		}
	}
	return maxResidual, maxResidual < eps
}

func FormatFloatSlice(values []float64) string {
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10g", v)
	}
	return out + "]"
}

func TopSpectrum(values []float64, n int) []float64 {
	out := append([]float64(nil), values...)
	sort.Sort(sort.Reverse(sort.Float64Slice(out)))
	if n > len(out) {
		n = len(out)
	}
	return append([]float64(nil), out[:n]...)
}

func FiniteHiggsPotential(p higgspotential.Analysis) []float64 {
	// Convenience hook for future bridge packages: the finite active spectrum is
	// the one-particle datum currently available for second quantization.
	return append([]float64(nil), p.ActiveContactSpectrum...)
}
