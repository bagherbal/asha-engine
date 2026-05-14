// Package kineticnorm performs Gate 60: finite generator kinetic-normalization
// and signed-Fierz coefficient audit.
//
// Gate 59 proved that the u(4)-shaped current generators act compatibly on the
// finite scalar LR projector and produce computable representation-overlap
// diagnostics.  This package takes the next step that can be done without
// importing continuum Fierz tables: it derives finite trace-normalization
// weights for the current sectors and normalizes the unsigned scalar-channel
// overlap coefficients against those kinetic traces.
//
// The result is deliberately strict.  Finite kinetic trace normalization is
// available.  However, this does not yet determine the Clifford/Lorentz Fierz
// signs, propagator/action signs, or attractive NJL scalar-channel coefficients.
// Therefore the native four-fermion kernel remains open.
package kineticnorm

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/currentprojection"
)

type SectorNormalization struct {
	Sector         string
	GeneratorCount int

	RawKineticTrace   float64
	RawOverlapTrace   float64
	NormalizedOverlap float64
	TraceWeight       float64

	UnitNormalized         bool
	SignedCoefficientKnown bool
	SignedCoefficient      float64
}

type Analysis struct {
	Current currentprojection.Analysis

	SectorNormalizations   []SectorNormalization
	TotalKineticTrace      float64
	MaxUnitOverlapResidual float64

	KineticTraceNormalizationDerived          bool
	UnsignedUnitProjectionCoefficientsDerived bool
	SignedCliffordFierzCoefficientsDerived    bool
	GeneratorPropagatorNormalizationDerived   bool
	AttractiveScalarChannelSignDerived        bool
	NativeFourFermionKernelDerived            bool
	UpDownSplittingDerived                    bool
	HiddenObservedInputUsed                   bool

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
		c, err := currentprojection.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(c, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(c currentprojection.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !c.UnsignedScalarProjectionCoefficientsKnown || len(c.SectorCoefficients) == 0 {
		return Analysis{}, fmt.Errorf("Gate 60 requires Gate 59 unsigned scalar-projection coefficients")
	}

	sectors := make([]SectorNormalization, 0, len(c.SectorCoefficients))
	totalTrace := 0.0
	maxResidual := 0.0
	allPositive := true
	allUnit := true

	for _, sc := range c.SectorCoefficients {
		if sc.RightNorm <= eps {
			allPositive = false
		}
		totalTrace += sc.RightNorm
	}
	if totalTrace <= eps {
		allPositive = false
	}

	for _, sc := range c.SectorCoefficients {
		residual := math.Abs(sc.NormalizedOverlap - 1.0)
		if residual > maxResidual {
			maxResidual = residual
		}
		if residual > 1e-8 {
			allUnit = false
		}
		weight := 0.0
		if totalTrace > eps {
			weight = sc.RightNorm / totalTrace
		}
		sectors = append(sectors, SectorNormalization{
			Sector:                 sc.Sector,
			GeneratorCount:         sc.GeneratorCount,
			RawKineticTrace:        sc.RightNorm,
			RawOverlapTrace:        sc.OverlapTrace,
			NormalizedOverlap:      sc.NormalizedOverlap,
			TraceWeight:            weight,
			UnitNormalized:         residual < 1e-8,
			SignedCoefficientKnown: false,
			SignedCoefficient:      0,
		})
	}

	sort.Slice(sectors, func(i, j int) bool { return sectors[i].Sector < sectors[j].Sector })

	truth := "Finite trace normalization of the current-sector generators is now available: the engine can normalize representation-overlap coefficients by their own kinetic traces.  The normalized scalar LR projection coefficients remain unit overlaps in this finite incidence representation.  This is still not a signed Fierz/NJL kernel: Clifford/Lorentz signs, propagator signs, and current kinetic/action normalization remain open."

	return Analysis{
		Current:                                   c,
		SectorNormalizations:                      sectors,
		TotalKineticTrace:                         totalTrace,
		MaxUnitOverlapResidual:                    maxResidual,
		KineticTraceNormalizationDerived:          allPositive,
		UnsignedUnitProjectionCoefficientsDerived: allUnit,
		SignedCliffordFierzCoefficientsDerived:    false,
		GeneratorPropagatorNormalizationDerived:   false,
		AttractiveScalarChannelSignDerived:        false,
		NativeFourFermionKernelDerived:            false,
		UpDownSplittingDerived:                    false,
		HiddenObservedInputUsed:                   false,
		TruthStatement:                            truth,
		RecommendedNextGate:                       "Gate 61 — Clifford/Lorentz Fierz Sign Construction",
		RemainingUnknowns: []string{
			"U-20D1B-CLIFFORD-TRACE-RULES: implement native gamma/pseudoscalar trace identities for LR scalar bilinears",
			"U-20D1D-SIGNED-FIERZ: compute signed current-current to scalar-channel coefficients",
			"U-20D2-PROPAGATOR-SIGN: derive the finite action/propagator sign for current exchange",
			"U-20D4-UP-DOWN-SPLITTING: break the up/down quark tie without observed Yukawa input",
			"U-20D5-G-HAT: combine signed coefficients and kinetic/propagator weights into native four-fermion strength",
		},
	}, nil
}

func FormatSectorNormalizations(xs []SectorNormalization) string {
	ys := append([]SectorNormalization(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Sector < ys[j].Sector })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(n=%d, K=%.10f, weight=%.6f, c_unsigned=%.10f, signed=open)", x.Sector, x.GeneratorCount, x.RawKineticTrace, x.TraceWeight, x.NormalizedOverlap))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
