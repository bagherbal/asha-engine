// Package couplingnorm audits whether the dimensionless topological action seal
// can normalize a gauge coupling.
//
// The finite contact index gives the dimensionless action seal S_top=8*pi^2 I_BG.
// In continuum Yang--Mills normalization a unit instanton has action
// S_YM=8*pi^2/g^2. That resemblance is important, but it is not by itself a
// derivation of a physical coupling: one must also derive the continuum index
// bridge and the trace/kinetic normalization that identifies the finite inner
// product with the continuum gauge kinetic term.
//
// This package therefore exposes the coupling family and the unit-trace
// diagnostic while rejecting any claim that alpha, the electromagnetic coupling,
// or an RG boundary condition has been derived.
package couplingnorm

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/actionscale"
)

type Analysis struct {
	Action actionscale.Analysis

	ContactIndex          float64
	TopologicalActionSeal float64
	InstantonWeight       float64

	ContinuumFormula string
	FiniteFormula    string
	CouplingFamily   string

	// UnitTrace* are diagnostics under the convention that the finite kinetic
	// trace normalization equals the continuum one. They are not physical
	// predictions until that convention is derived.
	UnitTraceNormalization       float64
	UnitTraceGaugeCouplingSq     float64
	UnitTraceGaugeCoupling       float64
	UnitTraceAlpha               float64
	UnitTraceInverseAlpha        float64
	UnitTraceElectromagneticLike bool

	ContinuumIndexBridgeDerived bool
	TraceNormalizationDerived   bool
	GaugeCouplingDerived        bool
	RGBoundaryDerived           bool
	FineStructureDerived        bool
	DimensionfulScaleDerived    bool
	HiddenObservedCouplingUsed  bool

	TruthStatement    string
	RemainingUnknowns []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		a, err := actionscale.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(a, 1e-12)
	})
	return defaultValue, defaultErr
}

func Build(a actionscale.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-12
	}
	if a.ContactIndex <= eps {
		return Analysis{}, fmt.Errorf("contact index must be positive")
	}
	if a.TopologicalActionSeal <= eps {
		return Analysis{}, fmt.Errorf("topological action seal must be positive")
	}

	// Unit-trace convention diagnostic only: S_fin=8*pi^2 I and
	// S_YM=8*pi^2 I/g^2 would give g^2=1. This is a convention until the
	// finite trace normalization is derived.
	unitTrace := 1.0
	g2 := 1.0 / unitTrace
	g := math.Sqrt(g2)
	alpha := g2 / (4 * math.Pi)
	invAlpha := 1 / alpha

	return Analysis{
		Action:                       a,
		ContactIndex:                 a.ContactIndex,
		TopologicalActionSeal:        a.TopologicalActionSeal,
		InstantonWeight:              a.ActionWeight,
		ContinuumFormula:             "S_YM(k=I_BG)=8π² I_BG / g²",
		FiniteFormula:                "S_top=8π² I_BG",
		CouplingFamily:               "if S_fin=κ·8π²I_BG is equated to S_YM, then g²=1/κ; κ is the missing trace/kinetic normalization",
		UnitTraceNormalization:       unitTrace,
		UnitTraceGaugeCouplingSq:     g2,
		UnitTraceGaugeCoupling:       g,
		UnitTraceAlpha:               alpha,
		UnitTraceInverseAlpha:        invAlpha,
		UnitTraceElectromagneticLike: false,
		ContinuumIndexBridgeDerived:  false,
		TraceNormalizationDerived:    false,
		GaugeCouplingDerived:         false,
		RGBoundaryDerived:            false,
		FineStructureDerived:         false,
		DimensionfulScaleDerived:     false,
		HiddenObservedCouplingUsed:   false,
		TruthStatement:               "The topological action seal can define a dimensionless coupling-normalization problem, not a physical gauge coupling by itself. Under the unit-trace convention it gives g²=1 and alpha^{-1}=4π, which is a normalization diagnostic, not the electromagnetic fine-structure constant.",
		RemainingUnknowns: []string{
			"U-20A-CONTINUUM-INDEX-BRIDGE: prove finite I_BG is carried to the continuum topological charge entering the Yang--Mills action",
			"U-20B-TRACE-KINETIC-NORMALIZATION: derive the finite-to-continuum trace normalization κ rather than choosing it by convention",
			"U-20C-RG-BOUNDARY: derive an RG boundary condition and identify which gauge factor it applies to",
			"U-20D-ELECTROMAGNETIC-PROJECTION: derive the electroweak mixing/projection before comparing to alpha_em",
			"U-20E-NO-COUPLING-FIT: reject choosing κ to force alpha^{-1}=137.036",
		},
	}, nil
}
