// Package generation2tauetacarrierpullbackobstructionaudit implements Gate 556:
// Tau-Eta Carrier Pullback Obstruction Audit.
//
// Gate 555 proved the native selector algebra and showed that tau_eta has a
// formal 2+1 selector shape only after a pullback. Gate 556 asks whether the
// project already contains the missing native source algebra and
// unit-preserving carrier representation. It resolves the historical tension
// between conditional tau_eta weak-plane/generation-capacity gates and strict
// native operator promotion: tau_eta is an eta-graded scalar/contact trace
// vector with conditional selector capacity, not currently an operator on
// W_spatial or a generation carrier.
package generation2tauetacarrierpullbackobstructionaudit

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE556-TAU-ETA-CARRIER-PULLBACK-OBSTRUCTION-AUDIT"

	StatusGate555Inherited                   = "CONDITIONAL_SUPPORT_GATE555_SELECTOR_THEOREM_INHERITED"
	StatusTauEtaTypedAsTraceVector           = "PASS_TAU_ETA_TYPED_AS_ETA_GRADED_TRACE_VALUE_VECTOR"
	StatusTauEtaNotNativeOperator            = "FAILED_ROUTE_TAU_ETA_NOT_NATIVE_OPERATOR_SPECTRUM_OR_ENDOMORPHISM"
	StatusNoNativeTauSourceAlgebra           = "FAILED_ROUTE_NO_NATIVE_TAU_SOURCE_ALGEBRA"
	StatusFormalTauAlgebrasQuarantined       = "CONDITIONAL_SUPPORT_FORMAL_TAU_POLYNOMIAL_ALGEBRAS_IDENTIFIED_AS_NON_NATIVE_FIXTURES"
	StatusNoUnitPreservingCarrierRep         = "FAILED_ROUTE_NO_TAU_ETA_UNIT_PRESERVING_CARRIER_REPRESENTATION"
	StatusTauEtaSelectorCapacitySealed       = "SEALED_SUPPORT_TAU_ETA_HAS_2PLUS1_SELECTOR_CAPACITY"
	StatusTauSelectorBasisDependent          = "FAILED_ROUTE_TAU_SELECTOR_BASIS_DEPENDENT_NO_CANONICAL_U12"
	StatusBMinusLCompatibilityConditional    = "CONDITIONAL_SUPPORT_FORMAL_SPATIAL_TAU_OPERATOR_WOULD_COMMUTE_WITH_B_MINUS_L"
	StatusSpectralTripleCompatibilityMissing = "FAILED_ROUTE_TAU_ETA_SPECTRAL_TRIPLE_COMPATIBILITY_DATA_MISSING"
	StatusNoCanonicalTwoPlusOneSelector      = "FAILED_ROUTE_TAU_ETA_DOES_NOT_PRODUCE_CANONICAL_2PLUS1_SELECTOR_ON_W_SPATIAL"
	StatusFirewallPreserved                  = "FIREWALL_PRESERVED_GATE556_TAU_ETA_TRACE_VECTOR_BOUNDARY"
)

type InheritedGate555Audit struct {
	SelectorTheoremProved      bool
	BMinusLFourToOnePlusThree  bool
	BMinusLUniqueWeakPlane     bool
	TauEtaPullbackValid        bool
	TauEtaSealedCapacity       bool
	ContactQuarticStillBlocked bool
	Verdict                    string
}

type TypeClassificationAudit struct {
	TauEta                           []int
	AbsTauEta                        []int
	IsTraceValueVector               bool
	IsSpectrumOfNativeOperator       bool
	IsDiagonalEndomorphism           bool
	IsCharacter                      bool
	IsCoefficientVectorInNativeBasis bool
	IsSealedBookkeepingDatum         bool
	SourceFunctional                 string
	SourceOperators                  []string
	PriorSupport                     []string
	Verdict                          string
	Reason                           string
}

type SourceAlgebraCandidate struct {
	Name                        string
	Presentation                string
	Unit                        string
	InsertedByHand              bool
	FoundAsNativeProjectAlgebra bool
	CarriesTauEta               bool
	RejectedReason              string
}

type SourceAlgebraAudit struct {
	Candidates                     []SourceAlgebraCandidate
	NativeSourceAlgebraExists      bool
	NativeSourceAlgebraName        string
	NativeElementName              string
	HasUnit                        bool
	TraceOrEigenvalueDataRecovered bool
	Verdict                        string
	Reason                         string
}

type RepresentationCandidate struct {
	Target                    string
	NativeSourceAvailable     bool
	RepresentationConstructed bool
	UnitPreserving            bool
	RhoOneEqualsIdentity      bool
	ActsOnWSpatial            bool
	ActsOnGenerationCarrier   bool
	CompatibilityKnown        bool
	RejectedReason            string
}

type RepresentationAudit struct {
	Candidates                           []RepresentationCandidate
	AnyValidUnitPreservingRepresentation bool
	ValidTarget                          string
	RhoOneIsIdentity                     bool
	Verdict                              string
	Reason                               string
}

type SelectorConsequenceAudit struct {
	FormalAbsTauEta                                     []int
	FormalMultiplicityPattern                           []int
	FormalCommutant                                     string
	FormalCommutantDimension                            int
	Gate555SelectorFormulaAppliesIfRepresentationExists bool
	ValidRepresentationExists                           bool
	ProducesNativeSelector                              bool
	CanonicalU12Selected                                bool
	BasisDependentIfForced                              bool
	Verdict                                             string
	Reason                                              string
}

type BMinusLCompatibilityAudit struct {
	ActsOnWSpatial                  bool
	ValidRepresentationExists       bool
	BMinusLRestrictedToWSpatial     string
	FormalCommutatorWithBMinusLZero bool
	NativeCompatibilityVerified     bool
	Verdict                         string
	Reason                          string
}

type SpectralTripleCompatibilityAudit struct {
	ProposedNativeRepresentation         bool
	GammaCompatibilityAvailable          bool
	JCompatibilityAvailable              bool
	DCompatibilityAvailable              bool
	FirstOrderCompatibilityAvailable     bool
	BMinusLCompatibilityAvailable        bool
	MissingData                          []string
	NativeSpectralTriplePromotionAllowed bool
	Verdict                              string
}

type FirewallAudit struct {
	PromotedToWeakIsospin             bool
	PromotedToGenerationMassHierarchy bool
	PromotedToHiggs                   bool
	PromotedToYukawa                  bool
	PromotedToCKMPMNS                 bool
	InsertedObservedFlavorData        bool
	InsertedFormalAlgebraAsNative     bool
	InsertedDiagonalMatrixAsNative    bool
	NativeRegistryPolluted            bool
	Verdict                           string
}

type FinalVerdict struct {
	TauEtaOperator                        bool
	TauEtaOnlyTraceVector                 bool
	NativeSourceAlgebraExists             bool
	UnitPreservingRepresentationExists    bool
	CanonicalTwoPlusOneSelectorOnWSpatial bool
	MissingNextTheorem                    string
	Verdict                               string
}

type Analysis struct {
	Inherited      InheritedGate555Audit
	Type           TypeClassificationAudit
	SourceAlgebra  SourceAlgebraAudit
	Representation RepresentationAudit
	Selector       SelectorConsequenceAudit
	BMinusL        BMinusLCompatibilityAudit
	SpectralTriple SpectralTripleCompatibilityAudit
	Firewall       FirewallAudit
	Final          FinalVerdict
	Truth          string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	inherited := buildInheritedGate555()
	typ := classifyTauEtaType()
	source := auditSourceAlgebra(typ)
	rep := auditRepresentations(source)
	selector := auditSelectorConsequence(rep)
	bminusl := auditBMinusLCompatibility(rep)
	spectral := auditSpectralTriple(rep, bminusl)
	firewall := auditFirewall()
	final := buildFinal(typ, source, rep, selector)
	a := Analysis{Inherited: inherited, Type: typ, SourceAlgebra: source, Representation: rep, Selector: selector, BMinusL: bminusl, SpectralTriple: spectral, Firewall: firewall, Final: final, Truth: truth(typ, source, rep, selector, spectral)}
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritedGate555() InheritedGate555Audit {
	return InheritedGate555Audit{
		SelectorTheoremProved:      true,
		BMinusLFourToOnePlusThree:  true,
		BMinusLUniqueWeakPlane:     false,
		TauEtaPullbackValid:        false,
		TauEtaSealedCapacity:       true,
		ContactQuarticStillBlocked: true,
		Verdict:                    StatusGate555Inherited + "; Gate 555 supplies the selector algebra but explicitly leaves tau_eta without a unit-preserving Fock/generation pullback",
	}
}

func classifyTauEtaType() TypeClassificationAudit {
	return TypeClassificationAudit{
		TauEta:                           []int{2, -2, 1},
		AbsTauEta:                        []int{2, 2, 1},
		IsTraceValueVector:               true,
		IsSpectrumOfNativeOperator:       false,
		IsDiagonalEndomorphism:           false,
		IsCharacter:                      false,
		IsCoefficientVectorInNativeBasis: false,
		IsSealedBookkeepingDatum:         true,
		SourceFunctional:                 "tau_eta(O)=Tr_HPhi(eta O)",
		SourceOperators:                  []string{"Q^T Q", "Z^T Z", "T3L^T Y_phi"},
		PriorSupport: []string{
			"Gate 244 traces tau_eta to eta-graded scalar-bundle trace records, not spatial Fock-mode projectors.",
			"Gate 259 uses tau_eta only under SpontaneousCarrierSeal as conditional orientation data, not native Fock pullback.",
			"Gate 555 records SEALED_SUPPORT_TAU_ETA_HAS_2PLUS1_SELECTOR_CAPACITY and FAILED_ROUTE_NO_TAU_ETA_FOCK_PULLBACK.",
		},
		Verdict: strings.Join([]string{StatusTauEtaTypedAsTraceVector, StatusTauEtaNotNativeOperator}, "; "),
		Reason:  "The current project supplies tau_eta as three stable eta-graded scalar/contact trace values. It does not supply an operator t_tau whose spectrum is tau_eta, a diagonal endomorphism on W_spatial, a character of a native algebra, or coefficients in a native spatial basis.",
	}
}

func auditSourceAlgebra(typ TypeClassificationAudit) SourceAlgebraAudit {
	candidates := []SourceAlgebraCandidate{
		{Name: "A_tau_signed", Presentation: "Q[t]/((t-2)(t+2)(t-1))", Unit: "1", InsertedByHand: true, FoundAsNativeProjectAlgebra: false, CarriesTauEta: true, RejectedReason: "This algebra can encode the three signed values formally, but the project does not derive it as a native source algebra."},
		{Name: "A_tau_abs", Presentation: "Q[t]/((t-2)(t-1))", Unit: "1", InsertedByHand: true, FoundAsNativeProjectAlgebra: false, CarriesTauEta: true, RejectedReason: "This algebra can encode the magnitude pattern (2,2,1) formally, but it collapses the sign data and is not found natively."},
	}
	_ = typ
	return SourceAlgebraAudit{
		Candidates:                     candidates,
		NativeSourceAlgebraExists:      false,
		NativeSourceAlgebraName:        "",
		NativeElementName:              "",
		HasUnit:                        false,
		TraceOrEigenvalueDataRecovered: false,
		Verdict:                        strings.Join([]string{StatusNoNativeTauSourceAlgebra, StatusFormalTauAlgebrasQuarantined}, "; "),
		Reason:                         "No existing ASHA package presents a native unit algebra A_tau with element t_tau whose representation-independent trace/eigenvalue data is tau_eta or |tau_eta|. The displayed polynomial algebras are mathematically possible fixtures, not project-derived native inputs.",
	}
}

func auditRepresentations(source SourceAlgebraAudit) RepresentationAudit {
	candidates := []RepresentationCandidate{
		{Target: "End(W_spatial)", NativeSourceAvailable: source.NativeSourceAlgebraExists, RepresentationConstructed: false, UnitPreserving: false, RhoOneEqualsIdentity: false, ActsOnWSpatial: true, CompatibilityKnown: false, RejectedReason: "No native A_tau and no rho_tau(1)=I_3 action on span_C{a_1^dagger,a_2^dagger,a_3^dagger} are present."},
		{Target: "End(C^3_gen)", NativeSourceAvailable: source.NativeSourceAlgebraExists, RepresentationConstructed: false, UnitPreserving: false, RhoOneEqualsIdentity: false, ActsOnGenerationCarrier: true, CompatibilityKnown: false, RejectedReason: "Prior gates support generation-breaking capacity, but no unit-preserving generation representation with spectral-triple compatibility is constructed."},
	}
	return RepresentationAudit{
		Candidates:                           candidates,
		AnyValidUnitPreservingRepresentation: false,
		ValidTarget:                          "",
		RhoOneIsIdentity:                     false,
		Verdict:                              StatusNoUnitPreservingCarrierRep,
		Reason:                               "Because there is no native source algebra and no constructed rho_tau, the unit test rho_tau(1_A_tau)=I cannot even be executed. Conditional orientation/generation-capacity language remains sealed support, not representation theory.",
	}
}

func auditSelectorConsequence(rep RepresentationAudit) SelectorConsequenceAudit {
	return SelectorConsequenceAudit{
		FormalAbsTauEta:           []int{2, 2, 1},
		FormalMultiplicityPattern: []int{2, 1},
		FormalCommutant:           "u(2)+u(1)",
		FormalCommutantDimension:  5,
		Gate555SelectorFormulaAppliesIfRepresentationExists: true,
		ValidRepresentationExists:                           rep.AnyValidUnitPreservingRepresentation,
		ProducesNativeSelector:                              false,
		CanonicalU12Selected:                                false,
		BasisDependentIfForced:                              true,
		Verdict:                                             strings.Join([]string{StatusTauEtaSelectorCapacitySealed, StatusTauSelectorBasisDependent, StatusNoCanonicalTwoPlusOneSelector}, "; "),
		Reason:                                              "If a valid unit-preserving spatial representation existed and if |tau_eta| were represented as diag(2,2,1), Gate 555 would give Comm=u(2)+u(1) of dimension 5. But the identification of the double eigenvalue with U_12 is a basis choice unless a native carrier map labels the tau slots by a_1^dagger,a_2^dagger,a_3^dagger.",
	}
}

func auditBMinusLCompatibility(rep RepresentationAudit) BMinusLCompatibilityAudit {
	return BMinusLCompatibilityAudit{
		ActsOnWSpatial:                  false,
		ValidRepresentationExists:       rep.AnyValidUnitPreservingRepresentation,
		BMinusLRestrictedToWSpatial:     "B-L = (1/3) I on W_spatial",
		FormalCommutatorWithBMinusLZero: true,
		NativeCompatibilityVerified:     false,
		Verdict:                         StatusBMinusLCompatibilityConditional,
		Reason:                          "Any formal endomorphism of W_spatial would commute with the restricted B-L scalar (1/3)I. This is only a conditional algebraic observation; without rho_tau, no native B-L refinement is verified.",
	}
}

func auditSpectralTriple(rep RepresentationAudit, b BMinusLCompatibilityAudit) SpectralTripleCompatibilityAudit {
	missing := []string{"native A_tau", "rho_tau(1)=I carrier representation", "gamma-compatibility datum", "J-compatibility datum", "D-compatibility datum", "first-order-condition check", "native B-L refinement proof"}
	return SpectralTripleCompatibilityAudit{
		ProposedNativeRepresentation:         rep.AnyValidUnitPreservingRepresentation,
		GammaCompatibilityAvailable:          false,
		JCompatibilityAvailable:              false,
		DCompatibilityAvailable:              false,
		FirstOrderCompatibilityAvailable:     false,
		BMinusLCompatibilityAvailable:        b.NativeCompatibilityVerified,
		MissingData:                          missing,
		NativeSpectralTriplePromotionAllowed: false,
		Verdict:                              StatusSpectralTripleCompatibilityMissing,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{Verdict: StatusFirewallPreserved}
}

func buildFinal(typ TypeClassificationAudit, source SourceAlgebraAudit, rep RepresentationAudit, sel SelectorConsequenceAudit) FinalVerdict {
	return FinalVerdict{
		TauEtaOperator:                        typ.IsSpectrumOfNativeOperator || typ.IsDiagonalEndomorphism,
		TauEtaOnlyTraceVector:                 typ.IsTraceValueVector && typ.IsSealedBookkeepingDatum,
		NativeSourceAlgebraExists:             source.NativeSourceAlgebraExists,
		UnitPreservingRepresentationExists:    rep.AnyValidUnitPreservingRepresentation,
		CanonicalTwoPlusOneSelectorOnWSpatial: sel.ProducesNativeSelector && sel.CanonicalU12Selected,
		MissingNextTheorem:                    "Construct a native unit algebra A_tau and a unit-preserving representation rho_tau:A_tau->End(W_spatial) or End(C^3_gen), prove rho_tau(1)=I, derive tau-slot-to-carrier labels basis-independently, and verify compatibility with gamma, J, D, first-order condition, and B-L.",
		Verdict:                               StatusNoCanonicalTwoPlusOneSelector,
	}
}

func validate(a Analysis) error {
	if !a.Inherited.SelectorTheoremProved || a.Inherited.BMinusLUniqueWeakPlane || a.Inherited.TauEtaPullbackValid {
		return fmt.Errorf("Gate 555 inheritance inconsistent")
	}
	if !a.Type.IsTraceValueVector || a.Type.IsSpectrumOfNativeOperator || a.Type.IsDiagonalEndomorphism || !a.Type.IsSealedBookkeepingDatum {
		return fmt.Errorf("tau_eta type classification illegally promoted trace data")
	}
	if a.SourceAlgebra.NativeSourceAlgebraExists || a.SourceAlgebra.HasUnit || a.SourceAlgebra.TraceOrEigenvalueDataRecovered {
		return fmt.Errorf("native tau source algebra was inserted without derivation")
	}
	if a.Representation.AnyValidUnitPreservingRepresentation || a.Representation.RhoOneIsIdentity {
		return fmt.Errorf("tau_eta representation was illegally promoted")
	}
	if a.Selector.ProducesNativeSelector || a.Selector.CanonicalU12Selected || a.Selector.FormalCommutantDimension != 5 {
		return fmt.Errorf("selector consequence audit invalid")
	}
	if a.SpectralTriple.NativeSpectralTriplePromotionAllowed || len(a.SpectralTriple.MissingData) == 0 {
		return fmt.Errorf("spectral triple firewall failed")
	}
	if a.Firewall.PromotedToWeakIsospin || a.Firewall.PromotedToGenerationMassHierarchy || a.Firewall.PromotedToHiggs || a.Firewall.PromotedToYukawa || a.Firewall.PromotedToCKMPMNS || a.Firewall.InsertedFormalAlgebraAsNative || a.Firewall.InsertedDiagonalMatrixAsNative || a.Firewall.NativeRegistryPolluted {
		return fmt.Errorf("firewall pollution detected")
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate555Inherited,
		StatusTauEtaTypedAsTraceVector,
		StatusTauEtaNotNativeOperator,
		StatusNoNativeTauSourceAlgebra,
		StatusFormalTauAlgebrasQuarantined,
		StatusNoUnitPreservingCarrierRep,
		StatusTauEtaSelectorCapacitySealed,
		StatusTauSelectorBasisDependent,
		StatusBMinusLCompatibilityConditional,
		StatusSpectralTripleCompatibilityMissing,
		StatusNoCanonicalTwoPlusOneSelector,
		StatusFirewallPreserved,
	}
}

func truth(typ TypeClassificationAudit, source SourceAlgebraAudit, rep RepresentationAudit, sel SelectorConsequenceAudit, st SpectralTripleCompatibilityAudit) string {
	return fmt.Sprintf("Gate 556 classifies tau_eta=%v as %s with source functional %s on %v. It finds no native A_tau (%t), no unit-preserving rho_tau (%t), and no spectral-triple compatibility package (%v). The magnitude pattern %v would have u(2)+u(1) selector capacity of dimension %d only after a valid carrier pullback; without that pullback, U_12 is basis-labelled rather than native.", typ.TauEta, "an eta-graded scalar/contact trace-value vector", typ.SourceFunctional, typ.SourceOperators, source.NativeSourceAlgebraExists, rep.AnyValidUnitPreservingRepresentation, st.MissingData, sel.FormalAbsTauEta, sel.FormalCommutantDimension)
}
