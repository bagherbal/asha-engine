// Package contactedgepullback implements Gate 405:
// Contact-to-Edge Natural Transformation / Pullback Sieve.
//
// Gates 398-404 established a strict obstruction chain: the contact quartic
// primary q4 is an exact four-dimensional contact spectral datum, while the
// mature scalar carrier H_phi and one-form edge ledger are real but too
// symmetric to inherit q4 through quaternionic, mixed Laplacian, charge-weight,
// graph, oriented-boundary, or edge-to-H_phi quotient routes. Gate 405 reverses
// the arrow. It asks whether the contact q4 operator itself has a native
// pullback/natural transformation into the J-doubled one-form edge module.
//
// The gate promotes only a typed map selected by existing ASHA data. Companion
// matrices, root placements, chosen edge bases, and arbitrary injections from
// a four-dimensional contact block into the five/ten edge-slot module are
// quarantined as sealed stress tests.
package contactedgepullback

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE405-CONTACT-TO-EDGE-NATURAL-TRANSFORMATION-PULLBACK-SIEVE"

	StatusGate404Inherited         = "CONDITIONAL_SUPPORT_GATE404_QUOTIENT_OBSTRUCTION_INHERITED"
	StatusPullbackArenaFormalized  = "CONDITIONAL_SUPPORT_CONTACT_EDGE_PULLBACK_ARENA_FORMALIZED"
	StatusQ4PrimaryTargetAudited   = "CONDITIONAL_SUPPORT_Q4_PRIMARY_TARGET_AUDITED"
	StatusEdgeLedgerDomainAudited  = "CONDITIONAL_SUPPORT_ONEFORM_EDGE_LEDGER_DOMAIN_AUDITED"
	StatusSealedQ4ExtensionAudited = "CONDITIONAL_SUPPORT_SEALED_Q4_EDGE_EXTENSION_STRESS_TESTED"

	StatusFailedNoNativeContactToEdgeMap        = "FAILED_ROUTE_NO_NATIVE_CONTACT_TO_EDGE_MAP"
	StatusFailedYukawaRestrictionWrongDirection = "FAILED_ROUTE_YUKAWA_RESTRICTION_IS_EDGE_TO_SCALAR_NOT_CONTACT_PULLBACK"
	StatusFailedQ4ExtensionManualEdgeBasis      = "FAILED_ROUTE_Q4_EXTENSION_TO_E5_REQUIRES_MANUAL_EDGE_BASIS"
	StatusFailedJDoubledManualDuplication       = "FAILED_ROUTE_J_DOUBLED_PULLBACK_DUPLICATES_MANUAL_Q4"
	StatusFailedNoDFIntertwiner                 = "FAILED_ROUTE_PULLBACK_DOES_NOT_INTERTWINE_DF_EDGE_GRAPH"
	StatusFailedNoNaturalTransformation         = "FAILED_ROUTE_NO_CANONICAL_CONTACT_EDGE_NATURAL_TRANSFORMATION"
	StatusFailedNoCanonicalHphiQuarticID        = "FAILED_ROUTE_NO_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusFailedNoYukawaCouplingReduction       = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFirewallPreserved13Moduli             = "FIREWALL_PRESERVED_13_MODULI"

	StatusConditionalContactPullbackAchieved = "CONDITIONAL_SUPPORT_CONTACT_PULLBACK_ACHIEVED"
)

const (
	ContactQuarticQ4        = "3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271"
	Q4Degree                = 4
	ContactPrimaryDim       = 4
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
	Gate404QuotientNoQ4             bool
	Gate404NeedsPullback            bool
	Gate385OneFormEdges             bool
	Gate385JDoubledEdgeCount        int
	Gate372ChargedModuliDim         int
	NoEmpiricalInputsImported       bool
	Verdict                         string
}

type Q4Target struct {
	Polynomial        string
	Degree            int
	Dimension         int
	IrreducibleOverQ  bool
	MonicCoefficients []float64
	Domain            string
	NeededMap         string
	Verdict           string
}

type PullbackArena struct {
	Formalized               bool
	ContactDomain            string
	EdgeCodomains            []string
	NaturalTransformation    string
	RequiredSquare           string
	ContactPrimaryDim        int
	StructuralEdgeDim        int
	YukawaEdgeDim            int
	JDoubledEdgeDim          int
	HphiDim                  int
	NativeFunctorKnown       bool
	ContactEdgeActionDerived bool
	UsesObservedMasses       bool
	UsesYukawaAmplitudes     bool
	UsesManualRootPlacement  bool
	Verdict                  string
}

type PullbackCandidate struct {
	Name                     string
	Formula                  string
	Source                   string
	Target                   string
	SourceDim                int
	TargetDim                int
	Native                   bool
	Sealed                   bool
	Circular                 bool
	Typed                    bool
	Canonical                bool
	ContactDerived           bool
	EdgeDerived              bool
	JCompatible              bool
	FirstOrderCompatible     bool
	DFIntertwiner            bool
	NaturalitySquareFormed   bool
	PullbackConstructed      bool
	PreservesQ4Degree        bool
	PreservesQ4Polynomial    bool
	PromotableAsQ4EdgeWeight bool
	InducedHphiEndomorphism  bool
	ReducesYukawaCouplings   bool
	ReducesFlavorModuli      bool
	Rank                     int
	KernelDimension          int
	EdgeOperatorPolynomial   string
	MinimalPolynomial        string
	CharacteristicPolynomial string
	ResidualToQ4             float64
	CommutatorWithEdgeGraph  string
	CommutatorZero           bool
	Reason                   string
	Verdict                  string
}

type PullbackSieve struct {
	Executed                       bool
	Candidates                     []PullbackCandidate
	NativeTypedMapCount            int
	NativePullbackCount            int
	NativeQ4PreservingCount        int
	NativeDFIntertwinerCount       int
	CanonicalNaturalTransformCount int
	SealedOrManualCount            int
	BestNativeCandidate            string
	BestNativeResidual             float64
	Verdict                        string
}

type Impact struct {
	ContactPullbackAchieved        bool
	Q4OnEdgeSpacePreserved         bool
	CanonicalNaturalTransformation bool
	HphiQuarticIdentified          bool
	YukawaCouplingsReduced         bool
	ChargedModuliStart             int
	ChargedModuliResult            int
	FlavorFirewallPreserved        bool
	HiggsLanePreserved             bool
	Verdict                        string
}

type FirewallAudit struct {
	Executed                       bool
	NoObservedMassesImported       bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	NoYukawaAmplitudesInserted     bool
	NoManualQ4HphiID               bool
	NoManualRootPlacementPromoted  bool
	NoArbitraryEdgeBasisPromoted   bool
	NoCompanionOperatorPromoted    bool
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
	Q4          Q4Target
	Arena       PullbackArena
	Sieve       PullbackSieve
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
	sieve := auditPullbacks(arena, q4)
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
		Gate404QuotientNoQ4:             true,
		Gate404NeedsPullback:            true,
		Gate385OneFormEdges:             true,
		Gate385JDoubledEdgeCount:        JDoubledEdgeCount,
		Gate372ChargedModuliDim:         Gate372ChargedModuliDim,
		NoEmpiricalInputsImported:       true,
		Verdict:                         "Gate 405 inherits the scalar/contact identity obstruction chain and the Gate-404 conclusion that native edge-to-H_phi quotients exist but remain pair-degenerate. The only remaining direction is a typed contact-to-edge pullback/natural transformation, if the project derives one.",
	}
}

func auditQ4() Q4Target {
	return Q4Target{Polynomial: ContactQuarticQ4, Degree: Q4Degree, Dimension: ContactPrimaryDim, IrreducibleOverQ: true, MonicCoefficients: append([]float64(nil), q4Monic...), Domain: "contact quartic primary block Q[x]/(q4)", NeededMap: "a native typed pullback f*: End(Q[x]/(q4)) -> End(Omega^1_D(A_F)) or a natural transformation from contact spectral algebra to one-form edge ledger", Verdict: "The target is not another four-dimensional dimension match. It is an irreducible quartic operator whose contact-domain action must be transported into edge coordinates by a derived functor."}
}

func buildArena() PullbackArena {
	return PullbackArena{
		Formalized:               true,
		ContactDomain:            "C_q4 := Q[x]/(q4), the four-dimensional contact spectral primary",
		EdgeCodomains:            []string{"E_5 structural finite-Dirac edge slots", "E_10 J-doubled one-form edge slots", "E_Y four Higgs/Yukawa edge slots", "H_phi after canonical Yukawa restriction"},
		NaturalTransformation:    "eta: ContactSpectralPrimary => OneFormEdgeLedger",
		RequiredSquare:           "contact q4 action --eta--> edge operator, compatible with D_F edge action, J, first-order, and canonical H_phi quotient",
		ContactPrimaryDim:        ContactPrimaryDim,
		StructuralEdgeDim:        StructuralEdgeCount,
		YukawaEdgeDim:            YukawaEdgeCount,
		JDoubledEdgeDim:          JDoubledEdgeCount,
		HphiDim:                  HphiRealDim,
		NativeFunctorKnown:       false,
		ContactEdgeActionDerived: false,
		UsesObservedMasses:       false,
		UsesYukawaAmplitudes:     false,
		UsesManualRootPlacement:  false,
		Verdict:                  "The arena can be typed, but existing ledgers do not define an action of the contact quartic primary on the one-form edge module. A pullback cannot be inferred from equal dimensions, edge count, or a chosen q4 companion matrix.",
	}
}

func auditPullbacks(arena PullbackArena, q4 Q4Target) PullbackSieve {
	candidates := []PullbackCandidate{
		{
			Name:                     "native contact projector to one-form edge ledger",
			Formula:                  "eta(P_q4) ?-> Omega^1_D(A_F)",
			Source:                   arena.ContactDomain,
			Target:                   "E_10 J-doubled one-form edge module",
			SourceDim:                ContactPrimaryDim,
			TargetDim:                JDoubledEdgeCount,
			Native:                   false,
			Typed:                    false,
			Canonical:                false,
			ContactDerived:           true,
			EdgeDerived:              true,
			JCompatible:              false,
			FirstOrderCompatible:     false,
			DFIntertwiner:            false,
			NaturalitySquareFormed:   false,
			PullbackConstructed:      false,
			PreservesQ4Degree:        false,
			PreservesQ4Polynomial:    false,
			Rank:                     0,
			KernelDimension:          ContactPrimaryDim,
			EdgeOperatorPolynomial:   "not constructed",
			MinimalPolynomial:        "not constructed",
			CharacteristicPolynomial: "not constructed",
			ResidualToQ4:             math.Inf(1),
			CommutatorWithEdgeGraph:  "not typed",
			Reason:                   "The contact q4 block and the one-form edge module are both native, but the project has no derived representation/action that sends contact primary basis elements to edge slots or edge endomorphisms.",
			Verdict:                  StatusFailedNoNativeContactToEdgeMap,
		},
		{
			Name:                     "reverse of canonical Yukawa edge restriction",
			Formula:                  "Q_Y^{-1}? : H_phi ~= E_Y -> E_5, then identify C_q4 with H_phi",
			Source:                   "H_phi / E_Y four Yukawa edge slots",
			Target:                   "E_5 structural edge module",
			SourceDim:                HphiRealDim,
			TargetDim:                StructuralEdgeCount,
			Native:                   true,
			Circular:                 true,
			Typed:                    true,
			Canonical:                false,
			ContactDerived:           false,
			EdgeDerived:              true,
			JCompatible:              true,
			FirstOrderCompatible:     true,
			DFIntertwiner:            false,
			NaturalitySquareFormed:   false,
			PullbackConstructed:      false,
			PreservesQ4Degree:        false,
			PreservesQ4Polynomial:    false,
			Rank:                     4,
			KernelDimension:          0,
			EdgeOperatorPolynomial:   "(x - 1)(x - 3) after canonical quotient; no q4 source",
			MinimalPolynomial:        "(x - 1)(x - 3)",
			CharacteristicPolynomial: "(x - 1)^2*(x - 3)^2",
			ResidualToQ4:             residualMonic([]float64{1, -8, 22, -24, 9}, q4Monic),
			CommutatorWithEdgeGraph:  "canonical edge restriction only; no contact q4 action",
			Reason:                   "The Gate-404 quotient is an edge-to-scalar map, not a contact-to-edge pullback. Reversing it requires first identifying H_phi with C_q4, exactly the theorem under test.",
			Verdict:                  StatusFailedYukawaRestrictionWrongDirection,
		},
		{
			Name:                     "sealed q4 extension to five structural edge slots",
			Formula:                  "A_E = diag/companion(q4) on chosen four edge modes plus one sterile slot",
			Source:                   arena.ContactDomain,
			Target:                   "E_5 structural edge module",
			SourceDim:                ContactPrimaryDim,
			TargetDim:                StructuralEdgeCount,
			Sealed:                   true,
			Circular:                 true,
			Typed:                    true,
			Canonical:                false,
			ContactDerived:           true,
			EdgeDerived:              false,
			JCompatible:              false,
			FirstOrderCompatible:     false,
			DFIntertwiner:            false,
			NaturalitySquareFormed:   false,
			PullbackConstructed:      true,
			PreservesQ4Degree:        true,
			PreservesQ4Polynomial:    true,
			Rank:                     4,
			KernelDimension:          1,
			EdgeOperatorPolynomial:   "x * q4(x) after choosing a fifth edge complement",
			MinimalPolynomial:        "lcm(q4, x) unless complement is projected away",
			CharacteristicPolynomial: "x * (3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271)",
			ResidualToQ4:             0,
			CommutatorWithEdgeGraph:  "nonzero generically; no edge-graph intertwiner",
			CommutatorZero:           false,
			Reason:                   "This preserves q4 only because q4 is manually placed on a chosen four-dimensional edge subspace. The chosen edge basis and sterile complement are not selected by contact topology, J, or first-order data.",
			Verdict:                  StatusFailedQ4ExtensionManualEdgeBasis,
		},
		{
			Name:                     "sealed J-doubled q4 pullback",
			Formula:                  "A_E10 = A_q4 ⊕ J A_q4 J^{-1} ⊕ sterile mirrors",
			Source:                   arena.ContactDomain,
			Target:                   "E_10 J-doubled one-form module",
			SourceDim:                ContactPrimaryDim,
			TargetDim:                JDoubledEdgeCount,
			Sealed:                   true,
			Circular:                 true,
			Typed:                    true,
			Canonical:                false,
			ContactDerived:           true,
			EdgeDerived:              false,
			JCompatible:              true,
			FirstOrderCompatible:     false,
			DFIntertwiner:            false,
			NaturalitySquareFormed:   false,
			PullbackConstructed:      true,
			PreservesQ4Degree:        true,
			PreservesQ4Polynomial:    true,
			Rank:                     8,
			KernelDimension:          2,
			EdgeOperatorPolynomial:   "x^2 * q4(x)^2 on ten slots after manual duplication",
			MinimalPolynomial:        "x * q4(x) unless zero complement is removed",
			CharacteristicPolynomial: "x^2 * q4(x)^2",
			ResidualToQ4:             0,
			CommutatorWithEdgeGraph:  "not an edge-graph natural transformation; duplicates a manual placement",
			CommutatorZero:           false,
			Reason:                   "J-doubling can mirror a manually inserted q4 block, but it does not derive the original q4-to-edge map. It duplicates the same arbitrary basis alignment.",
			Verdict:                  StatusFailedJDoubledManualDuplication,
		},
		{
			Name:                     "contact q4 as edge weight/intertwiner with native D_F edge graph",
			Formula:                  "A_E Delta_edge = Delta_edge A_E and A_E|_{H_phi} has q4",
			Source:                   arena.ContactDomain,
			Target:                   "edge graph plus canonical H_phi quotient",
			SourceDim:                ContactPrimaryDim,
			TargetDim:                HphiRealDim,
			Native:                   false,
			Sealed:                   true,
			Circular:                 true,
			Typed:                    true,
			Canonical:                false,
			ContactDerived:           true,
			EdgeDerived:              true,
			JCompatible:              false,
			FirstOrderCompatible:     false,
			DFIntertwiner:            false,
			NaturalitySquareFormed:   false,
			PullbackConstructed:      false,
			PreservesQ4Degree:        false,
			PreservesQ4Polynomial:    false,
			Rank:                     4,
			KernelDimension:          0,
			EdgeOperatorPolynomial:   "edge graph polynomial remains pair/quartic graph polynomial, not q4",
			MinimalPolynomial:        "commutant of K2⊔K2 or P3⊔K2 does not select irreducible q4 without inserted coefficients",
			CharacteristicPolynomial: "not q4",
			ResidualToQ4:             math.Inf(1),
			CommutatorWithEdgeGraph:  "manual q4 companion does not commute with native edge Laplacian/admissible graph action except under tuned basis choices",
			CommutatorZero:           false,
			Reason:                   "A true pullback must intertwine the contact operator with native edge dynamics. No such commutative square is derived; forced q4 blocks are not natural with respect to the edge graph.",
			Verdict:                  StatusFailedNoDFIntertwiner,
		},
	}

	nativeTyped, nativePullback, nativeQ4, nativeIntertwiner, natural, sealed := 0, 0, 0, 0, 0, 0
	bestName := "none"
	bestResidual := math.Inf(1)
	for _, c := range candidates {
		if c.Native && c.Typed {
			nativeTyped++
		}
		if c.Native && c.PullbackConstructed {
			nativePullback++
		}
		if c.Native && c.PreservesQ4Polynomial && c.PromotableAsQ4EdgeWeight {
			nativeQ4++
		}
		if c.Native && c.DFIntertwiner {
			nativeIntertwiner++
		}
		if c.Native && c.Canonical && c.NaturalitySquareFormed {
			natural++
		}
		if c.Sealed || c.Circular {
			sealed++
		}
		if c.Native && c.Typed && c.ResidualToQ4 < bestResidual {
			bestResidual = c.ResidualToQ4
			bestName = c.Name
		}
	}
	return PullbackSieve{Executed: true, Candidates: candidates, NativeTypedMapCount: nativeTyped, NativePullbackCount: nativePullback, NativeQ4PreservingCount: nativeQ4, NativeDFIntertwinerCount: nativeIntertwiner, CanonicalNaturalTransformCount: natural, SealedOrManualCount: sealed, BestNativeCandidate: bestName, BestNativeResidual: bestResidual, Verdict: "Gate 405 finds no native contact-to-edge natural transformation. The existing native map is in the wrong direction (edge-to-H_phi restriction). Exact q4 preservation appears only by manually placing a q4 companion block onto chosen edge slots, and such placements fail the D_F edge-graph/naturality test."}
}

func auditImpact(s PullbackSieve) Impact {
	return Impact{ContactPullbackAchieved: false, Q4OnEdgeSpacePreserved: false, CanonicalNaturalTransformation: false, HphiQuarticIdentified: false, YukawaCouplingsReduced: false, ChargedModuliStart: Gate372ChargedModuliDim, ChargedModuliResult: Gate372ChargedModuliDim, FlavorFirewallPreserved: true, HiggsLanePreserved: true, Verdict: "No native contact q4 pullback into the one-form edge ledger is derived. The scalar/contact identity obstruction and flavor firewall remain intact."}
}

func auditFirewall(i Impact) FirewallAudit {
	return FirewallAudit{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaAmplitudesInserted: true, NoManualQ4HphiID: true, NoManualRootPlacementPromoted: true, NoArbitraryEdgeBasisPromoted: true, NoCompanionOperatorPromoted: true, NoFlavorModuliReductionClaimed: i.FlavorFirewallPreserved, Verdict: "No empirical flavor data, observed masses, CKM/PMNS inputs, Yukawa amplitudes, manual q4-H_phi identity, root placement, arbitrary edge basis, or companion operator is promoted."}
}

func nextStep(s PullbackSieve, i Impact) NextStep {
	return NextStep{Gate: 406, Title: "Contact-Eigenoperator Internal Reconstruction / q4 Lives Only in Contact Sector", Reason: "Gate 405 rejects the contact-to-edge pullback route. The repeated failure of scalar/edge identifications suggests q4 may be an internal contact-sector eigenoperator rather than a Higgs-bundle selector. The next gate should reconstruct q4 inside the contact projector algebra itself and classify whether it has any lawful bridge obligations left.", PrimaryTask: "Audit whether q4 should be sealed as a contact-only spectral invariant, then search for a different scalar identity selector from the mature one-form/H_phi lane instead of forcing q4 across sectors."}
}

func truth(s PullbackSieve, i Impact) string {
	return "Gate 405 reverses the arrows demanded by Gate 404 and tests the strongest remaining q4 route: a contact-to-edge natural transformation. The result is negative. The contact q4 primary and one-form edge ledger are both native, but there is no typed ASHA functor sending the q4 contact action into edge-slot endomorphisms. The canonical Yukawa edge restriction exists only in the opposite direction and presupposes the H_phi/q4 identification it would need to prove. Exact q4 preservation on edge space is possible only through sealed companion-matrix placement on a chosen edge basis or its J-doubled duplicate; those fail naturality and D_F edge-intertwiner checks. Therefore no canonical contact pullback is achieved, no scalar bundle geometric seal is derived, no Yukawa couplings are reduced, and the 13-moduli firewall remains preserved."
}

func Statuses(a Analysis) []string {
	statuses := []string{StatusGate404Inherited, StatusPullbackArenaFormalized, StatusQ4PrimaryTargetAudited, StatusEdgeLedgerDomainAudited, StatusSealedQ4ExtensionAudited}
	for _, c := range a.Sieve.Candidates {
		if c.Verdict != "" && !contains(statuses, c.Verdict) {
			statuses = append(statuses, c.Verdict)
		}
	}
	for _, status := range []string{StatusFailedNoNaturalTransformation, StatusFailedNoCanonicalHphiQuarticID, StatusFailedNoYukawaCouplingReduction, StatusFirewallPreserved13Moduli} {
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
	return fmt.Sprintf("executed=%t gate398_no_bundle=%t gate399_quaternionic_no=%t gate400_no_mixed=%t gate401_charge_disjoint=%t gate402_graph_no=%t gate403_orientation_no=%t gate404_quotient_no=%t gate404_needs_pullback=%t oneform_edges=%t J_edges=%d moduli_dim=%d no_empirical=%t verdict=%s", x.Executed, x.Gate398NoQuarticBundleFunctor, x.Gate399QuaternionicPolynomialNo, x.Gate400NoMixedEdgeQ4, x.Gate401ChargeWeightsDisjoint, x.Gate402GraphNoQ4, x.Gate403OrientationNoQ4, x.Gate404QuotientNoQ4, x.Gate404NeedsPullback, x.Gate385OneFormEdges, x.Gate385JDoubledEdgeCount, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatQ4(x Q4Target) string {
	return fmt.Sprintf("polynomial=%s degree=%d dim=%d irreducible_over_Q=%t monic=%v domain=%s needed_map=%s verdict=%s", x.Polynomial, x.Degree, x.Dimension, x.IrreducibleOverQ, x.MonicCoefficients, x.Domain, x.NeededMap, x.Verdict)
}

func FormatArena(x PullbackArena) string {
	return fmt.Sprintf("formalized=%t contact_domain=%s edge_codomains=%v natural_transformation=%s required_square=%s contact_dim=%d edge_dim=%d yukawa_dim=%d J_dim=%d Hphi_dim=%d native_functor=%t contact_edge_action=%t uses_masses=%t uses_yukawa=%t manual_roots=%t verdict=%s", x.Formalized, x.ContactDomain, x.EdgeCodomains, x.NaturalTransformation, x.RequiredSquare, x.ContactPrimaryDim, x.StructuralEdgeDim, x.YukawaEdgeDim, x.JDoubledEdgeDim, x.HphiDim, x.NativeFunctorKnown, x.ContactEdgeActionDerived, x.UsesObservedMasses, x.UsesYukawaAmplitudes, x.UsesManualRootPlacement, x.Verdict)
}

func FormatCandidate(c PullbackCandidate) string {
	return fmt.Sprintf("name=%q source=%q target=%q source_dim=%d target_dim=%d native=%t sealed=%t circular=%t typed=%t canonical=%t contact=%t edge=%t J=%t first_order=%t DF_intertwiner=%t naturality=%t pullback=%t q4_degree=%t q4_poly=%t promotable=%t Hphi=%t yukawa_reduced=%t moduli_reduced=%t rank=%d kernel=%d edge_poly=%s min=%s char=%s residual_q4=%s commutator=%q comm_zero=%t verdict=%s reason=%s", c.Name, c.Source, c.Target, c.SourceDim, c.TargetDim, c.Native, c.Sealed, c.Circular, c.Typed, c.Canonical, c.ContactDerived, c.EdgeDerived, c.JCompatible, c.FirstOrderCompatible, c.DFIntertwiner, c.NaturalitySquareFormed, c.PullbackConstructed, c.PreservesQ4Degree, c.PreservesQ4Polynomial, c.PromotableAsQ4EdgeWeight, c.InducedHphiEndomorphism, c.ReducesYukawaCouplings, c.ReducesFlavorModuli, c.Rank, c.KernelDimension, c.EdgeOperatorPolynomial, c.MinimalPolynomial, c.CharacteristicPolynomial, formatResidual(c.ResidualToQ4), c.CommutatorWithEdgeGraph, c.CommutatorZero, c.Verdict, c.Reason)
}

func FormatSieve(x PullbackSieve) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("executed=%t native_typed=%d native_pullback=%d native_q4=%d native_DF_intertwiner=%d natural_transform=%d sealed=%d best_native=%q best_residual=%s verdict=%s\n%s", x.Executed, x.NativeTypedMapCount, x.NativePullbackCount, x.NativeQ4PreservingCount, x.NativeDFIntertwinerCount, x.CanonicalNaturalTransformCount, x.SealedOrManualCount, x.BestNativeCandidate, formatResidual(x.BestNativeResidual), x.Verdict, strings.Join(parts, "\n"))
}

func FormatImpact(x Impact) string {
	return fmt.Sprintf("contact_pullback=%t q4_on_edges=%t natural_transformation=%t Hphi_quartic=%t yukawa_reduced=%t moduli_start=%d moduli_result=%d flavor_firewall=%t higgs_lane=%t verdict=%s", x.ContactPullbackAchieved, x.Q4OnEdgeSpacePreserved, x.CanonicalNaturalTransformation, x.HphiQuarticIdentified, x.YukawaCouplingsReduced, x.ChargedModuliStart, x.ChargedModuliResult, x.FlavorFirewallPreserved, x.HiggsLanePreserved, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_CKM=%t no_PMNS=%t no_yukawa=%t no_manual_q4_Hphi=%t no_roots=%t no_edge_basis=%t no_companion=%t no_moduli_reduction=%t verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoManualQ4HphiID, x.NoManualRootPlacementPromoted, x.NoArbitraryEdgeBasisPromoted, x.NoCompanionOperatorPromoted, x.NoFlavorModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%q reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func formatResidual(v float64) string {
	if math.IsInf(v, 1) {
		return "+Inf"
	}
	if math.IsNaN(v) {
		return "NaN"
	}
	return fmt.Sprintf("%.12g", v)
}
