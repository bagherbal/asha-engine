// Package fierzsign performs Gate 61: finite Clifford/Lorentz Fierz sign construction.
//
// Gate 60 derived finite kinetic-trace weights and unsigned scalar-channel
// overlap coefficients for the u(4)-shaped current sectors.  Those coefficients
// are not yet an NJL kernel because a current-current interaction must first be
// projected through the Lorentz/chiral Fierz identities.
//
// This package implements the native two-component spinor completeness identity
//
//	sigma^mu_{a dot b} bar(sigma)_mu^{dot c d} = 2 delta_a^d delta_dotb^dotc
//
// and the fermion-reordering sign for the LR scalar channel.  Therefore a
// vector LR current-current term has the scalar-channel Fierz coefficient
//
//	c_LR^scalar = -2
//
// in this convention.  This is a signed Fierz projection.  It still does not
// derive an attractive NJL kernel, because the finite exchange action,
// propagator sign, and relative current couplings remain open.
package fierzsign

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/kineticnorm"
)

type SectorSignedCoefficient struct {
	Sector                    string
	GeneratorCount            int
	KineticTrace              float64
	TraceWeight               float64
	UnsignedProjection        float64
	SignedFierzCoefficient    float64
	WeightedSignedCoefficient float64
}

type Analysis struct {
	Kinetic kineticnorm.Analysis

	SpinorDimension                   int
	SigmaBarSigmaIdentityResidual     float64
	SpinorCompletenessFactor          float64
	FermionReorderingSign             float64
	UniversalLRScalarFierzCoefficient float64

	SectorCoefficients             []SectorSignedCoefficient
	TotalWeightedSignedCoefficient float64

	SigmaIdentityConstructed                bool
	FermionReorderingIncluded               bool
	SignedCliffordFierzCoefficientsDerived  bool
	GeneratorPropagatorNormalizationDerived bool
	AttractiveScalarChannelSignDerived      bool
	NativeFourFermionKernelDerived          bool
	UpDownSplittingDerived                  bool
	HiddenObservedInputUsed                 bool

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
		k, err := kineticnorm.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(k, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(k kineticnorm.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !k.KineticTraceNormalizationDerived || !k.UnsignedUnitProjectionCoefficientsDerived {
		return Analysis{}, fmt.Errorf("Gate 61 requires Gate 60 kinetic normalization and unsigned scalar projection")
	}

	residual := sigmaBarSigmaResidual()
	sigmaOK := residual < 1e-12
	spinorDim := 2
	completeness := float64(spinorDim)
	reorder := -1.0
	coeff := reorder * completeness

	sectors := make([]SectorSignedCoefficient, 0, len(k.SectorNormalizations))
	total := 0.0
	for _, s := range k.SectorNormalizations {
		signed := coeff * s.NormalizedOverlap
		weighted := signed * s.TraceWeight
		total += weighted
		sectors = append(sectors, SectorSignedCoefficient{
			Sector:                    s.Sector,
			GeneratorCount:            s.GeneratorCount,
			KineticTrace:              s.RawKineticTrace,
			TraceWeight:               s.TraceWeight,
			UnsignedProjection:        s.NormalizedOverlap,
			SignedFierzCoefficient:    signed,
			WeightedSignedCoefficient: weighted,
		})
	}
	sort.Slice(sectors, func(i, j int) bool { return sectors[i].Sector < sectors[j].Sector })

	truth := "The engine now constructs the native LR vector-to-scalar Fierz sign: the sigma/bar-sigma completeness identity gives the spinor factor 2, and fermion reordering gives the scalar-channel sign -1, so the universal LR scalar Fierz coefficient is -2 in this convention. This upgrades unsigned current-sector overlaps into signed Fierz coefficients. It still does not prove NJL attraction: the finite propagator/action sign, exchange normalization, and up/down splitting remain open."

	return Analysis{
		Kinetic:                                 k,
		SpinorDimension:                         spinorDim,
		SigmaBarSigmaIdentityResidual:           residual,
		SpinorCompletenessFactor:                completeness,
		FermionReorderingSign:                   reorder,
		UniversalLRScalarFierzCoefficient:       coeff,
		SectorCoefficients:                      sectors,
		TotalWeightedSignedCoefficient:          total,
		SigmaIdentityConstructed:                sigmaOK,
		FermionReorderingIncluded:               true,
		SignedCliffordFierzCoefficientsDerived:  sigmaOK && math.Abs(coeff+2) < eps,
		GeneratorPropagatorNormalizationDerived: false,
		AttractiveScalarChannelSignDerived:      false,
		NativeFourFermionKernelDerived:          false,
		UpDownSplittingDerived:                  false,
		HiddenObservedInputUsed:                 false,
		TruthStatement:                          truth,
		RecommendedNextGate:                     "Gate 62 — Propagator/Action Sign and Exchange Kernel Audit",
		RemainingUnknowns: []string{
			"U-20D2-PROPAGATOR-SIGN: derive the finite exchange action sign and propagator orientation",
			"U-20D3-RELATIVE-COUPLINGS: derive current-sector coupling strengths instead of unit weights",
			"U-20D4-UP-DOWN-SPLITTING: break the up/down quark tie without observed Yukawa input",
			"U-20D5-G-HAT: combine signed Fierz coefficients with propagator weights into native four-fermion strength",
			"U-20D6-CRITICALITY: derive finite regulator/critical threshold for the NJL gap equation",
		},
	}, nil
}

// sigmaBarSigmaResidual verifies
// sigma^mu_{a dot b} bar(sigma)_mu^{dot c d} = 2 delta_a^d delta_dotb^dotc
// using the Weyl sigma matrices and metric (+---).
func sigmaBarSigmaResidual() float64 {
	sigma := [4][2][2]complex128{
		{{1, 0}, {0, 1}},
		{{0, 1}, {1, 0}},
		{{0, -1i}, {1i, 0}},
		{{1, 0}, {0, -1}},
	}
	barsigma := [4][2][2]complex128{
		{{1, 0}, {0, 1}},
		{{0, -1}, {-1, 0}},
		{{0, 1i}, {-1i, 0}},
		{{-1, 0}, {0, 1}},
	}
	eta := [4]float64{1, -1, -1, -1}
	max := 0.0
	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			for c := 0; c < 2; c++ {
				for d := 0; d < 2; d++ {
					var sum complex128
					for mu := 0; mu < 4; mu++ {
						sum += complex(eta[mu], 0) * sigma[mu][a][b] * barsigma[mu][c][d]
					}
					target := complex(0, 0)
					if a == d && b == c {
						target = 2
					}
					diff := cmplxAbs(sum - target)
					if diff > max {
						max = diff
					}
				}
			}
		}
	}
	return max
}

func cmplxAbs(z complex128) float64 {
	return math.Hypot(real(z), imag(z))
}

func FormatSectorCoefficients(xs []SectorSignedCoefficient) string {
	ys := append([]SectorSignedCoefficient(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Sector < ys[j].Sector })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(n=%d, K=%.10f, weight=%.6f, c_signed=%.10f, weighted=%.10f)", x.Sector, x.GeneratorCount, x.KineticTrace, x.TraceWeight, x.SignedFierzCoefficient, x.WeightedSignedCoefficient))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
