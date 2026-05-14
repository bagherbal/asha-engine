// Package abeliancoupling implements Gate 81: abelian coupling normalization
// from the diagonal anomaly-constrained Hessian audit.
//
// Gates 75-80 established:
//
//  1. The charge-level hypercharge direction is selected as
//     Y = T3_R + (B-L)/2, with central u(1) rejected as a hypercharge
//     component.
//  2. The available finite U(1) trace-Gram diagnostics are diagonal:
//     central: Tr(I^2)=16, B-L: Tr((B-L)^2)=16/3, contact-u1:
//     Tr(T_phi^2)=1.
//  3. All known off-diagonal kinetic-mixing sources cancel, and that
//     cancellation is anomaly-shadow supported.
//
// This gate asks whether the surviving diagonal trace-Gram data is enough to
// normalize physical U(1) couplings.  The answer is conservative: it is enough
// to define representation-metric and canonical-generator diagnostics, but not
// enough to derive physical gauge couplings.  A finite action must still select
// the diagonal trace-Gram matrix as the gauge kinetic Hessian and must provide
// RG boundary/matching data before alpha_em can be claimed.
package abeliancoupling

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/anomalykinetic"
	"github.com/bagherbal/asha-engine/pkg/bridge/ewprojection"
)

type CouplingDiagnostic struct {
	Name               string
	TraceNormSquared   float64
	TraceNorm          float64
	InverseNormSquared float64
	Role               string
}

type HyperchargeDiagnostic struct {
	MatterBLCoefficient       float64
	MatterBLNormSquared       float64
	MatterBLContribution      float64
	ScalarContactNormSquared  float64
	ScalarContactContribution float64
	CombinedBridgeNorm        float64
	ChargeTableKY             float64
	BoundarySin2              float64
}

type Analysis struct {
	AnomalyKinetic anomalykinetic.Analysis
	EWProjection   ewprojection.Analysis

	Fields      []CouplingDiagnostic
	Hypercharge HyperchargeDiagnostic

	DiagonalTraceGramSelectedAsRepresentationMetric bool
	CanonicalGeneratorDiagnosticsDerived            bool
	DiagonalTraceGramSelectedAsGaugeKineticHessian  bool
	PhysicalGaugeCouplingsDerived                   bool
	FineStructureDerived                            bool
	RGBoundaryScaleDerived                          bool
	HiddenObservedInputUsed                         bool

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
		ak, err := anomalykinetic.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ew, err := ewprojection.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ak, ew, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(ak anomalykinetic.Analysis, ew ewprojection.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !ak.DiagonalPositive || !ak.AllKnownOffDiagonalSourcesCancel {
		return Analysis{}, fmt.Errorf("Gate 81 requires Gate 80 diagonal positive/no-source U(1) diagnostic")
	}
	if math.Abs(ew.HyperchargeNormalizationKY-5.0/3.0) > 1e-8 {
		return Analysis{}, fmt.Errorf("Gate 81 expected inherited k_Y=5/3, got %.12f", ew.HyperchargeNormalizationKY)
	}

	central := ak.Kinetic.Central.Trace2
	bl := ak.Kinetic.BMinusL.Trace2
	contact := ak.Kinetic.ContactU1.Trace2
	fields := []CouplingDiagnostic{
		makeDiagnostic("central u(1)", central, "universal finite current; rejected as a hypercharge component but still has a representation norm"),
		makeDiagnostic("B-L", bl, "matter-side abelian generator entering Y=T3_R+(B-L)/2"),
		makeDiagnostic("contact-u1 / T_phi", contact, "scalar/contact abelian generator with spectrum (+1/2,+1/2,-1/2,-1/2)"),
	}

	// Charge-level hypercharge uses the matter coefficient 1/2 for B-L.  The
	// scalar/contact charge is represented separately on the Higgs/contact factor.
	// The combined norm below is only a bridge diagnostic: the full one-generation
	// charge-table normalization k_Y=5/3 is inherited from Gate 41, not replaced by
	// this two-carrier representation metric.
	blCoeff := 0.5
	matterBLContribution := blCoeff * blCoeff * bl
	scalarContribution := contact
	combined := matterBLContribution + scalarContribution

	truth := "Gate 81 shows that the surviving diagonal U(1) trace-Gram data gives canonical representation-metric diagnostics for central u(1), B-L, and contact-u1. It does not by itself become the physical gauge kinetic Hessian. Therefore charge-level hypercharge and k_Y=5/3 remain valid, but g_Y, alpha_em, and low-energy coupling values remain open until a finite action selects the kinetic Hessian and an RG boundary/matching rule."

	return Analysis{
		AnomalyKinetic: ak,
		EWProjection:   ew,
		Fields:         fields,
		Hypercharge: HyperchargeDiagnostic{
			MatterBLCoefficient:       blCoeff,
			MatterBLNormSquared:       bl,
			MatterBLContribution:      matterBLContribution,
			ScalarContactNormSquared:  contact,
			ScalarContactContribution: scalarContribution,
			CombinedBridgeNorm:        combined,
			ChargeTableKY:             ew.HyperchargeNormalizationKY,
			BoundarySin2:              ew.EqualNormalizedCouplingBoundarySin2,
		},
		DiagonalTraceGramSelectedAsRepresentationMetric: true,
		CanonicalGeneratorDiagnosticsDerived:            true,
		DiagonalTraceGramSelectedAsGaugeKineticHessian:  false,
		PhysicalGaugeCouplingsDerived:                   false,
		FineStructureDerived:                            false,
		RGBoundaryScaleDerived:                          false,
		HiddenObservedInputUsed:                         false,
		TruthStatement:                                  truth,
		RecommendedNextGate:                             "Gate 82 — Gauge Kinetic Action Selection / RG Boundary Coupling Audit",
		RemainingUnknowns: []string{
			"U-20D3F5-ACTION-SELECTS-U1-HESSIAN: prove or reject that the diagonal trace-Gram matrix is the gauge kinetic Hessian",
			"U-20D3F6-COUPLING-FROM-HESSIAN: derive physical g_Y and g_2 from finite kinetic terms, not from charge normalization alone",
			"U-20D3F7-RG-BOUNDARY: derive the boundary scale/coupling before alpha_em can be computed",
			"U-20D3F8-CENTRAL-U1-FATE: decide whether the central u(1) is projected, massive, global, or separately gauged",
		},
	}, nil
}

func makeDiagnostic(name string, norm2 float64, role string) CouplingDiagnostic {
	return CouplingDiagnostic{
		Name:               name,
		TraceNormSquared:   norm2,
		TraceNorm:          math.Sqrt(norm2),
		InverseNormSquared: 1.0 / norm2,
		Role:               role,
	}
}
