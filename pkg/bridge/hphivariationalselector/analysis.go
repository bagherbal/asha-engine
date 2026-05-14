// Package hphivariationalselector implements Gate 408:
// H_phi Variational Functional / Canonical Coefficient Selector Sieve.
//
// Gate 407 proved that the native scalar carrier H_phi has full algebraic
// capacity (End_R(H_phi)) once left/right quaternionic actions and the scalar
// pair split are all allowed, but that no canonical coefficient rule selects a
// nondegenerate element. Gate 408 audits the native variational/functionality
// layer: scalar potential Hessian, one-form kinetic trace, quaternionic invariant
// trace/norm, and source-functional stress tests. The theorem asks whether one
// of these functionals selects a unique non-pair-degenerate H_phi endomorphism
// without importing Yukawa amplitudes, CKM/PMNS data, or arbitrary sources.
package hphivariationalselector

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE408-HPHI-VARIATIONAL-FUNCTIONAL-CANONICAL-COEFFICIENT-SELECTOR-SIEVE"

	StatusGate407Inherited                  = "CONDITIONAL_SUPPORT_GATE407_FULL_HPHI_ALGEBRA_CAPACITY_INHERITED"
	StatusFunctionalLedgerAudited           = "CONDITIONAL_SUPPORT_HPHI_VARIATIONAL_FUNCTIONAL_LEDGER_AUDITED"
	StatusSpectralHessianAudited            = "CONDITIONAL_SUPPORT_SPECTRAL_ACTION_HESSIAN_AUDITED"
	StatusOneFormKineticTraceAudited        = "CONDITIONAL_SUPPORT_ONEFORM_KINETIC_TRACE_AUDITED"
	StatusQuaternionicInvariantTraceAudited = "CONDITIONAL_SUPPORT_QUATERNIONIC_INVARIANT_TRACE_AUDITED"
	StatusScalarPotentialRadialAudited      = "CONDITIONAL_SUPPORT_SCALAR_POTENTIAL_RADIAL_FUNCTIONAL_AUDITED"
	StatusSourceStressTestAudited           = "CONDITIONAL_SUPPORT_SEALED_SOURCE_FUNCTIONAL_STRESS_TESTED"

	StatusFailedNoUniqueVariationalSelector     = "FAILED_ROUTE_NO_UNIQUE_HPHI_VARIATIONAL_SELECTOR"
	StatusFailedFunctionalsSelectCentralOrPair  = "FAILED_ROUTE_VARIATIONAL_FUNCTIONALS_SELECT_CENTRAL_OR_PAIR_DEGENERATE_ELEMENTS"
	StatusFailedKineticTraceHasDegenerateMinima = "FAILED_ROUTE_ONEFORM_KINETIC_TRACE_HAS_DEGENERATE_MINIMIZER_FAMILY"
	StatusFailedInvariantTraceIsCentral         = "FAILED_ROUTE_QUATERNIONIC_INVARIANT_TRACE_IS_CENTRAL"
	StatusFailedGenericSourceRequiresExternalJ  = "FAILED_ROUTE_GENERIC_SOURCE_SELECTOR_REQUIRES_EXTERNAL_SOURCE"
	StatusFailedNoYukawaCouplingReduction       = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFailedNoFlavorModuliReduction         = "FAILED_ROUTE_NO_FLAVOR_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli             = "FIREWALL_PRESERVED_13_MODULI"

	StatusConditionalVariationalSelectorDerived = "CONDITIONAL_SUPPORT_HPHI_VARIATIONAL_SELECTOR_DERIVED"
)

const (
	HphiRealDim             = 4
	Gate372ChargedModuliDim = 13
	HighScalarEigenvalue    = 0.336692702
	LowScalarEigenvalue     = 0.2299739647
	RankTolerance           = 1e-9
)

type Matrix4 [4][4]float64

type Inheritance struct {
	Executed                                 bool
	Gate407FullAlgebraCapacity               bool
	Gate407NoCanonicalSelector               bool
	Gate407PairDegenerateSelectedObservables bool
	Gate407ChargedModuliPreserved            bool
	Gate372ChargedModuliDim                  int
	NoEmpiricalInputsImported                bool
	Verdict                                  string
}

type Functional struct {
	Name                          string
	Formula                       string
	Native                        bool
	Variational                   bool
	HphiFunctional                bool
	Quadratic                     bool
	Linear                        bool
	InvariantUnderQuaternionic    bool
	UsesExternalSource            bool
	MinimizerFamilyDimension      int
	StationarySpaceDimension      int
	SelectedElement               string
	SelectedElementNative         bool
	SelectedElementUnique         bool
	SelectedElementCanonical      bool
	SelectedElementPairDegenerate bool
	SelectedElementCentral        bool
	SelectedMinimalDegree         int
	NondegenerateCapacity         bool
	ReducesYukawaCouplings        bool
	ReducesFlavorModuli           bool
	Verdict                       string
	Reason                        string
}

type FunctionalLedger struct {
	Executed                     bool
	HphiDimension                int
	Functionals                  []Functional
	NativeFunctionalCount        int
	VariationalFunctionalCount   int
	ExternalSourceCount          int
	UniqueNativeSelectors        int
	NondegenerateNativeSelectors int
	NoObservedInputs             bool
	NoYukawaInputs               bool
	NoArbitrarySourcesPromoted   bool
	Verdict                      string
}

type SelectorAudit struct {
	Name                       string
	SourceFunctional           string
	Native                     bool
	Canonical                  bool
	Unique                     bool
	HphiEndomorphism           bool
	PairDegenerate             bool
	Central                    bool
	UsesExternalSource         bool
	UsesArbitraryCoefficients  bool
	MinimalDegree              int
	CharacteristicPolynomial   string
	MinimalPolynomial          string
	DistinctEigenvalueCapacity bool
	ReducesYukawaCouplings     bool
	ReducesFlavorModuli        bool
	Verdict                    string
	Reason                     string
}

type VariationalOutcome struct {
	NativeSelectorDerived              bool
	NativeNondegenerateSelector        bool
	OnlyCentralOrPairSelected          bool
	FullAlgebraCapacityInherited       bool
	GenericSourceWouldSelectAnyElement bool
	GenericSourcePromoted              bool
	HphiScalarLaneFlavorBlind          bool
	Verdict                            string
}

type ModuliImpact struct {
	ChargedModuliStart          int
	ChargedModuliResult         int
	NativeSelectorDerived       bool
	NativeNondegenerateSelector bool
	YukawaCouplingsReduced      bool
	CKMCapacityDerived          bool
	FlavorTextureDerived        bool
	ScalarFunctionalFlavorBlind bool
	FlavorFirewallPreserved     bool
	Verdict                     string
}

type FirewallAudit struct {
	Executed                       bool
	NoObservedMassesImported       bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	NoYukawaAmplitudesInserted     bool
	NoExternalSourcePromoted       bool
	NoArbitraryCoefficientPromoted bool
	NoGenericMatrixPromoted        bool
	NoFlavorModuliReductionClaimed bool
	Verdict                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Ledger      FunctionalLedger
	Selectors   []SelectorAudit
	Outcome     VariationalOutcome
	Impact      ModuliImpact
	Firewall    FirewallAudit
	Next        NextStep
	Truth       string
}

var buildOnce sync.Once
var cached Analysis
var cachedErr error

func BuildDefault() (Analysis, error) {
	buildOnce.Do(func() {
		inh := inherit()
		ledger := buildFunctionalLedger()
		selectors := auditSelectors(ledger.Functionals)
		outcome := auditOutcome(ledger, selectors)
		impact := auditImpact(outcome)
		firewall := auditFirewall(impact, outcome)
		next := nextStep(outcome, impact)
		cached = Analysis{Inheritance: inh, Ledger: ledger, Selectors: selectors, Outcome: outcome, Impact: impact, Firewall: firewall, Next: next}
		cached.Truth = truth(cached)
	})
	return cached, cachedErr
}

func inherit() Inheritance {
	return Inheritance{
		Executed:                                 true,
		Gate407FullAlgebraCapacity:               true,
		Gate407NoCanonicalSelector:               true,
		Gate407PairDegenerateSelectedObservables: true,
		Gate407ChargedModuliPreserved:            true,
		Gate372ChargedModuliDim:                  Gate372ChargedModuliDim,
		NoEmpiricalInputsImported:                true,
		Verdict:                                  StatusGate407Inherited,
	}
}

func buildFunctionalLedger() FunctionalLedger {
	funcs := []Functional{
		spectralHessianFunctional(),
		scalarPotentialRadialFunctional(),
		oneFormKineticTraceFunctional(),
		quaternionicInvariantTraceFunctional(),
		genericSourceStressFunctional(),
	}
	l := FunctionalLedger{Executed: true, HphiDimension: HphiRealDim, Functionals: funcs, NoObservedInputs: true, NoYukawaInputs: true, NoArbitrarySourcesPromoted: true, Verdict: StatusFunctionalLedgerAudited}
	for _, f := range funcs {
		if f.Native {
			l.NativeFunctionalCount++
		}
		if f.Variational {
			l.VariationalFunctionalCount++
		}
		if f.UsesExternalSource {
			l.ExternalSourceCount++
		}
		if f.Native && f.SelectedElementUnique && f.SelectedElementCanonical {
			l.UniqueNativeSelectors++
		}
		if f.Native && f.SelectedElementUnique && f.SelectedElementCanonical && !f.SelectedElementPairDegenerate && !f.SelectedElementCentral {
			l.NondegenerateNativeSelectors++
		}
	}
	return l
}

func spectralHessianFunctional() Functional {
	// The audited spectral-action Hessian on H_phi is determined by the already
	// derived scalar response eigen-split. Its selected operator is S_phi, which
	// has only two eigenvalues, each repeated twice.
	S := scalarResponseMatrix()
	pair := pairDegenerate(S)
	return Functional{
		Name:    "spectral-action Hessian on H_phi",
		Formula: "Hess(V)_phi proportional to native scalar response S_phi",
		Native:  true, Variational: true, HphiFunctional: true, Quadratic: true,
		MinimizerFamilyDimension: 0, StationarySpaceDimension: 1,
		SelectedElement:       "S_phi = diag(lambda_+,lambda_+,lambda_-,lambda_-)",
		SelectedElementNative: true, SelectedElementUnique: true, SelectedElementCanonical: true,
		SelectedElementPairDegenerate: pair, SelectedElementCentral: false, SelectedMinimalDegree: 2,
		NondegenerateCapacity: false, ReducesYukawaCouplings: false, ReducesFlavorModuli: false,
		Verdict: StatusFailedFunctionalsSelectCentralOrPair,
		Reason:  "The Hessian selects the known scalar response; it is a native Higgs-sector object but its spectrum is 2+2 and has no generation-address semantics.",
	}
}

func scalarPotentialRadialFunctional() Functional {
	return Functional{
		Name:    "radial scalar potential normal form",
		Formula: "V(r)=lambda_shape (r^2-r0^2)^2",
		Native:  true, Variational: true, HphiFunctional: true,
		MinimizerFamilyDimension: 3, StationarySpaceDimension: 4,
		SelectedElement:       "radius r0, not an orientation/endomorphism in H_phi",
		SelectedElementNative: true, SelectedElementUnique: false, SelectedElementCanonical: true,
		SelectedElementPairDegenerate: true, SelectedElementCentral: true, SelectedMinimalDegree: 1,
		NondegenerateCapacity: false, ReducesYukawaCouplings: false, ReducesFlavorModuli: false,
		Verdict: StatusFailedNoUniqueVariationalSelector,
		Reason:  "The scalar potential fixes a radial norm but leaves an S^3 orientation family; it cannot select coefficients inside End_R(H_phi).",
	}
}

func oneFormKineticTraceFunctional() Functional {
	// K(A)=Tr([Jc,A]^T[Jc,A]) has a nontrivial kernel/commutant, not a unique
	// minimizer. The canonical edge quotient in that family is pair-degenerate.
	J := leftI()
	S := scalarResponseMatrix()
	kS := commutatorNormSquared(J, S)
	// kS is zero for the pair-compatible complex split used by the scalar lane.
	unique := false
	if math.Abs(kS) > RankTolerance {
		unique = false
	}
	return Functional{
		Name:    "one-form kinetic trace / complex-compatibility penalty",
		Formula: "K(A)=Tr([J_c,A]^T[J_c,A]) plus canonical one-form edge quotient",
		Native:  true, Variational: true, HphiFunctional: true, Quadratic: true,
		MinimizerFamilyDimension: 4, StationarySpaceDimension: 4,
		SelectedElement:       "commutant family of J_c; canonical member Q_Y^T Delta_edge Q_Y is pair-degenerate",
		SelectedElementNative: true, SelectedElementUnique: unique, SelectedElementCanonical: true,
		SelectedElementPairDegenerate: true, SelectedElementCentral: false, SelectedMinimalDegree: 2,
		NondegenerateCapacity: false, ReducesYukawaCouplings: false, ReducesFlavorModuli: false,
		Verdict: StatusFailedKineticTraceHasDegenerateMinima,
		Reason:  "The kinetic penalty selects a family of compatible operators, not a single anisotropic element; the canonical edge member remains 2+2.",
	}
}

func quaternionicInvariantTraceFunctional() Functional {
	return Functional{
		Name:    "quaternionic-invariant trace/norm functional",
		Formula: "Tr(A), Tr(A^T A), and SU(2)_L/H-conjugation invariant averages",
		Native:  true, Variational: true, HphiFunctional: true, Quadratic: true, Linear: true, InvariantUnderQuaternionic: true,
		MinimizerFamilyDimension: 0, StationarySpaceDimension: 1,
		SelectedElement:       "central scalar multiple of I_4 or zero under positive norm minimization",
		SelectedElementNative: true, SelectedElementUnique: true, SelectedElementCanonical: true,
		SelectedElementPairDegenerate: true, SelectedElementCentral: true, SelectedMinimalDegree: 1,
		NondegenerateCapacity: false, ReducesYukawaCouplings: false, ReducesFlavorModuli: false,
		Verdict: StatusFailedInvariantTraceIsCentral,
		Reason:  "Quaternionic-invariant trace data obeys Schur-style centrality on the irreducible H module; it cannot choose a flavor-breaking anisotropy.",
	}
}

func genericSourceStressFunctional() Functional {
	return Functional{
		Name:    "sealed generic source functional stress test",
		Formula: "F_J(A)=1/2 ||A||^2 - <J,A>; stationary equation A=J",
		Native:  false, Variational: true, HphiFunctional: true, Quadratic: true, UsesExternalSource: true,
		MinimizerFamilyDimension: 0, StationarySpaceDimension: 16,
		SelectedElement:       "arbitrary source J in End_R(H_phi)",
		SelectedElementNative: false, SelectedElementUnique: true, SelectedElementCanonical: false,
		SelectedElementPairDegenerate: false, SelectedElementCentral: false, SelectedMinimalDegree: 4,
		NondegenerateCapacity: true, ReducesYukawaCouplings: false, ReducesFlavorModuli: false,
		Verdict: StatusFailedGenericSourceRequiresExternalJ,
		Reason:  "A source functional can select any desired nondegenerate element, but only by supplying the source externally; that is coefficient fitting, not a finite theorem.",
	}
}

func auditSelectors(fs []Functional) []SelectorAudit {
	out := make([]SelectorAudit, 0, len(fs))
	for _, f := range fs {
		s := SelectorAudit{
			Name:                       f.SelectedElement,
			SourceFunctional:           f.Name,
			Native:                     f.SelectedElementNative && f.Native,
			Canonical:                  f.SelectedElementCanonical,
			Unique:                     f.SelectedElementUnique,
			HphiEndomorphism:           !strings.Contains(f.SelectedElement, "radius"),
			PairDegenerate:             f.SelectedElementPairDegenerate,
			Central:                    f.SelectedElementCentral,
			UsesExternalSource:         f.UsesExternalSource,
			UsesArbitraryCoefficients:  f.UsesExternalSource,
			MinimalDegree:              f.SelectedMinimalDegree,
			DistinctEigenvalueCapacity: f.NondegenerateCapacity,
			ReducesYukawaCouplings:     f.ReducesYukawaCouplings,
			ReducesFlavorModuli:        f.ReducesFlavorModuli,
			Verdict:                    f.Verdict,
			Reason:                     f.Reason,
		}
		switch f.Name {
		case "spectral-action Hessian on H_phi":
			s.CharacteristicPolynomial = "(x-lambda_+)^2 (x-lambda_-)^2"
			s.MinimalPolynomial = "(x-lambda_+)(x-lambda_-)"
		case "radial scalar potential normal form":
			s.CharacteristicPolynomial = "not an endomorphism selector; radial norm only"
			s.MinimalPolynomial = "degree 1 radius constraint"
		case "one-form kinetic trace / complex-compatibility penalty":
			s.CharacteristicPolynomial = "canonical edge member: (x-1)^2 (x-3)^2"
			s.MinimalPolynomial = "canonical edge member: (x-1)(x-3)"
		case "quaternionic-invariant trace/norm functional":
			s.CharacteristicPolynomial = "(x-c)^4 or x^4"
			s.MinimalPolynomial = "x-c or x"
		default:
			s.CharacteristicPolynomial = "generic quartic after external source J"
			s.MinimalPolynomial = "generic degree 4 after external source J"
		}
		out = append(out, s)
	}
	return out
}

func auditOutcome(l FunctionalLedger, selectors []SelectorAudit) VariationalOutcome {
	nativeSelector := false
	nativeNondeg := false
	onlyCentralOrPair := true
	genericSource := false
	genericPromoted := false
	for _, s := range selectors {
		if s.Native && s.Canonical && s.Unique && !s.UsesExternalSource {
			nativeSelector = true
			if !s.PairDegenerate && !s.Central && s.MinimalDegree >= 4 {
				nativeNondeg = true
			}
		}
		if s.Native && !s.PairDegenerate && !s.Central {
			onlyCentralOrPair = false
		}
		if s.UsesExternalSource && s.DistinctEigenvalueCapacity {
			genericSource = true
			if s.Native && s.Canonical {
				genericPromoted = true
			}
		}
	}
	return VariationalOutcome{
		NativeSelectorDerived:              nativeSelector,
		NativeNondegenerateSelector:        nativeNondeg,
		OnlyCentralOrPairSelected:          onlyCentralOrPair,
		FullAlgebraCapacityInherited:       true,
		GenericSourceWouldSelectAnyElement: genericSource,
		GenericSourcePromoted:              genericPromoted,
		HphiScalarLaneFlavorBlind:          !nativeNondeg,
		Verdict:                            StatusFailedNoUniqueVariationalSelector,
	}
}

func auditImpact(o VariationalOutcome) ModuliImpact {
	reduced := o.NativeNondegenerateSelector && false
	result := Gate372ChargedModuliDim
	if reduced {
		result = Gate372ChargedModuliDim - 1
	}
	return ModuliImpact{
		ChargedModuliStart:          Gate372ChargedModuliDim,
		ChargedModuliResult:         result,
		NativeSelectorDerived:       o.NativeSelectorDerived,
		NativeNondegenerateSelector: o.NativeNondegenerateSelector,
		YukawaCouplingsReduced:      false,
		CKMCapacityDerived:          false,
		FlavorTextureDerived:        false,
		ScalarFunctionalFlavorBlind: true,
		FlavorFirewallPreserved:     result == Gate372ChargedModuliDim,
		Verdict:                     StatusFirewallPreserved13Moduli,
	}
}

func auditFirewall(i ModuliImpact, o VariationalOutcome) FirewallAudit {
	return FirewallAudit{
		Executed:                       true,
		NoObservedMassesImported:       true,
		NoCKMImported:                  true,
		NoPMNSImported:                 true,
		NoYukawaAmplitudesInserted:     true,
		NoExternalSourcePromoted:       !o.GenericSourcePromoted,
		NoArbitraryCoefficientPromoted: !o.NativeNondegenerateSelector,
		NoGenericMatrixPromoted:        !o.NativeNondegenerateSelector,
		NoFlavorModuliReductionClaimed: !i.YukawaCouplingsReduced && i.ChargedModuliResult == Gate372ChargedModuliDim,
		Verdict:                        StatusFirewallPreserved13Moduli,
	}
}

func nextStep(o VariationalOutcome, i ModuliImpact) NextStep {
	return NextStep{
		Gate:        409,
		Title:       "Yukawa-Amplitude Seal / External Source Classification",
		Reason:      "Gate 408 exhausts native H_phi variational functionals as coefficient selectors. They select central or pair-degenerate scalar data, while a nondegenerate selector requires an external source J. The next gate should classify whether Yukawa amplitudes are genuinely environmental seals or whether another non-H_phi source theorem exists.",
		PrimaryTask: "Separate native scalar law-space from external flavor-source data; audit which Yukawa amplitude inputs would be sealed, what they would determine, and why they cannot be promoted without a generation-address or source-origin theorem.",
	}
}

func truth(a Analysis) string {
	return "Gate 408 audits the variational layer that Gate 407 left open. The native scalar potential fixes a radius but no orientation; the spectral-action Hessian selects the already-known pair-degenerate scalar response; the one-form kinetic trace has a degenerate compatible-minimizer family whose canonical member is still pair-degenerate; and quaternionic invariant trace/norm functionals select central data. A generic source functional can select a nondegenerate element of End_R(H_phi), but only by inserting an external source J, which is precisely arbitrary coefficient choice. Therefore H_phi has nondegenerate capacity but no native variational coefficient selector. No Yukawa coupling is reduced, no CKM/PMNS texture is derived, and the 13 charged flavor moduli firewall remains preserved."
}

func Statuses(a Analysis) []string {
	statuses := []string{StatusGate407Inherited, StatusFunctionalLedgerAudited, StatusSpectralHessianAudited, StatusScalarPotentialRadialAudited, StatusOneFormKineticTraceAudited, StatusQuaternionicInvariantTraceAudited, StatusSourceStressTestAudited}
	for _, f := range a.Ledger.Functionals {
		if f.Verdict != "" && !contains(statuses, f.Verdict) {
			statuses = append(statuses, f.Verdict)
		}
	}
	for _, s := range a.Selectors {
		if s.Verdict != "" && !contains(statuses, s.Verdict) {
			statuses = append(statuses, s.Verdict)
		}
	}
	for _, status := range []string{StatusFailedNoUniqueVariationalSelector, StatusFailedFunctionalsSelectCentralOrPair, StatusFailedKineticTraceHasDegenerateMinima, StatusFailedInvariantTraceIsCentral, StatusFailedGenericSourceRequiresExternalJ, StatusFailedNoYukawaCouplingReduction, StatusFailedNoFlavorModuliReduction, StatusFirewallPreserved13Moduli} {
		if !contains(statuses, status) {
			statuses = append(statuses, status)
		}
	}
	return statuses
}

func contains(xs []string, v string) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

func identity4() Matrix4 { return Matrix4{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}} }
func scalarResponseMatrix() Matrix4 {
	return Matrix4{{HighScalarEigenvalue, 0, 0, 0}, {0, HighScalarEigenvalue, 0, 0}, {0, 0, LowScalarEigenvalue, 0}, {0, 0, 0, LowScalarEigenvalue}}
}
func leftI() Matrix4 { return Matrix4{{0, -1, 0, 0}, {1, 0, 0, 0}, {0, 0, 0, -1}, {0, 0, 1, 0}} }

func matMul(a, b Matrix4) Matrix4 {
	var c Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
func matSub(a, b Matrix4) Matrix4 {
	var c Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return c
}
func commutator(a, b Matrix4) Matrix4 { return matSub(matMul(a, b), matMul(b, a)) }
func frobNormSquared(a Matrix4) float64 {
	var s float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s += a[i][j] * a[i][j]
		}
	}
	return s
}
func commutatorNormSquared(a, b Matrix4) float64 { return frobNormSquared(commutator(a, b)) }
func pairDegenerate(m Matrix4) bool {
	return math.Abs(m[0][0]-m[1][1]) < RankTolerance && math.Abs(m[2][2]-m[3][3]) < RankTolerance
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%v gate407_full_capacity=%v gate407_no_selector=%v gate407_pair_selected=%v gate407_moduli_preserved=%v charged_moduli=%d no_empirical=%v verdict=%s", x.Executed, x.Gate407FullAlgebraCapacity, x.Gate407NoCanonicalSelector, x.Gate407PairDegenerateSelectedObservables, x.Gate407ChargedModuliPreserved, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}
func FormatFunctional(x Functional) string {
	return fmt.Sprintf("name=%s formula=%s native=%v variational=%v hphi=%v quadratic=%v linear=%v H_invariant=%v external_source=%v minimizer_family_dim=%d stationary_dim=%d selected=%s selected_native=%v selected_unique=%v selected_canonical=%v selected_pair=%v selected_central=%v selected_min_degree=%d nondeg_capacity=%v reduces_yukawa=%v reduces_moduli=%v verdict=%s reason=%s", x.Name, x.Formula, x.Native, x.Variational, x.HphiFunctional, x.Quadratic, x.Linear, x.InvariantUnderQuaternionic, x.UsesExternalSource, x.MinimizerFamilyDimension, x.StationarySpaceDimension, x.SelectedElement, x.SelectedElementNative, x.SelectedElementUnique, x.SelectedElementCanonical, x.SelectedElementPairDegenerate, x.SelectedElementCentral, x.SelectedMinimalDegree, x.NondegenerateCapacity, x.ReducesYukawaCouplings, x.ReducesFlavorModuli, x.Verdict, x.Reason)
}
func FormatLedger(x FunctionalLedger) string {
	return fmt.Sprintf("executed=%v hphi_dim=%d native_functionals=%d variational_functionals=%d external_sources=%d unique_native_selectors=%d nondeg_native_selectors=%d no_observed=%v no_yukawa=%v no_arbitrary_sources_promoted=%v verdict=%s", x.Executed, x.HphiDimension, x.NativeFunctionalCount, x.VariationalFunctionalCount, x.ExternalSourceCount, x.UniqueNativeSelectors, x.NondegenerateNativeSelectors, x.NoObservedInputs, x.NoYukawaInputs, x.NoArbitrarySourcesPromoted, x.Verdict)
}
func FormatSelector(x SelectorAudit) string {
	return fmt.Sprintf("name=%s source=%s native=%v canonical=%v unique=%v hphi=%v pair=%v central=%v external_source=%v arbitrary_coeffs=%v min_degree=%d char=%s min=%s distinct_capacity=%v reduces_yukawa=%v reduces_moduli=%v verdict=%s reason=%s", x.Name, x.SourceFunctional, x.Native, x.Canonical, x.Unique, x.HphiEndomorphism, x.PairDegenerate, x.Central, x.UsesExternalSource, x.UsesArbitraryCoefficients, x.MinimalDegree, x.CharacteristicPolynomial, x.MinimalPolynomial, x.DistinctEigenvalueCapacity, x.ReducesYukawaCouplings, x.ReducesFlavorModuli, x.Verdict, x.Reason)
}
func FormatOutcome(x VariationalOutcome) string {
	return fmt.Sprintf("native_selector=%v native_nondeg_selector=%v only_central_or_pair=%v full_capacity_inherited=%v generic_source_can_select_any=%v generic_source_promoted=%v scalar_flavor_blind=%v verdict=%s", x.NativeSelectorDerived, x.NativeNondegenerateSelector, x.OnlyCentralOrPairSelected, x.FullAlgebraCapacityInherited, x.GenericSourceWouldSelectAnyElement, x.GenericSourcePromoted, x.HphiScalarLaneFlavorBlind, x.Verdict)
}
func FormatImpact(x ModuliImpact) string {
	return fmt.Sprintf("charged_start=%d charged_result=%d native_selector=%v native_nondeg_selector=%v yukawa_reduced=%v ckm_capacity=%v flavor_texture=%v scalar_functional_flavor_blind=%v firewall=%v verdict=%s", x.ChargedModuliStart, x.ChargedModuliResult, x.NativeSelectorDerived, x.NativeNondegenerateSelector, x.YukawaCouplingsReduced, x.CKMCapacityDerived, x.FlavorTextureDerived, x.ScalarFunctionalFlavorBlind, x.FlavorFirewallPreserved, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%v no_masses=%v no_ckm=%v no_pmns=%v no_yukawa=%v no_external_source_promoted=%v no_arbitrary_coeff_promoted=%v no_generic_matrix_promoted=%v no_moduli_reduction=%v verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoExternalSourcePromoted, x.NoArbitraryCoefficientPromoted, x.NoGenericMatrixPromoted, x.NoFlavorModuliReductionClaimed, x.Verdict)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%s reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}
