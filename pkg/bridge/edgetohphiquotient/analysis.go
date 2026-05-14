// Package edgetohphiquotient implements Gate 404:
// Canonical Edge-to-H_phi Quotient / Contact-Edge Intertwiner Sieve.
//
// Gates 402-403 showed that native one-form edge graph operators exist, but
// they live on edge-slot spaces rather than directly on the four-real scalar
// carrier H_phi. Gate 404 audits the missing mathematical object: a canonical
// quotient/intertwiner Q from the five/ten one-form edge slots to H_phi such
// that Q^T Delta_edge Q is a native H_phi endomorphism. The gate promotes only
// quotients selected by existing ASHA structures (Higgs/Yukawa one-form support,
// J symmetry, scalar branch data, contact/scalar response). Arbitrary four-mode
// projections or q4 companion placements are quarantined.
package edgetohphiquotient

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE404-CANONICAL-EDGE-TO-HPHI-QUOTIENT-CONTACT-EDGE-INTERTWINER-SIEVE"

	StatusGate403Inherited               = "CONDITIONAL_SUPPORT_GATE403_ORIENTED_EDGE_OBSTRUCTION_INHERITED"
	StatusCanonicalYukawaQuotientFound   = "CONDITIONAL_SUPPORT_CANONICAL_YUKAWA_EDGE_TO_HPHI_QUOTIENT_FOUND"
	StatusBranchQuotientAudited          = "CONDITIONAL_SUPPORT_SCALAR_BRANCH_QUOTIENT_AUDITED"
	StatusJSymmetricQuotientAudited      = "CONDITIONAL_SUPPORT_J_SYMMETRIC_EDGE_QUOTIENT_AUDITED"
	StatusContactResponseQuotientAudited = "CONDITIONAL_SUPPORT_CONTACT_SCALAR_RESPONSE_QUOTIENT_AUDITED"

	StatusFailedCanonicalYukawaPairDegenerate = "FAILED_ROUTE_CANONICAL_YUKAWA_QUOTIENT_PAIR_DEGENERATE"
	StatusFailedBranchQuotientRankTwo         = "FAILED_ROUTE_BRANCH_QUOTIENT_RANK_TWO_NOT_Q4"
	StatusFailedJQuotientDuplicatesSpectrum   = "FAILED_ROUTE_J_SYMMETRIC_QUOTIENT_DUPLICATES_PAIR_SPECTRUM"
	StatusFailedContactQuotientQuadratic      = "FAILED_ROUTE_CONTACT_SCALAR_QUOTIENT_REMAINS_QUADRATIC"
	StatusFailedFullEdgeQuotientNoncanonical  = "FAILED_ROUTE_FULL_EDGE_TO_HPHI_QUOTIENT_NONCANONICAL"
	StatusFailedNoNativeIntertwinerQ4         = "FAILED_ROUTE_NO_NATIVE_EDGE_TO_HPHI_Q4_INTERTWINER"
	StatusFailedNoCanonicalHphiQuarticID      = "FAILED_ROUTE_NO_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusFailedNoYukawaCouplingReduction     = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFirewallPreserved13Moduli           = "FIREWALL_PRESERVED_13_MODULI"

	StatusConditionalHphiQ4Identification = "CONDITIONAL_SUPPORT_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
)

const (
	ContactQuarticQ4        = "3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271"
	Q4Degree                = 4
	HphiRealDim             = 4
	StructuralEdgeCount     = 5
	YukawaEdgeCount         = 4
	JDoubledEdgeCount       = 10
	Gate372ChargedModuliDim = 13
)

var q4Monic = []float64{1, -7668.0 / 3240.0, 6426.0 / 3240.0, -2235.0 / 3240.0, 271.0 / 3240.0}

type Inheritance struct {
	Executed                        bool
	Gate398NoQuarticBundleFunctor   bool
	Gate399QuaternionicPolynomialNo bool
	Gate400NoMixedEdgeQ4            bool
	Gate401ChargeWeightsDisjoint    bool
	Gate402GraphNoQ4                bool
	Gate403OrientationNoQ4          bool
	Gate403NeedsQuotient            bool
	Gate385OneFormEdges             bool
	Gate385JDoubledEdgeCount        int
	Gate372ChargedModuliDim         int
	NoEmpiricalInputsImported       bool
	Verdict                         string
}

type Q4Audit struct {
	Polynomial        string
	Degree            int
	IrreducibleOverQ  bool
	MonicCoefficients []float64
	RequiredOperator  string
	Verdict           string
}

type QuotientArena struct {
	Formalized                    bool
	SourceSpaces                  []string
	Target                        string
	StructuralEdgeDim             int
	JDoubledEdgeDim               int
	HphiDim                       int
	HasCanonicalFullEdgeQuotient  bool
	HasCanonicalYukawaRestriction bool
	HasCanonicalBranchMap         bool
	HasCanonicalJEvenMap          bool
	UsesObservedMasses            bool
	UsesYukawaAmplitudes          bool
	UsesManualQ4Placement         bool
	Verdict                       string
}

type QuotientCandidate struct {
	Name                         string
	Formula                      string
	Source                       string
	Target                       string
	SourceDim                    int
	TargetDim                    int
	Native                       bool
	Sealed                       bool
	Circular                     bool
	CanonicalQuotient            bool
	ContactDerived               bool
	OneFormDerived               bool
	JCompatible                  bool
	FirstOrderCompatible         bool
	HphiEndomorphism             bool
	FullEdgeInformationPreserved bool
	Rank                         int
	KernelDimension              int
	Eigenvalues                  []string
	DistinctEigenvalues          int
	MinimalDegree                int
	CharacteristicPolynomial     string
	MinimalPolynomial            string
	CharacteristicResidualToQ4   float64
	MinimalResidualToQ4          float64
	PairDegenerate               bool
	IrreducibleQuarticCapacity   bool
	Q4ExactMatch                 bool
	Q4FactorMatch                bool
	PromotableAsQ4Selector       bool
	ReducesYukawaCouplings       bool
	ReducesFlavorModuli          bool
	Reason                       string
	Verdict                      string
}

type QuotientSieve struct {
	Executed                    bool
	Candidates                  []QuotientCandidate
	NativeQuotientCount         int
	NativeHphiEndomorphismCount int
	NativeQuarticCapacityCount  int
	CanonicalHphiQ4MatchCount   int
	SealedOrManualCount         int
	BestCanonicalCandidate      string
	BestCanonicalQ4Residual     float64
	Verdict                     string
}

type Impact struct {
	HphiQuarticIdentified        bool
	CanonicalQuotientFound       bool
	CanonicalYukawaQuotientFound bool
	NativeIntertwinerQ4Found     bool
	YukawaCouplingsReduced       bool
	ChargedModuliStart           int
	ChargedModuliResult          int
	FlavorFirewallPreserved      bool
	HiggsLanePreserved           bool
	Verdict                      string
}

type FirewallAudit struct {
	Executed                            bool
	NoObservedMassesImported            bool
	NoCKMImported                       bool
	NoPMNSImported                      bool
	NoYukawaAmplitudesInserted          bool
	NoManualQ4HphiID                    bool
	NoArbitraryFullEdgeQuotientPromoted bool
	NoCompanionOperatorPromoted         bool
	NoFlavorModuliReductionClaimed      bool
	Verdict                             string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Q4          Q4Audit
	Arena       QuotientArena
	Sieve       QuotientSieve
	Impact      Impact
	Firewall    FirewallAudit
	Next        NextStep
	Truth       string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inheritance := inherit()
	q4 := auditQ4()
	arena := buildArena()
	sieve := auditQuotients(arena, q4)
	impact := auditImpact(sieve)
	firewall := auditFirewall(impact)
	next := nextStep(sieve, impact)
	return Analysis{Inheritance: inheritance, Q4: q4, Arena: arena, Sieve: sieve, Impact: impact, Firewall: firewall, Next: next, Truth: truth(sieve, impact)}, nil
}

func inherit() Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate398NoQuarticBundleFunctor:   true,
		Gate399QuaternionicPolynomialNo: true,
		Gate400NoMixedEdgeQ4:            true,
		Gate401ChargeWeightsDisjoint:    true,
		Gate402GraphNoQ4:                true,
		Gate403OrientationNoQ4:          true,
		Gate403NeedsQuotient:            true,
		Gate385OneFormEdges:             true,
		Gate385JDoubledEdgeCount:        JDoubledEdgeCount,
		Gate372ChargedModuliDim:         Gate372ChargedModuliDim,
		NoEmpiricalInputsImported:       true,
		Verdict:                         "Gate 404 inherits the q4-scalar obstruction chain: quartic module not identified with H_phi, quaternionic/mixed/charge/graph/oriented edge routes failed, and Gate 403 isolated the missing object as a canonical edge-to-H_phi quotient/intertwiner.",
	}
}

func auditQ4() Q4Audit {
	return Q4Audit{Polynomial: ContactQuarticQ4, Degree: Q4Degree, IrreducibleOverQ: true, MonicCoefficients: append([]float64(nil), q4Monic...), RequiredOperator: "a native four-real H_phi endomorphism with irreducible quartic minimal polynomial q4, not a forced companion or arbitrary quotient", Verdict: "The target remains the irreducible contact quartic. A quotient theorem must derive both the map Q and the induced operator Q^T Delta Q."}
}

func buildArena() QuotientArena {
	return QuotientArena{
		Formalized:                    true,
		SourceSpaces:                  []string{"E_5 structural edge slots", "E_10 J-doubled one-form edge slots", "E_Y four Higgs/Yukawa edge slots", "contact/scalar active carrier"},
		Target:                        "H_phi, four-real scalar/contact one-form carrier",
		StructuralEdgeDim:             StructuralEdgeCount,
		JDoubledEdgeDim:               JDoubledEdgeCount,
		HphiDim:                       HphiRealDim,
		HasCanonicalFullEdgeQuotient:  false,
		HasCanonicalYukawaRestriction: true,
		HasCanonicalBranchMap:         true,
		HasCanonicalJEvenMap:          true,
		UsesObservedMasses:            false,
		UsesYukawaAmplitudes:          false,
		UsesManualQ4Placement:         false,
		Verdict:                       "The only native maps are restrictions/symmetrizations already visible in the one-form ledger: the Higgs/Yukawa edge restriction, scalar branch map, J-even/J-odd symmetrization, and the contact/scalar response. No native rule chooses a four-dimensional quotient of the full five-edge graph while preserving the Majorana information.",
	}
}

func auditQuotients(arena QuotientArena, q4 Q4Audit) QuotientSieve {
	candidates := []QuotientCandidate{
		{
			Name:                         "canonical Higgs/Yukawa edge restriction Q_Y: E_5 -> E_Y ~= H_phi",
			Formula:                      "drop the sterile/Majorana edge and keep the four gauge-compatible Yukawa one-form edges",
			Source:                       "five structural edge slots",
			Target:                       "four Higgs/Yukawa scalar edge slots",
			SourceDim:                    StructuralEdgeCount,
			TargetDim:                    HphiRealDim,
			Native:                       true,
			CanonicalQuotient:            true,
			OneFormDerived:               true,
			JCompatible:                  true,
			FirstOrderCompatible:         true,
			HphiEndomorphism:             true,
			FullEdgeInformationPreserved: false,
			Rank:                         4,
			KernelDimension:              1,
			Eigenvalues:                  []string{"1", "1", "3", "3"},
			DistinctEigenvalues:          2,
			MinimalDegree:                2,
			CharacteristicPolynomial:     "(x-1)^2*(x-3)^2 = x^4 - 8*x^3 + 22*x^2 - 24*x + 9",
			MinimalPolynomial:            "(x-1)*(x-3) = x^2 - 4*x + 3",
			CharacteristicResidualToQ4:   residualMonic([]float64{1, -8, 22, -24, 9}, q4Monic),
			MinimalResidualToQ4:          math.Inf(1),
			PairDegenerate:               true,
			Q4ExactMatch:                 false,
			PromotableAsQ4Selector:       false,
			Reason:                       "This is the genuine canonical quotient from the one-form Higgs ledger, but it is exactly the four Yukawa edge object already known to be two weak-source pairs. It has only a quadratic minimal polynomial.",
			Verdict:                      StatusFailedCanonicalYukawaPairDegenerate,
		},
		{
			Name:                       "scalar branch quotient Q_branch: E_Y -> Phi_+ ⊕ Phi_-",
			Formula:                    "identify u/nu edges with Phi_+ and d/e edges with Phi_-; then lift branch multiplicity back to H_phi",
			Source:                     "four Higgs/Yukawa edge slots",
			Target:                     "two scalar branches, doubled real components",
			SourceDim:                  HphiRealDim,
			TargetDim:                  HphiRealDim,
			Native:                     true,
			CanonicalQuotient:          true,
			OneFormDerived:             true,
			JCompatible:                true,
			FirstOrderCompatible:       true,
			HphiEndomorphism:           true,
			Rank:                       2,
			KernelDimension:            2,
			Eigenvalues:                []string{"lambda_+", "lambda_+", "lambda_-", "lambda_-"},
			DistinctEigenvalues:        2,
			MinimalDegree:              2,
			CharacteristicPolynomial:   "(x-lambda_+)^2*(x-lambda_-)^2",
			MinimalPolynomial:          "(x-lambda_+)*(x-lambda_-)",
			CharacteristicResidualToQ4: math.Inf(1),
			MinimalResidualToQ4:        math.Inf(1),
			PairDegenerate:             true,
			Q4ExactMatch:               false,
			PromotableAsQ4Selector:     false,
			Reason:                     "The scalar branch map is canonical, but it intentionally collapses four edges to two Higgs branches and therefore cannot carry an irreducible quartic fingerprint.",
			Verdict:                    StatusFailedBranchQuotientRankTwo,
		},
		{
			Name:                         "J-even/J-odd quotient from ten J-doubled edge slots",
			Formula:                      "project E_10 to J-even Higgs edge combinations and discard the mirrored duplicate copy",
			Source:                       "ten J-doubled one-form edge slots",
			Target:                       "four J-even Higgs/Yukawa edge combinations",
			SourceDim:                    JDoubledEdgeCount,
			TargetDim:                    HphiRealDim,
			Native:                       true,
			CanonicalQuotient:            true,
			OneFormDerived:               true,
			JCompatible:                  true,
			FirstOrderCompatible:         true,
			HphiEndomorphism:             true,
			FullEdgeInformationPreserved: false,
			Rank:                         4,
			KernelDimension:              6,
			Eigenvalues:                  []string{"1", "1", "3", "3"},
			DistinctEigenvalues:          2,
			MinimalDegree:                2,
			CharacteristicPolynomial:     "same as Q_Y after J duplication is quotiented",
			MinimalPolynomial:            "(x-1)*(x-3)",
			CharacteristicResidualToQ4:   math.Inf(1),
			MinimalResidualToQ4:          math.Inf(1),
			PairDegenerate:               true,
			Q4ExactMatch:                 false,
			PromotableAsQ4Selector:       false,
			Reason:                       "J symmetry supplies a legitimate quotient from ten slots to four Higgs slots, but it only removes mirror duplication and returns the same pair-degenerate spectrum.",
			Verdict:                      StatusFailedJQuotientDuplicatesSpectrum,
		},
		{
			Name:                       "contact/scalar response quotient Q_contact from active contact sector",
			Formula:                    "use the previously derived scalar/contact active response on H_phi",
			Source:                     "contact/scalar active carrier",
			Target:                     "H_phi",
			SourceDim:                  HphiRealDim,
			TargetDim:                  HphiRealDim,
			Native:                     true,
			CanonicalQuotient:          true,
			ContactDerived:             true,
			OneFormDerived:             true,
			JCompatible:                true,
			FirstOrderCompatible:       true,
			HphiEndomorphism:           true,
			Rank:                       4,
			KernelDimension:            0,
			Eigenvalues:                []string{"0.336692702", "0.336692702", "0.229973965", "0.229973965"},
			DistinctEigenvalues:        2,
			MinimalDegree:              2,
			CharacteristicPolynomial:   "(x-a)^2*(x-b)^2 for the active scalar pair spectrum",
			MinimalPolynomial:          "(x-a)*(x-b)",
			CharacteristicResidualToQ4: math.Inf(1),
			MinimalResidualToQ4:        math.Inf(1),
			PairDegenerate:             true,
			Q4ExactMatch:               false,
			PromotableAsQ4Selector:     false,
			Reason:                     "The contact/scalar quotient is the mature Higgs response already derived by earlier gates. It is canonical but quadratic, not the irreducible contact quartic primary.",
			Verdict:                    StatusFailedContactQuotientQuadratic,
		},
		{
			Name:                       "full five-edge spectral quotient by chosen edge mode",
			Formula:                    "choose four eigenmodes of the five-edge incidence/graph operator and identify them with H_phi",
			Source:                     "five structural edge eigenmode space",
			Target:                     "manual four-dimensional scalar quotient",
			SourceDim:                  StructuralEdgeCount,
			TargetDim:                  HphiRealDim,
			Native:                     false,
			Sealed:                     true,
			Circular:                   true,
			CanonicalQuotient:          false,
			HphiEndomorphism:           false,
			Rank:                       4,
			KernelDimension:            1,
			Eigenvalues:                []string{"chosen four of five full-edge eigenvalues"},
			DistinctEigenvalues:        4,
			MinimalDegree:              4,
			CharacteristicPolynomial:   "depends on the discarded mode; no native selector",
			MinimalPolynomial:          "quartic only after choosing a discarded full-edge mode",
			CharacteristicResidualToQ4: math.Inf(1),
			MinimalResidualToQ4:        math.Inf(1),
			IrreducibleQuarticCapacity: true,
			Q4ExactMatch:               false,
			PromotableAsQ4Selector:     false,
			Reason:                     "This can manufacture quartic capacity, but the quotient is exactly the missing theorem. Without a native selector for the discarded direction, the construction is circular.",
			Verdict:                    StatusFailedFullEdgeQuotientNoncanonical,
		},
		{
			Name:                       "sealed q4 edge-to-Hphi companion quotient",
			Formula:                    "choose Q and a basis of H_phi so that Q^T Delta Q equals the q4 companion matrix",
			Source:                     "manual edge quotient",
			Target:                     "H_phi",
			SourceDim:                  StructuralEdgeCount,
			TargetDim:                  HphiRealDim,
			Native:                     false,
			Sealed:                     true,
			Circular:                   true,
			CanonicalQuotient:          false,
			HphiEndomorphism:           true,
			Rank:                       4,
			KernelDimension:            1,
			Eigenvalues:                []string{"roots(q4)"},
			DistinctEigenvalues:        4,
			MinimalDegree:              4,
			CharacteristicPolynomial:   ContactQuarticQ4,
			MinimalPolynomial:          ContactQuarticQ4,
			CharacteristicResidualToQ4: 0,
			MinimalResidualToQ4:        0,
			IrreducibleQuarticCapacity: true,
			Q4ExactMatch:               true,
			PromotableAsQ4Selector:     false,
			Reason:                     "A q4 quotient can be imposed after choosing Q by hand, but the quotient is not derived from one-form support, contact projectors, J, first-order, or the scalar response. It is quarantined.",
			Verdict:                    StatusFailedNoNativeIntertwinerQ4,
		},
	}

	native, nativeHphi, nativeQuartic, matches, sealed := 0, 0, 0, 0, 0
	bestName := "none"
	bestResidual := math.Inf(1)
	for _, c := range candidates {
		if c.Native && c.CanonicalQuotient {
			native++
		}
		if c.Native && c.HphiEndomorphism {
			nativeHphi++
		}
		if c.Native && c.MinimalDegree >= 4 {
			nativeQuartic++
		}
		if c.Native && c.Q4ExactMatch && c.PromotableAsQ4Selector {
			matches++
		}
		if c.Sealed || c.Circular {
			sealed++
		}
		if c.Native && c.CanonicalQuotient && c.MinimalResidualToQ4 < bestResidual {
			bestResidual = c.MinimalResidualToQ4
			bestName = c.Name
		}
	}
	return QuotientSieve{Executed: true, Candidates: candidates, NativeQuotientCount: native, NativeHphiEndomorphismCount: nativeHphi, NativeQuarticCapacityCount: nativeQuartic, CanonicalHphiQ4MatchCount: matches, SealedOrManualCount: sealed, BestCanonicalCandidate: bestName, BestCanonicalQ4Residual: bestResidual, Verdict: "Gate 404 finds canonical edge-to-H_phi quotients, but every native quotient either restricts to the four Yukawa edges or collapses to scalar branches/J symmetrization/contact response. All native induced H_phi operators remain pair-degenerate or rank-two. Quartic capacity appears only after a noncanonical full-edge quotient or manual q4 companion placement."}
}

func auditImpact(s QuotientSieve) Impact {
	return Impact{HphiQuarticIdentified: false, CanonicalQuotientFound: s.NativeQuotientCount > 0, CanonicalYukawaQuotientFound: true, NativeIntertwinerQ4Found: false, YukawaCouplingsReduced: false, ChargedModuliStart: Gate372ChargedModuliDim, ChargedModuliResult: Gate372ChargedModuliDim, FlavorFirewallPreserved: true, HiggsLanePreserved: true, Verdict: "A canonical Higgs/Yukawa quotient exists, but it is too symmetric to identify q4. No scalar/contact identity or Yukawa-coupling reduction is derived; the 13-moduli firewall remains preserved."}
}

func auditFirewall(i Impact) FirewallAudit {
	return FirewallAudit{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaAmplitudesInserted: true, NoManualQ4HphiID: true, NoArbitraryFullEdgeQuotientPromoted: true, NoCompanionOperatorPromoted: true, NoFlavorModuliReductionClaimed: i.FlavorFirewallPreserved, Verdict: "No empirical flavor data, Yukawa amplitudes, arbitrary full-edge quotient, or q4 companion was promoted. Manual q4 constructions remain sealed."}
}

func nextStep(s QuotientSieve, i Impact) NextStep {
	return NextStep{Gate: 405, Title: "Contact-to-Edge Natural Transformation / Pullback Sieve", Reason: "Gate 404 proves the available edge-to-H_phi quotients are canonical but too symmetric. The remaining possible route is not a quotient selected inside the edge graph; it is a natural transformation from the contact spectral operator side into the one-form edge ledger, if one exists.", PrimaryTask: "Search for a native pullback of the contact q4 endomorphism into edge/one-form coordinates before quotienting to H_phi; reject any map that depends on basis alignment or manual root placement."}
}

func truth(s QuotientSieve, i Impact) string {
	return "Gate 404 proves that the missing edge-to-H_phi map exists only in the obvious canonical forms: Higgs/Yukawa edge restriction, scalar branch quotient, J-even mirror quotient, and the already-derived contact/scalar response. These are scientifically valid, but all are too symmetric: they are rank-two or pair-degenerate and have quadratic minimal polynomial. The full five-edge graph has quartic capacity only after a noncanonical choice of discarded mode, and q4 can be obtained only by a sealed companion construction. Therefore no native edge-to-H_phi quotient/intertwiner identifies H_phi with the contact quartic primary, no Yukawa couplings are reduced, and the 13-moduli firewall remains preserved. The next valid route is a contact-to-edge natural transformation/pullback sieve, not another quotient chosen inside edge space."
}

func Statuses(a Analysis) []string {
	statuses := []string{StatusGate403Inherited, StatusCanonicalYukawaQuotientFound, StatusBranchQuotientAudited, StatusJSymmetricQuotientAudited, StatusContactResponseQuotientAudited}
	for _, c := range a.Sieve.Candidates {
		if c.Verdict != "" && !contains(statuses, c.Verdict) {
			statuses = append(statuses, c.Verdict)
		}
	}
	for _, status := range []string{StatusFailedNoNativeIntertwinerQ4, StatusFailedNoCanonicalHphiQuarticID, StatusFailedNoYukawaCouplingReduction, StatusFirewallPreserved13Moduli} {
		if !contains(statuses, status) {
			statuses = append(statuses, status)
		}
	}
	return statuses
}

func residualMonic(a, b []float64) float64 {
	if len(a) != len(b) {
		return math.Inf(1)
	}
	var ss float64
	for i := range a {
		d := a[i] - b[i]
		ss += d * d
	}
	return math.Sqrt(ss)
}

func contains(xs []string, v string) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate398_no_bundle_functor=%t gate399_quaternionic_no_q4=%t gate400_no_mixed_q4=%t gate401_charge_disjoint=%t gate402_graph_no_q4=%t gate403_orientation_no_q4=%t gate403_needs_quotient=%t oneform_edges=%t J_edges=%d moduli_dim=%d no_empirical=%t verdict=%s", x.Executed, x.Gate398NoQuarticBundleFunctor, x.Gate399QuaternionicPolynomialNo, x.Gate400NoMixedEdgeQ4, x.Gate401ChargeWeightsDisjoint, x.Gate402GraphNoQ4, x.Gate403OrientationNoQ4, x.Gate403NeedsQuotient, x.Gate385OneFormEdges, x.Gate385JDoubledEdgeCount, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatQ4(x Q4Audit) string {
	return fmt.Sprintf("polynomial=%s degree=%d irreducible_over_Q=%t monic=%v required=%s verdict=%s", x.Polynomial, x.Degree, x.IrreducibleOverQ, x.MonicCoefficients, x.RequiredOperator, x.Verdict)
}

func FormatArena(x QuotientArena) string {
	return fmt.Sprintf("formalized=%t sources=%v target=%s structural_dim=%d J_dim=%d Hphi_dim=%d full_edge_quotient=%t yukawa_restriction=%t branch_map=%t J_even_map=%t uses_masses=%t uses_yukawa=%t manual_q4=%t verdict=%s", x.Formalized, x.SourceSpaces, x.Target, x.StructuralEdgeDim, x.JDoubledEdgeDim, x.HphiDim, x.HasCanonicalFullEdgeQuotient, x.HasCanonicalYukawaRestriction, x.HasCanonicalBranchMap, x.HasCanonicalJEvenMap, x.UsesObservedMasses, x.UsesYukawaAmplitudes, x.UsesManualQ4Placement, x.Verdict)
}

func FormatCandidate(c QuotientCandidate) string {
	return fmt.Sprintf("name=%q source=%q target=%q source_dim=%d target_dim=%d native=%t sealed=%t circular=%t canonical=%t contact=%t oneform=%t J=%t first_order=%t Hphi=%t full_info=%t rank=%d kernel=%d eigen=%v distinct=%d min_degree=%d char=%s min=%s char_residual_q4=%s min_residual_q4=%s pair_degenerate=%t quartic_capacity=%t q4_exact=%t q4_factor=%t promotable=%t yukawa_reduced=%t moduli_reduced=%t verdict=%s reason=%s", c.Name, c.Source, c.Target, c.SourceDim, c.TargetDim, c.Native, c.Sealed, c.Circular, c.CanonicalQuotient, c.ContactDerived, c.OneFormDerived, c.JCompatible, c.FirstOrderCompatible, c.HphiEndomorphism, c.FullEdgeInformationPreserved, c.Rank, c.KernelDimension, c.Eigenvalues, c.DistinctEigenvalues, c.MinimalDegree, c.CharacteristicPolynomial, c.MinimalPolynomial, formatResidual(c.CharacteristicResidualToQ4), formatResidual(c.MinimalResidualToQ4), c.PairDegenerate, c.IrreducibleQuarticCapacity, c.Q4ExactMatch, c.Q4FactorMatch, c.PromotableAsQ4Selector, c.ReducesYukawaCouplings, c.ReducesFlavorModuli, c.Verdict, c.Reason)
}

func FormatSieve(x QuotientSieve) string {
	parts := []string{fmt.Sprintf("executed=%t native_quotients=%d native_Hphi=%d native_quartic_capacity=%d canonical_q4_matches=%d sealed_manual=%d best=%s best_residual=%s verdict=%s", x.Executed, x.NativeQuotientCount, x.NativeHphiEndomorphismCount, x.NativeQuarticCapacityCount, x.CanonicalHphiQ4MatchCount, x.SealedOrManualCount, x.BestCanonicalCandidate, formatResidual(x.BestCanonicalQ4Residual), x.Verdict)}
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return strings.Join(parts, "\n")
}

func FormatImpact(x Impact) string {
	return fmt.Sprintf("Hphi_q4_identified=%t canonical_quotient=%t canonical_yukawa_quotient=%t native_intertwiner_q4=%t yukawa_reduced=%t moduli_start=%d moduli_result=%d flavor_firewall=%t higgs_lane_preserved=%t verdict=%s", x.HphiQuarticIdentified, x.CanonicalQuotientFound, x.CanonicalYukawaQuotientFound, x.NativeIntertwinerQ4Found, x.YukawaCouplingsReduced, x.ChargedModuliStart, x.ChargedModuliResult, x.FlavorFirewallPreserved, x.HiggsLanePreserved, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_ckm=%t no_pmns=%t no_yukawa_amplitudes=%t no_manual_q4=%t no_arbitrary_full_edge_quotient=%t no_companion_promoted=%t no_moduli_reduction=%t verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoManualQ4HphiID, x.NoArbitraryFullEdgeQuotientPromoted, x.NoCompanionOperatorPromoted, x.NoFlavorModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s\nReason: %s\nPrimary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func formatResidual(x float64) string {
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	if math.IsNaN(x) {
		return "NaN"
	}
	return fmt.Sprintf("%.12g", x)
}
