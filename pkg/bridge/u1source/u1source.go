// Package u1source implements Gate 76: the contact-u1 / B-L kinetic-Hessian
// source search.
//
// Gate 75 derived finite trace-Gram diagnostics for the abelian generators on
// their separate carriers:
//
//	central u(1), B-L      on the Fock/Pati-Salam carrier,
//	contact-u1 / T_phi     on the scalar/contact carrier.
//
// This package asks whether the existing finite data selects a cross-carrier
// kinetic source term such as K(B-L, T_phi).  The result is deliberately strict:
// the factorized trace candidates vanish because both B-L and T_phi are
// traceless, and no non-factorized finite action term has yet been derived.
// Thus the charge-level hypercharge direction remains valid, but the physical
// U(1)_Y kinetic Hessian and coupling are still bridge-open.
package u1source

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/u1kinetic"
)

type CrossTerm struct {
	Name      string
	Left      string
	Right     string
	Candidate float64
	Derived   bool
	Detail    string
}

type Analysis struct {
	U1Kinetic u1kinetic.Analysis

	Fields []string

	FactorizedCentralContact  float64
	FactorizedBLContact       float64
	TensorTraceCentralContact float64
	TensorTraceBLContact      float64

	FactorizedTraceSourceDerived bool
	NonFactorizedActionDerived   bool
	CrossCarrierSourceDerived    bool
	FullHessianDerived           bool

	CandidateCrossTerms []CrossTerm

	CentralContactForbiddenForHypercharge bool
	BLContactSourceNonzero                bool
	ContactU1TraceZero                    bool
	BMinusLTraceZero                      bool

	PhysicalU1CouplingDerived bool
	FineStructureDerived      bool
	HiddenObservedInputUsed   bool

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
		u1, err := u1kinetic.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(u1, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(u1 u1kinetic.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !u1.ChargeLevelHyperchargeSelected {
		return Analysis{}, fmt.Errorf("Gate 75 requires the Gate 74 charge-level hypercharge selection")
	}

	// The only cross-source candidates currently available without a new action
	// are factorized trace pairings between matter-carrier abelian generators and
	// the scalar/contact T_phi generator.  They vanish because T_phi is traceless,
	// while B-L is also traceless on the 16-state Fock carrier.
	centralContact := u1.Central.Trace * u1.ContactU1.Trace
	blContact := u1.BMinusL.Trace * u1.ContactU1.Trace

	// The tensor-product trace gives the same factorized result.
	tensorCentralContact := centralContact
	tensorBLContact := blContact

	contactTraceZero := math.Abs(u1.ContactU1.Trace) <= eps
	blTraceZero := math.Abs(u1.BMinusL.Trace) <= eps
	blContactNonzero := math.Abs(blContact) > eps || math.Abs(tensorBLContact) > eps

	terms := []CrossTerm{
		{
			Name:      "central/contact factorized trace",
			Left:      "central u(1)",
			Right:     "contact-u1 / T_phi",
			Candidate: centralContact,
			Derived:   false,
			Detail:    "factorized trace vanishes because Tr(T_phi)=0; central u(1) was already rejected as a hypercharge component",
		},
		{
			Name:      "B-L/contact factorized trace",
			Left:      "B-L",
			Right:     "contact-u1 / T_phi",
			Candidate: blContact,
			Derived:   false,
			Detail:    "factorized trace vanishes because Tr(B-L)=0 and Tr(T_phi)=0",
		},
	}

	truth := "Gate 76 tests the only cross-carrier U(1) kinetic-source candidates available from the current finite data: factorized trace pairings between matter abelian charges and contact-u1. They vanish. No non-factorized finite action term coupling B-L to contact-u1 is derived. Therefore the hypercharge direction remains charge-level, while the physical U(1)_Y kinetic Hessian and coupling remain open."

	return Analysis{
		U1Kinetic:                             u1,
		Fields:                                []string{"central u(1)", "B-L", "contact-u1 / T_phi"},
		FactorizedCentralContact:              centralContact,
		FactorizedBLContact:                   blContact,
		TensorTraceCentralContact:             tensorCentralContact,
		TensorTraceBLContact:                  tensorBLContact,
		FactorizedTraceSourceDerived:          true,
		NonFactorizedActionDerived:            false,
		CrossCarrierSourceDerived:             false,
		FullHessianDerived:                    false,
		CandidateCrossTerms:                   terms,
		CentralContactForbiddenForHypercharge: true,
		BLContactSourceNonzero:                blContactNonzero,
		ContactU1TraceZero:                    contactTraceZero,
		BMinusLTraceZero:                      blTraceZero,
		PhysicalU1CouplingDerived:             false,
		FineStructureDerived:                  false,
		HiddenObservedInputUsed:               false,
		TruthStatement:                        truth,
		RecommendedNextGate:                   "Gate 77 — Non-Factorized Abelian Action / Kinetic-Mixing Search",
		RemainingUnknowns: []string{
			"U-20D3B1-NONFACTORIZED-U1-ACTION: derive a finite action term that pairs B-L with contact-u1 beyond factorized trace",
			"U-20D3B2-CONTACT-U1-HESSIAN: compute the contact-u1 second variation from scalar/contact dynamics rather than using trace norm only",
			"U-20D3B3-HYPERCHARGE-KINETIC-MATRIX: derive the physical U(1)_Y kinetic matrix after abelian projection",
			"U-20D3B4-ALPHA: no alpha_em follows until the kinetic matrix, RG scale, and threshold matching are derived",
		},
	}, nil
}
