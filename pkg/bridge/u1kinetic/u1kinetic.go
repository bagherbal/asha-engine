// Package u1kinetic implements Gate 75: the U(1) kinetic-mixing / gauge-
// coupling Hessian search.
//
// Gate 74 selected the charge-level hypercharge direction:
//
//	Y = T3_R + (B-L)/2,
//
// and rejected the central u(1) as a hypercharge component.  That is still only
// a charge-table result.  A physical U(1)_Y coupling requires a kinetic Hessian
// for the abelian gauge fields.  This package constructs the finite trace-Gram
// diagnostics for the available abelian generators:
//
//	central u(1) on the Fock/Pati-Salam carrier,
//	B-L on the Fock/Pati-Salam carrier,
//	contact-u1 / scalar T_phi on the contact scalar carrier.
//
// The result is deliberately strict.  The matter-carrier central/B-L Gram block
// and the scalar contact-u1 norm are computable, but the cross-carrier kinetic
// mixing entries are not selected by the finite action.  Therefore no physical
// U(1) gauge coupling, fine-structure constant, or low-energy alpha is derived.
package u1kinetic

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/abelianmixing"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type GeneratorTrace struct {
	Name      string
	Carrier   string
	Dimension int
	Trace     float64
	Trace2    float64
	Role      string
}

type Analysis struct {
	AbelianMixing abelianmixing.Analysis

	MatterCarrierDimension int
	ScalarCarrierDimension int
	AbelianFieldCount      int

	Central   GeneratorTrace
	BMinusL   GeneratorTrace
	ContactU1 GeneratorTrace

	MatterGram              linear.Matrix // central/B-L trace Gram on the 16D Fock carrier.
	BlockDiagonalDiagnostic linear.Matrix // central, B-L, contact-u1, with cross-carrier entries set to zero diagnostically.

	CentralBMinusLOrthogonal         bool
	MatterGramDerived                bool
	ContactU1NormDerived             bool
	CrossCarrierKineticMixingDerived bool
	FullU1KineticHessianDerived      bool

	MatterGramDeterminant      float64
	BlockDiagnosticDeterminant float64
	CentralNorm                float64
	BMinusLNorm                float64
	ContactU1Norm              float64

	ChargeLevelHyperchargeSelected bool
	PhysicalU1CouplingDerived      bool
	FineStructureDerived           bool
	HiddenObservedInputUsed        bool

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
		am, err := abelianmixing.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(am, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(am abelianmixing.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !am.ChargeLevelHyperchargeBridge {
		return Analysis{}, fmt.Errorf("Gate 74 must select the charge-level hypercharge bridge before U(1) kinetic audit")
	}

	// One-generation Fock/Pati-Salam carrier has 16 occupation states.  The
	// central generator is the identity.  B-L is q=-N0+(N1+N2+N3)/3.  Summing over
	// all 16 states gives Tr(B-L)=0 and Tr((B-L)^2)=16/3.
	centralTrace := 16.0
	centralTrace2 := 16.0
	bTrace := 0.0
	bTrace2 := 16.0 / 3.0
	centralB := 0.0

	// The scalar/contact u1 is the finite scalar doublet charge T_phi with active
	// real spectrum (+1/2,+1/2,-1/2,-1/2), so Tr(T_phi)=0 and Tr(T_phi^2)=1.
	contactTrace := 0.0
	contactTrace2 := 1.0

	matterGram, err := linear.FromRows([][]float64{{centralTrace2, centralB}, {centralB, bTrace2}})
	if err != nil {
		return Analysis{}, err
	}
	block, err := linear.FromRows([][]float64{{centralTrace2, centralB, 0}, {centralB, bTrace2, 0}, {0, 0, contactTrace2}})
	if err != nil {
		return Analysis{}, err
	}
	matterDet := centralTrace2*bTrace2 - centralB*centralB
	blockDet := matterDet * contactTrace2

	truth := "Gate 75 derives the finite trace-Gram block for central u(1) and B-L on the 16-state Fock carrier, and the scalar/contact-u1 norm on the 4-real-dimensional scalar carrier. It does not derive the cross-carrier kinetic-mixing Hessian. Therefore the charge-level hypercharge direction is known, but the physical U(1)_Y gauge coupling and alpha_em remain open."

	return Analysis{
		AbelianMixing:                    am,
		MatterCarrierDimension:           16,
		ScalarCarrierDimension:           4,
		AbelianFieldCount:                3,
		Central:                          GeneratorTrace{Name: "central u(1)", Carrier: "Fock/Pati-Salam", Dimension: 16, Trace: centralTrace, Trace2: centralTrace2, Role: "universal finite current; rejected as hypercharge component by Gate 74"},
		BMinusL:                          GeneratorTrace{Name: "B-L", Carrier: "Fock/Pati-Salam", Dimension: 16, Trace: bTrace, Trace2: bTrace2, Role: "matter-side abelian charge entering Y=T3_R+(B-L)/2"},
		ContactU1:                        GeneratorTrace{Name: "contact-u1 / T_phi", Carrier: "scalar/contact", Dimension: 4, Trace: contactTrace, Trace2: contactTrace2, Role: "scalar-side doublet charge with spectrum (+1/2,+1/2,-1/2,-1/2)"},
		MatterGram:                       matterGram,
		BlockDiagonalDiagnostic:          block,
		CentralBMinusLOrthogonal:         math.Abs(centralB) <= eps,
		MatterGramDerived:                true,
		ContactU1NormDerived:             true,
		CrossCarrierKineticMixingDerived: false,
		FullU1KineticHessianDerived:      false,
		MatterGramDeterminant:            matterDet,
		BlockDiagnosticDeterminant:       blockDet,
		CentralNorm:                      math.Sqrt(centralTrace2),
		BMinusLNorm:                      math.Sqrt(bTrace2),
		ContactU1Norm:                    math.Sqrt(contactTrace2),
		ChargeLevelHyperchargeSelected:   am.ChargeLevelHyperchargeBridge,
		PhysicalU1CouplingDerived:        false,
		FineStructureDerived:             false,
		HiddenObservedInputUsed:          false,
		TruthStatement:                   truth,
		RecommendedNextGate:              "Gate 76 — Contact-u1 / B-L Kinetic Hessian Source Search",
		RemainingUnknowns: []string{
			"U-20D3A-U1-KINETIC-HESSIAN: derive the full kinetic matrix for central, B-L, and contact-u1 fields",
			"U-20D3B-CROSS-CARRIER-MIXING: derive or reject off-diagonal B-L/contact-u1 kinetic terms from a finite action",
			"U-20D3C-CENTRAL-U1-FATE: decide whether central u(1) is projected out, massive, global, or separately gauged",
			"U-20D3D-PHYSICAL-COUPLING: no g_Y, alpha_em, or low-energy weak angle follows until the kinetic Hessian and RG scale are derived",
		},
	}, nil
}
