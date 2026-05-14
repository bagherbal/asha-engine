// Package casimirkernel performs Gate 67: current-sector Casimir / propagator
// diagnostic.
//
// Gate 66 constructed genuine current-sector Casimir operators
//
//	C_A = sum_a T_a^T T_a
//
// on the one-generation 1+3 Pati-Salam flavor carrier.  Gate 67 asks the next
// question: can those Casimirs themselves become propagator denominators or
// exchange kernels?
//
// The answer is deliberately conservative.  The package exposes several
// candidate kernel families built from C_A — direct, trace-normalized, inverse
// nonzero, and inverse-trace-normalized — but marks all of them as diagnostics
// until a finite exchange action selects one.  A positive Casimir is
// representation data.  It is not automatically a propagator mass, kinetic
// denominator, or NJL kernel.
package casimirkernel

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/sectoroperators"
)

type SectorDiagnostic struct {
	Sector                  string
	GeneratorCount          int
	Rank                    int
	Nullity                 int
	Trace                   float64
	DirectDiagonal          []float64
	TraceNormalizedDiagonal []float64
	InverseNonZeroDiagonal  []float64
	InverseTrace            float64
	HasZeroMode             bool
	ColorBlindWeakBlind     bool
	SelectedAsPropagator    bool
	Interpretation          string
}

type Analysis struct {
	Previous sectoroperators.Analysis

	Diagnostics                    []SectorDiagnostic
	SectorCount                    int
	FlavorDimension                int
	TotalDirectTrace               float64
	TotalInverseTrace              float64
	AllCasimirDiagnosticsBuilt     bool
	DirectKernelFamilyExposed      bool
	InverseKernelFamilyExposed     bool
	TraceNormalizedFamilyExposed   bool
	AnySectorSelectedAsPropagator  bool
	FiniteActionSelectionDerived   bool
	PropagatorDenominatorsDerived  bool
	ExchangeKernelUpdated          bool
	AttractiveScalarChannelDerived bool
	UpDownSplittingDerived         bool
	CondensationClaimAllowed       bool
	HiddenObservedInputUsed        bool

	DominantDirectSector  string
	DominantInverseSector string
	ColorSectorZeroMode   bool
	AmbiguityStatement    string
	TruthStatement        string
	RecommendedNextGate   string
	RemainingUnknowns     []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := sectoroperators.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev sectoroperators.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	diagnostics := make([]SectorDiagnostic, 0, len(prev.Operators))
	totalDirect := 0.0
	totalInverse := 0.0
	dominantDirect := ""
	dominantDirectTrace := math.Inf(-1)
	dominantInverse := ""
	dominantInverseTrace := math.Inf(-1)
	colorZero := false

	for _, op := range prev.Operators {
		if len(op.DiagonalSpectrum) == 0 {
			return Analysis{}, fmt.Errorf("sector %s has empty diagonal spectrum", op.Sector)
		}
		direct := append([]float64(nil), op.DiagonalSpectrum...)
		trace := op.Trace
		if math.Abs(trace) < eps {
			return Analysis{}, fmt.Errorf("sector %s has zero trace", op.Sector)
		}
		traceNorm := make([]float64, len(direct))
		inverse := make([]float64, len(direct))
		inverseTrace := 0.0
		nullity := 0
		hasZero := false
		for i, v := range direct {
			traceNorm[i] = v / trace
			if v <= eps {
				nullity++
				hasZero = true
				inverse[i] = 0
				continue
			}
			inverse[i] = 1.0 / v
			inverseTrace += inverse[i]
		}
		if op.Sector == "color-su3" && hasZero {
			colorZero = true
		}
		if trace > dominantDirectTrace {
			dominantDirectTrace = trace
			dominantDirect = op.Sector
		}
		if inverseTrace > dominantInverseTrace {
			dominantInverseTrace = inverseTrace
			dominantInverse = op.Sector
		}
		totalDirect += trace
		totalInverse += inverseTrace

		diagnostics = append(diagnostics, SectorDiagnostic{
			Sector:                  op.Sector,
			GeneratorCount:          op.GeneratorCount,
			Rank:                    op.Rank,
			Nullity:                 nullity,
			Trace:                   trace,
			DirectDiagonal:          direct,
			TraceNormalizedDiagonal: traceNorm,
			InverseNonZeroDiagonal:  inverse,
			InverseTrace:            inverseTrace,
			HasZeroMode:             hasZero,
			ColorBlindWeakBlind:     true,
			SelectedAsPropagator:    false,
			Interpretation:          diagnosticInterpretation(op.Sector),
		})
	}

	truth := "Gate 67 exposes Casimir-built kernel families but does not select one. Direct C_A, trace-normalized C_A, and nonzero inverse C_A are useful diagnostics; none is a propagator denominator without a finite exchange-action theorem. The color sector correctly has a lepton zero mode, but these Casimirs still act only on the 1+3 flavor carrier and do not split up/down weak channels or select top condensation."
	amb := "direct Casimir weighting favors the color sector, inverse-nonzero weighting favors small-charge abelian data; because these diagnostics disagree, the finite action must select the propagator rule before G_hat can be updated"

	return Analysis{
		Previous:                       prev,
		Diagnostics:                    diagnostics,
		SectorCount:                    len(diagnostics),
		FlavorDimension:                4,
		TotalDirectTrace:               totalDirect,
		TotalInverseTrace:              totalInverse,
		AllCasimirDiagnosticsBuilt:     len(diagnostics) == 4,
		DirectKernelFamilyExposed:      true,
		InverseKernelFamilyExposed:     true,
		TraceNormalizedFamilyExposed:   true,
		AnySectorSelectedAsPropagator:  false,
		FiniteActionSelectionDerived:   false,
		PropagatorDenominatorsDerived:  false,
		ExchangeKernelUpdated:          false,
		AttractiveScalarChannelDerived: false,
		UpDownSplittingDerived:         false,
		CondensationClaimAllowed:       false,
		HiddenObservedInputUsed:        false,
		DominantDirectSector:           dominantDirect,
		DominantInverseSector:          dominantInverse,
		ColorSectorZeroMode:            colorZero,
		AmbiguityStatement:             amb,
		TruthStatement:                 truth,
		RecommendedNextGate:            "Gate 68 — Finite Exchange-Action Selection Principle",
		RemainingUnknowns: []string{
			"U-20D2B3-CURRENT-KINETIC-OPERATOR: decide whether direct Casimir, inverse Casimir, Laplacian spectrum, or another operator controls exchange",
			"U-20D2B4-ACTION-SELECTION: derive the finite action that chooses a propagator rule rather than choosing by diagnostic convenience",
			"U-20D3-RELATIVE-COUPLINGS: derive sector coupling weights after the propagator rule is selected",
			"U-20D4-UP-DOWN-SPLITTING: sector Casimirs live on lepton/color flavor and still do not distinguish up from down",
			"U-20D6-CRITICAL-REGULATOR: derive NJL criticality only after a real exchange kernel exists",
		},
	}, nil
}

func diagnosticInterpretation(sector string) string {
	switch sector {
	case "central":
		return "uniform flavor Casimir; direct and inverse kernels are both completely degenerate"
	case "b-minus-l":
		return "charge-square Casimir; inverse diagnostic heavily weights color seeds because |B-L|=1/3 there"
	case "color-su3":
		return "color Casimir; annihilates the lepton seed and acts uniformly on colors"
	case "leptoquark":
		return "lepton-color exchange Casimir; larger on lepton than on each color seed"
	default:
		return "sector Casimir diagnostic"
	}
}

func FormatDiagnostics(xs []SectorDiagnostic) string {
	ys := append([]SectorDiagnostic(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Sector < ys[j].Sector })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(n=%d, rank=%d, nullity=%d, TrC=%.10f, C=%s, C/tr=%s, C+diag=%s, TrC+=%.10f, selected=%v)", x.Sector, x.GeneratorCount, x.Rank, x.Nullity, x.Trace, formatFloatList(x.DirectDiagonal), formatFloatList(x.TraceNormalizedDiagonal), formatFloatList(x.InverseNonZeroDiagonal), x.InverseTrace, x.SelectedAsPropagator))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}

func formatFloatList(xs []float64) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%.10f", x))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
