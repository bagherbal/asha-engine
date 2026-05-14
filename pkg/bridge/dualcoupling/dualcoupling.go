// Package dualcoupling implements Gate 73: the dual-carrier coupling tensor /
// action search.
//
// Gate 72 established that the Pati-Salam/u(4) current carrier and the
// Boolean/contact electroweak block carrier must remain distinct finite
// carriers.  This gate searches the possible coupling-action terms between
// them without collapsing the two carriers into a single target space.  The
// important result is a reduction of the naive 64-dimensional coupling tensor
// to a small symmetry-compatible bridge domain, while still refusing to select
// coefficients without a finite action theorem.
package dualcoupling

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/dualcarrier"
)

type CouplingSector struct {
	Name            string
	SourceSector    string
	TargetSector    string
	Dimension       int
	SymmetryAllowed bool
	Derived         bool
	Role            string
	Obstruction     string
}

type ActionTerm struct {
	Name                 string
	Formula              string
	Dimension            int
	SymmetryAllowed      bool
	CoefficientsSelected bool
	Derived              bool
	Obstruction          string
}

type Analysis struct {
	Previous dualcarrier.Analysis

	NaiveTensorDimension        int
	SymmetryCompatibleDimension int
	RejectedDimension           int

	SectorCouplings []CouplingSector
	ActionTerms     []ActionTerm

	Direct64TensorRejected            bool
	ColorContactCouplingRejected      bool
	LeptoquarkContactCouplingRejected bool
	AbelianBridgeDomainExposed        bool
	AbelianBridgeDimension            int
	AbelianCoefficientsSelected       bool

	ScalarCurrentCouplingDomainExposed bool
	ScalarCurrentCoefficientsSelected  bool
	CouplingTensorSelected             bool
	CouplingActionDerived              bool
	DualCarrierHessianComputable       bool
	ExchangeKernelUpdated              bool
	AttractiveScalarDerived            bool
	UpDownSplittingDerived             bool
	CondensationClaimAllowed           bool
	HiddenObservedInputUsed            bool

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
		prev, err := dualcarrier.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev dualcarrier.Analysis) (Analysis, error) {
	if !prev.DualCarrierSplitDefined || prev.PatiSalamCarrier.Dimension != 16 || prev.ContactCarrier.Dimension != 4 {
		return Analysis{}, fmt.Errorf("Gate 72 dual-carrier input mismatch: PS=%d contact=%d split=%v", prev.PatiSalamCarrier.Dimension, prev.ContactCarrier.Dimension, prev.DualCarrierSplitDefined)
	}

	naive := prev.CouplingTensorDimension

	sectors := []CouplingSector{
		{
			Name:            "central-contact-u1",
			SourceSector:    "central",
			TargetSector:    "contact-u1",
			Dimension:       1,
			SymmetryAllowed: true,
			Derived:         false,
			Role:            "abelian kinetic/current mixing candidate",
			Obstruction:     "central u(1) versus hypercharge/contact-u1 normalization is not selected",
		},
		{
			Name:            "b-minus-l-contact-u1",
			SourceSector:    "b-minus-l",
			TargetSector:    "contact-u1",
			Dimension:       1,
			SymmetryAllowed: true,
			Derived:         false,
			Role:            "B-L to contact-u1/hypercharge bridge candidate",
			Obstruction:     "B-L/contact-u1 mixing coefficient is open; this is the likely hypercharge bridge slot",
		},
		{
			Name:            "color-su3-contact",
			SourceSector:    "color-su3",
			TargetSector:    "contact-su2+contact-u1",
			Dimension:       8 * prev.ContactCarrier.Dimension,
			SymmetryAllowed: false,
			Derived:         false,
			Role:            "forbidden direct color-contact coupling",
			Obstruction:     "contact carrier has no color adjoint target; color remains on the Pati-Salam carrier",
		},
		{
			Name:            "leptoquark-contact",
			SourceSector:    "leptoquark",
			TargetSector:    "contact-su2+contact-u1",
			Dimension:       6 * prev.ContactCarrier.Dimension,
			SymmetryAllowed: false,
			Derived:         false,
			Role:            "forbidden direct leptoquark-contact coupling",
			Obstruction:     "no finite leptoquark carrier exists in the contact block; this requires a separate Pati-Salam-sector action",
		},
	}

	compatible := 0
	rejected := 0
	for _, s := range sectors {
		if s.SymmetryAllowed {
			compatible += s.Dimension
		} else {
			rejected += s.Dimension
		}
	}
	// The abstract tensor has 64 entries.  The sector audit explicitly accounts
	// for 58 entries: 2 symmetry-compatible abelian slots and 56 rejected
	// nonabelian/direct slots.  The remaining 6 formal entries correspond to
	// contact-su2 coupling to abelian source components; those would break the
	// carrier roles unless a scalar current source is derived, so they remain in
	// the scalar-current bridge term below rather than the raw j-A tensor.
	scalarCurrentBridgeDim := 6

	terms := []ActionTerm{
		{
			Name:                 "abelian kinetic bridge",
			Formula:              "S_ab = κ_c j_c A_u1 + κ_BL j_BL A_u1",
			Dimension:            compatible,
			SymmetryAllowed:      true,
			CoefficientsSelected: false,
			Derived:              false,
			Obstruction:          "κ_c and κ_BL are not fixed by the finite action; arbitrary values would be fitting",
		},
		{
			Name:                 "scalar-current bridge",
			Formula:              "S_scalar = <j_abelian, Φ†TΦ> + possible contact-su2 scalar currents",
			Dimension:            scalarCurrentBridgeDim,
			SymmetryAllowed:      true,
			CoefficientsSelected: false,
			Derived:              false,
			Obstruction:          "finite scalar current and its normalization are not yet derived",
		},
		{
			Name:                 "color-contact direct term",
			Formula:              "j_su3 · A_contact",
			Dimension:            8 * prev.ContactCarrier.Dimension,
			SymmetryAllowed:      false,
			CoefficientsSelected: false,
			Derived:              false,
			Obstruction:          "rejected: no contact color representation",
		},
		{
			Name:                 "leptoquark-contact direct term",
			Formula:              "j_LQ · A_contact",
			Dimension:            6 * prev.ContactCarrier.Dimension,
			SymmetryAllowed:      false,
			CoefficientsSelected: false,
			Derived:              false,
			Obstruction:          "rejected: no contact leptoquark representation",
		},
	}

	truth := "Gate 73 turns the raw 64-parameter dual-carrier tensor into a structured coupling-action search. Direct color/contact and leptoquark/contact terms are rejected by carrier representation. The only symmetry-compatible bridge currently exposed is a small abelian/scalar-current coupling domain, but its coefficients are not selected by the finite action. The next missing object is the finite abelian mixing and scalar-current normalization theorem."

	return Analysis{
		Previous:                           prev,
		NaiveTensorDimension:               naive,
		SymmetryCompatibleDimension:        compatible + scalarCurrentBridgeDim,
		RejectedDimension:                  rejected,
		SectorCouplings:                    sectors,
		ActionTerms:                        terms,
		Direct64TensorRejected:             true,
		ColorContactCouplingRejected:       true,
		LeptoquarkContactCouplingRejected:  true,
		AbelianBridgeDomainExposed:         compatible == 2,
		AbelianBridgeDimension:             compatible,
		AbelianCoefficientsSelected:        false,
		ScalarCurrentCouplingDomainExposed: scalarCurrentBridgeDim > 0,
		ScalarCurrentCoefficientsSelected:  false,
		CouplingTensorSelected:             false,
		CouplingActionDerived:              false,
		DualCarrierHessianComputable:       false,
		ExchangeKernelUpdated:              false,
		AttractiveScalarDerived:            false,
		UpDownSplittingDerived:             false,
		CondensationClaimAllowed:           false,
		HiddenObservedInputUsed:            false,
		TruthStatement:                     truth,
		RecommendedNextGate:                "Gate 74 — Abelian Mixing / Hypercharge Coupling Normalization Search",
		RemainingUnknowns: []string{
			"U-20D2C1-ABELIAN-MIXING: derive κ_c and κ_BL from finite kinetic/action data",
			"U-20D2C2-SCALAR-CURRENT: construct finite scalar/contact current Φ†TΦ and its normalization",
			"U-20D2C3-DUAL-HESSIAN: compute the coupled-carrier Hessian only after S_coupling is derived",
			"U-20D2C4-COLOR-DYNAMICS: keep color exchange on the Pati-Salam carrier with its own action",
			"U-20D2C5-LEPTOQUARK-DYNAMICS: derive or reject leptoquark exchange independently of the contact block",
			"U-20D2C6-UP-DOWN-SPLIT: no coupling-action term yet distinguishes top-like up from bottom-like down",
		},
	}, nil
}

func FormatSectorCouplings(xs []CouplingSector) string {
	ys := append([]CouplingSector(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Name < ys[j].Name })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(%s→%s, dim=%d, allowed=%v, derived=%v, obstruction=%s)", x.Name, x.SourceSector, x.TargetSector, x.Dimension, x.SymmetryAllowed, x.Derived, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatActionTerms(xs []ActionTerm) string {
	ys := append([]ActionTerm(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Name < ys[j].Name })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(dim=%d, allowed=%v, coeffs=%v, derived=%v, formula=%s, obstruction=%s)", x.Name, x.Dimension, x.SymmetryAllowed, x.CoefficientsSelected, x.Derived, x.Formula, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	return strings.Join(xs, "; ")
}
