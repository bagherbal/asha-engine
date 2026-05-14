// Package dualcarrier implements Gate 72: the dual-carrier gauge architecture
// split.
//
// Gate 71 showed that the full u(4) Fock/Pati-Salam current inventory should
// not be forced into the four-generator Boolean/contact block target.  This
// gate records the corrected architecture: Pati-Salam/color currents and the
// contact electroweak/scalar block live on two typed finite carriers.  The
// missing problem becomes a coupling tensor/action between carriers, not a
// direct embedding of one carrier into the other.
package dualcarrier

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/currentcontact"
)

type Carrier struct {
	Name        string
	Dimension   int
	Role        string
	Sectors     []CarrierSector
	Derived     bool
	Description string
}

type CarrierSector struct {
	Name      string
	Dimension int
	Role      string
}

type CouplingProblem struct {
	Name            string
	Domain          string
	Dimension       int
	Derived         bool
	Obstruction     string
	AllowedAsBridge bool
}

type Analysis struct {
	Previous currentcontact.Analysis

	PatiSalamCarrier Carrier
	ContactCarrier   Carrier

	ForcedEmbeddingRejected       bool
	DualCarrierSplitDefined       bool
	PatiSalamCarrierPreservesU4   bool
	ContactCarrierPreservesEWSeed bool
	ColorCarrierPreserved         bool
	LeptoquarkCarrierPreserved    bool
	AbelianSeparationStillOpen    bool

	DirectEmbeddingDimension int
	DirectEmbeddingKernelMin int
	CouplingTensorDimension  int

	CouplingProblems []CouplingProblem

	DualCarrierActionTemplateAvailable bool
	CouplingTensorSelected             bool
	CouplingActionDerived              bool
	CurrentHessianComputable           bool
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
		prev, err := currentcontact.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev currentcontact.Analysis) (Analysis, error) {
	if prev.SourceGeneratorCount != 16 || prev.TargetBlockSeedCount != 4 {
		return Analysis{}, fmt.Errorf("Gate 71 input mismatch: source=%d target=%d", prev.SourceGeneratorCount, prev.TargetBlockSeedCount)
	}

	ps := Carrier{
		Name:        "Pati-Salam / Fock current carrier",
		Dimension:   prev.SourceGeneratorCount,
		Role:        "u(4)-shaped matter-current carrier acting on Fock/Yukawa sectors",
		Derived:     true,
		Description: "keeps color-su3 and leptoquark currents on the Fock/Pati-Salam side instead of compressing them into the contact block target",
		Sectors: []CarrierSector{
			{Name: "central", Dimension: 1, Role: "central u(1) current"},
			{Name: "color-su3", Dimension: 8, Role: "SU(3)c adjoint current"},
			{Name: "b-minus-l", Dimension: 1, Role: "Pati-Salam B-L current"},
			{Name: "leptoquark", Dimension: 6, Role: "off-diagonal lepton-color current"},
		},
	}
	contact := Carrier{
		Name:        "Boolean/contact electroweak block carrier",
		Dimension:   prev.TargetBlockSeedCount,
		Role:        "contact-preserving su(2)+u(1)-shaped block carrier on K⊕K⊥ and the scalar/contact sector",
		Derived:     prev.TargetHasSU2U1Shape,
		Description: "keeps the contact centralizer/electroweak-scalar block as its own finite carrier",
		Sectors: []CarrierSector{
			{Name: "contact-su2", Dimension: 3, Role: "derived nonabelian contact-preserving seed"},
			{Name: "contact-u1", Dimension: 1, Role: "derived abelian contact-preserving seed"},
		},
	}

	directDim := prev.AbstractMapDimension
	kernelMin := prev.MinimalKernelDimension
	couplingDim := ps.Dimension * contact.Dimension

	problems := []CouplingProblem{
		{
			Name:            "current-contact coupling tensor",
			Domain:          "Hom(V_u4, V_contact) or V_u4*⊗V_contact",
			Dimension:       couplingDim,
			Derived:         false,
			AllowedAsBridge: true,
			Obstruction:     "64 formal coefficients exist, but no finite action selects them",
		},
		{
			Name:            "block-diagonal carrier action",
			Domain:          "S_PS[j] + S_contact[A,Φ] + S_coupling[j,A,Φ]",
			Dimension:       0,
			Derived:         false,
			AllowedAsBridge: true,
			Obstruction:     "the two carrier actions exist as typed slots, but the cross-term is not constructed",
		},
		{
			Name:            "abelian mixing bridge",
			Domain:          "span{central,B-L} ↔ contact-u1/hypercharge",
			Dimension:       2,
			Derived:         false,
			AllowedAsBridge: true,
			Obstruction:     "central versus B-L separation and kinetic mixing are not selected by current finite data",
		},
		{
			Name:            "color-contact non-identification",
			Domain:          "su(3)c carrier × contact carrier",
			Dimension:       8 * contact.Dimension,
			Derived:         false,
			AllowedAsBridge: false,
			Obstruction:     "color remains on the Pati-Salam carrier; there is no direct contact su(3) target",
		},
	}

	truth := "Gate 72 replaces the failed single-carrier embedding with a dual-carrier gauge architecture. The u(4) Fock/Pati-Salam currents remain on a sixteen-dimensional matter-current carrier, while the Boolean/contact su(2)+u(1) seed remains on a four-dimensional contact/electroweak-scalar block carrier. This preserves color and leptoquark structure instead of crushing them into the contact target. The next missing object is not an embedding, but a finite coupling action/tensor between the two carriers."

	return Analysis{
		Previous:                           prev,
		PatiSalamCarrier:                   ps,
		ContactCarrier:                     contact,
		ForcedEmbeddingRejected:            !prev.CurrentToContactMapDerived,
		DualCarrierSplitDefined:            true,
		PatiSalamCarrierPreservesU4:        ps.Dimension == 16 && len(ps.Sectors) == 4,
		ContactCarrierPreservesEWSeed:      contact.Dimension == 4 && contact.Derived,
		ColorCarrierPreserved:              true,
		LeptoquarkCarrierPreserved:         true,
		AbelianSeparationStillOpen:         prev.AbelianAmbiguity,
		DirectEmbeddingDimension:           directDim,
		DirectEmbeddingKernelMin:           kernelMin,
		CouplingTensorDimension:            couplingDim,
		CouplingProblems:                   problems,
		DualCarrierActionTemplateAvailable: true,
		CouplingTensorSelected:             false,
		CouplingActionDerived:              false,
		CurrentHessianComputable:           false,
		ExchangeKernelUpdated:              false,
		AttractiveScalarDerived:            false,
		UpDownSplittingDerived:             false,
		CondensationClaimAllowed:           false,
		HiddenObservedInputUsed:            false,
		TruthStatement:                     truth,
		RecommendedNextGate:                "Gate 73 — Dual-Carrier Coupling Tensor / Action Search",
		RemainingUnknowns: []string{
			"U-20D2B5C-DUAL-CARRIER-COUPLING: derive the coupling action between u(4) matter currents and the contact/electroweak block",
			"U-20D2B5F-ABELIAN-SEPARATION: separate central u(1), B-L, and contact-u1/hypercharge kinetically",
			"U-20D2B6-ACTION-HESSIAN: compute K_current only after the dual-carrier coupling action is derived",
			"U-20D2B7-COLOR-CARRIER-DYNAMICS: keep SU(3)c on its native Fock/Pati-Salam carrier and derive its independent kinetic action",
			"U-20D2B8-LEPTOQUARK-DYNAMICS: derive or reject finite leptoquark exchange as a separate Pati-Salam-sector effect",
			"U-20D2B9-NJL-KERNEL: no native four-fermion kernel follows until the carrier coupling and propagator rule are derived",
		},
	}, nil
}

func FormatCarrier(c Carrier) string {
	parts := make([]string, 0, len(c.Sectors))
	for _, s := range c.Sectors {
		parts = append(parts, fmt.Sprintf("%s:%d", s.Name, s.Dimension))
	}
	return fmt.Sprintf("%s(dim=%d, sectors=[%s], derived=%v)", c.Name, c.Dimension, strings.Join(parts, ", "), c.Derived)
}

func FormatCouplingProblems(xs []CouplingProblem) string {
	ys := append([]CouplingProblem(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Name < ys[j].Name })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(domain=%s, dim=%d, derived=%v, bridge=%v, obstruction=%s)", x.Name, x.Domain, x.Dimension, x.Derived, x.AllowedAsBridge, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
