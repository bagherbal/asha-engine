// Package exchangekernel performs Gate 62: propagator/action sign and exchange-kernel audit.
//
// Gate 61 derived the native LR scalar Fierz sign c_LR^scalar = -2 from the
// sigma/bar-sigma completeness identity plus fermion reordering.  A signed
// Fierz coefficient is still not a physical NJL kernel: a current-current
// exchange also needs a finite action orientation, propagator normalization,
// and relative current-sector couplings.
//
// This package keeps that distinction explicit.  It computes the conditional
// scalar-kernel values under the two possible exchange orientations and exposes
// the one that would be attractive in the common NJL convention
//
//	L_NJL = +G (PsiBar_R Psi_L)(PsiBar_L Psi_R),   G > 0.
//
// The engine does not claim that this orientation is derived.  It only says:
// if the finite vector-exchange action has the attractive orientation, the
// signed Fierz data would give a positive unit-normalized scalar kernel.  The
// action sign, propagator weights, up/down splitting, and regulator criticality
// remain open.
package exchangekernel

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fierzsign"
)

type SectorExchangeDiagnostic struct {
	Sector                    string
	GeneratorCount            int
	TraceWeight               float64
	SignedFierzCoefficient    float64
	WeightedSignedCoefficient float64

	ScalarKernelIfExchangePlus  float64
	ScalarKernelIfExchangeMinus float64
	AttractiveUnitContribution  float64
}

type Analysis struct {
	Fierz fierzsign.Analysis

	NJLAttractiveConvention string
	CurrentCurrentTemplate  string

	ExchangeSignDerived             bool
	PropagatorMagnitudeDerived      bool
	RelativeCurrentCouplingsDerived bool

	UnitExchangePlusTotal           float64
	UnitExchangeMinusTotal          float64
	ConditionalAttractiveUnitKernel float64
	ConditionalRepulsiveUnitKernel  float64

	SectorDiagnostics []SectorExchangeDiagnostic

	SignedFierzAvailable                 bool
	ConditionalAttractiveBranchAvailable bool
	AttractiveScalarChannelSignDerived   bool
	NativeFourFermionKernelDerived       bool
	UpDownSplittingDerived               bool
	RegulatorCriticalityDerived          bool
	CondensationClaimAllowed             bool
	HiddenObservedInputUsed              bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		f, err := fierzsign.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(f, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(f fierzsign.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !f.SignedCliffordFierzCoefficientsDerived {
		return Analysis{}, fmt.Errorf("Gate 62 requires Gate 61 signed Clifford/Lorentz Fierz coefficients")
	}

	sectors := make([]SectorExchangeDiagnostic, 0, len(f.SectorCoefficients))
	plusTotal := 0.0
	minusTotal := 0.0
	for _, s := range f.SectorCoefficients {
		// exchangePlus corresponds to an effective +J^2 orientation.
		// exchangeMinus corresponds to an effective -J^2 orientation.
		// Since the scalar Fierz coefficient is negative in Gate 61, the
		// -J^2 orientation yields a positive scalar coefficient in the
		// chosen NJL convention.
		plus := s.WeightedSignedCoefficient
		minus := -s.WeightedSignedCoefficient
		plusTotal += plus
		minusTotal += minus
		sectors = append(sectors, SectorExchangeDiagnostic{
			Sector:                      s.Sector,
			GeneratorCount:              s.GeneratorCount,
			TraceWeight:                 s.TraceWeight,
			SignedFierzCoefficient:      s.SignedFierzCoefficient,
			WeightedSignedCoefficient:   s.WeightedSignedCoefficient,
			ScalarKernelIfExchangePlus:  plus,
			ScalarKernelIfExchangeMinus: minus,
			AttractiveUnitContribution:  math.Max(plus, minus),
		})
	}
	sort.Slice(sectors, func(i, j int) bool { return sectors[i].Sector < sectors[j].Sector })

	conditionalAttractive := math.Max(plusTotal, minusTotal)
	conditionalRepulsive := math.Min(plusTotal, minusTotal)
	conditionalBranch := conditionalAttractive > eps && conditionalRepulsive < -eps

	truth := "Gate 62 converts the signed Fierz result into a conditional exchange-kernel audit. With the common NJL convention L_NJL=+G(Ψ̄_RΨ_L)(Ψ̄_LΨ_R), the Gate 61 coefficient c_LR=-2 becomes attractive only if the finite current exchange has the -J·J orientation. Under unit exchange this would give G_hat_unit=2, but the engine does not claim that orientation, propagator magnitude, or relative current couplings are derived. Therefore attraction and condensation remain bridge-open."

	return Analysis{
		Fierz:                                f,
		NJLAttractiveConvention:              "L_NJL=+G(Ψ̄_RΨ_L)(Ψ̄_LΨ_R), G>0",
		CurrentCurrentTemplate:               "L_eff = η_A ρ_A J_A·J_A; η_A, ρ_A are finite action/propagator data and remain open",
		ExchangeSignDerived:                  false,
		PropagatorMagnitudeDerived:           false,
		RelativeCurrentCouplingsDerived:      false,
		UnitExchangePlusTotal:                plusTotal,
		UnitExchangeMinusTotal:               minusTotal,
		ConditionalAttractiveUnitKernel:      conditionalAttractive,
		ConditionalRepulsiveUnitKernel:       conditionalRepulsive,
		SectorDiagnostics:                    sectors,
		SignedFierzAvailable:                 true,
		ConditionalAttractiveBranchAvailable: conditionalBranch,
		AttractiveScalarChannelSignDerived:   false,
		NativeFourFermionKernelDerived:       false,
		UpDownSplittingDerived:               false,
		RegulatorCriticalityDerived:          false,
		CondensationClaimAllowed:             false,
		HiddenObservedInputUsed:              false,
		TruthStatement:                       truth,
		RecommendedNextGate:                  "Gate 63 — Finite Exchange Action / Propagator Normalization Search",
		RemainingUnknowns: []string{
			"U-20D2A-EXCHANGE-ACTION-SIGN: derive whether finite current exchange contributes +J·J or -J·J to the scalar channel",
			"U-20D2B-PROPAGATOR-MAGNITUDE: derive current-sector exchange weights ρ_A from kinetic/action normalization",
			"U-20D3-RELATIVE-COUPLINGS: derive relative current-sector couplings instead of unit weights",
			"U-20D4-UP-DOWN-SPLITTING: derive a native operator that distinguishes top-like up from bottom-like down",
			"U-20D6-CRITICAL-REGULATOR: derive C_reg for the finite NJL gap equation",
		},
	}, nil
}

func FormatSectorDiagnostics(xs []SectorExchangeDiagnostic) string {
	ys := append([]SectorExchangeDiagnostic(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Sector < ys[j].Sector })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(n=%d, weight=%.6f, c=%.10f, +J²=%.10f, -J²=%.10f)", x.Sector, x.GeneratorCount, x.TraceWeight, x.SignedFierzCoefficient, x.ScalarKernelIfExchangePlus, x.ScalarKernelIfExchangeMinus))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
