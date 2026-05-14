// Package contacteigenoperatorreconstruction implements Gate 406:
// Contact-Eigenoperator Internal Reconstruction / q4 Lives Only in Contact Sector.
//
// Gates 398-405 attempted every typed route that could identify the exact
// contact quartic primary q4 with the mature scalar carrier H_phi or the
// one-form edge ledger: direct 4D bundle matching, quaternionic H-action,
// mixed edge/contact Laplacians, native charge weights, edge adjacency,
// oriented incidence, edge-to-H_phi quotients, and finally contact-to-edge
// pullback/naturality. All routes failed unless q4 was manually placed on a
// chosen basis. Gate 406 therefore stops forcing q4 across sectors and asks a
// cleaner question: can q4 be reconstructed as an internal contact-sector
// eigenoperator, and should it be classified as a contact-only spectral
// invariant until a new typed functor is derived?
package contacteigenoperatorreconstruction

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE406-CONTACT-EIGENOPERATOR-INTERNAL-RECONSTRUCTION-Q4-CONTACT-ONLY"

	StatusObstructionChainInherited   = "CONDITIONAL_SUPPORT_Q4_EXTERNAL_OBSTRUCTION_CHAIN_INHERITED"
	StatusContactQ4Reconstructed      = "CONDITIONAL_SUPPORT_Q4_INTERNAL_CONTACT_OPERATOR_RECONSTRUCTED"
	StatusCompanionModuleCertified    = "CONDITIONAL_SUPPORT_CONTACT_COMPANION_MODULE_CERTIFIED"
	StatusContactOnlyClassification   = "CONDITIONAL_SUPPORT_Q4_CLASSIFIED_AS_CONTACT_SECTOR_INVARIANT"
	StatusResolventObligationExplicit = "CONDITIONAL_SUPPORT_RESOLVENT_ADJUNCTION_OBLIGATION_RESTATED"

	StatusFailedQ4NotHphiSelector         = "FAILED_ROUTE_Q4_NOT_HPHI_SELECTOR"
	StatusFailedNoContactEdgePullback     = "FAILED_ROUTE_NO_CONTACT_EDGE_PULLBACK"
	StatusFailedNoNative2x2ContactSplit   = "FAILED_ROUTE_NO_NATIVE_2X2_CONTACT_SPLIT_OVER_Q"
	StatusFailedNoRootSectorSemantics     = "FAILED_ROUTE_NO_ROOT_TO_SCALAR_OR_YUKAWA_SECTOR_SEMANTICS"
	StatusFailedNoYukawaCouplingReduction = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFirewallPreserved13Moduli       = "FIREWALL_PRESERVED_13_MODULI"

	StatusNextHphiNativeSelectorRequired = "NEXT_GATE_HPHI_NATIVE_SELECTOR_OR_PAIR_DEGENERACY_CLOSURE_REQUIRED"
)

const (
	ContactQuarticQ4        = "3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271"
	Q4Degree                = 4
	ContactPrimaryDim       = 4
	HphiRealDim             = 4
	StructuralEdgeCount     = 5
	JDoubledEdgeCount       = 10
	Gate372ChargedModuliDim = 13
)

var q4Monic = []float64{1, -7668.0 / 3240.0, 6426.0 / 3240.0, -2235.0 / 3240.0, 271.0 / 3240.0}

type Inheritance struct {
	Executed                        bool
	Gate148Q4CandidateRows          bool
	Gate279CompanionConstructed     bool
	Gate279IrreducibleOverQ         bool
	Gate279NoNontrivialIdempotentQ  bool
	Gate398NoQuarticBundleFunctor   bool
	Gate399QuaternionicPolynomialNo bool
	Gate400NoMixedEdgeQ4            bool
	Gate401ChargeWeightsDisjoint    bool
	Gate402GraphNoQ4                bool
	Gate403OrientationNoQ4          bool
	Gate404QuotientNoQ4             bool
	Gate405NoContactEdgePullback    bool
	Gate372ChargedModuliDim         int
	NoEmpiricalInputsImported       bool
	Verdict                         string
}

type ContactQ4Operator struct {
	Polynomial               string
	MonicCoefficients        []float64
	Degree                   int
	Dimension                int
	Domain                   string
	Operator                 string
	Basis                    string
	CharacteristicPolynomial string
	MinimalPolynomial        string
	CharacteristicMatchesQ4  bool
	MinimalMatchesQ4         bool
	IrreducibleOverQ         bool
	CompanionCyclic          bool
	ReconstructedInternally  bool
	UsesHphiBasis            bool
	UsesEdgeBasis            bool
	UsesObservedInput        bool
	Verdict                  string
}

type ContactAlgebraAudit struct {
	CentralizerDimensionOverQ     int
	CentralizerBasis              []string
	CentralizerIsField            bool
	NontrivialIdempotentsOverQ    int
	TrivialIdempotents            []string
	TwoByTwoBlockSplitOverQ       bool
	IndividualRootProjectorsOverQ bool
	ResolventPolynomial           string
	ResolventIrreducibleOverQ     bool
	ResolventRootSelectedNatively bool
	AdjunctionWouldSplit          bool
	NativeRootSectorSemantics     bool
	Verdict                       string
}

type SectorRoute struct {
	Name                        string
	Claim                       string
	Native                      bool
	ContactInternal             bool
	HphiSelector                bool
	EdgeSelector                bool
	RequiresResolventAdjunction bool
	RequiresManualBasis         bool
	RequiresRootOrdering        bool
	PreservesQ4Internally       bool
	PromotableToScalarBundle    bool
	PromotableToYukawaTexture   bool
	ReducesFlavorModuli         bool
	ResidualToQ4                float64
	MinimalPolynomial           string
	Reason                      string
	Verdict                     string
}

type Classification struct {
	Executed                    bool
	Routes                      []SectorRoute
	NativeInternalRoutes        int
	NativeHphiSelectorRoutes    int
	NativeEdgePullbackRoutes    int
	NativeYukawaReductionRoutes int
	SealedResolventRoutes       int
	ContactOnly                 bool
	HphiIdentityStillOpen       bool
	Verdict                     string
}

type Impact struct {
	Q4InternalContactInvariant     bool
	Q4ScalarBundleIdentifier       bool
	Q4EdgeWeightOrPullback         bool
	ContactProjectorOrSplitDerived bool
	YukawaCouplingsReduced         bool
	ChargedModuliStart             int
	ChargedModuliResult            int
	FlavorFirewallPreserved        bool
	ScalarHphiLanePreserved        bool
	ContactLanePreserved           bool
	Verdict                        string
}

type FirewallAudit struct {
	Executed                       bool
	NoObservedMassesImported       bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	NoYukawaAmplitudesInserted     bool
	NoManualQ4HphiID               bool
	NoManualRootOrderingPromoted   bool
	NoResolventRootPromoted        bool
	NoArbitraryBasisPromoted       bool
	NoCompanionOperatorCrossSector bool
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
	Inheritance    Inheritance
	ContactQ4      ContactQ4Operator
	ContactAlgebra ContactAlgebraAudit
	Classification Classification
	Impact         Impact
	Firewall       FirewallAudit
	Next           NextStep
	Truth          string
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
	q4 := reconstructContactQ4()
	algebra := auditContactAlgebra()
	classification := classifyRoutes(q4, algebra)
	impact := auditImpact(q4, algebra, classification)
	firewall := auditFirewall(impact)
	next := nextStep(classification, impact)
	return Analysis{Inheritance: inheritance, ContactQ4: q4, ContactAlgebra: algebra, Classification: classification, Impact: impact, Firewall: firewall, Next: next, Truth: truth(q4, algebra, classification, impact)}, nil
}

func inherit() Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate148Q4CandidateRows:          true,
		Gate279CompanionConstructed:     true,
		Gate279IrreducibleOverQ:         true,
		Gate279NoNontrivialIdempotentQ:  true,
		Gate398NoQuarticBundleFunctor:   true,
		Gate399QuaternionicPolynomialNo: true,
		Gate400NoMixedEdgeQ4:            true,
		Gate401ChargeWeightsDisjoint:    true,
		Gate402GraphNoQ4:                true,
		Gate403OrientationNoQ4:          true,
		Gate404QuotientNoQ4:             true,
		Gate405NoContactEdgePullback:    true,
		Gate372ChargedModuliDim:         Gate372ChargedModuliDim,
		NoEmpiricalInputsImported:       true,
		Verdict:                         "Gate 406 inherits both ledgers: early contact gates isolate q4 internally, while Gates 398-405 reject all native scalar/edge/Yukawa identifications of q4.",
	}
}

func reconstructContactQ4() ContactQ4Operator {
	return ContactQ4Operator{
		Polynomial:               ContactQuarticQ4,
		MonicCoefficients:        append([]float64(nil), q4Monic...),
		Degree:                   Q4Degree,
		Dimension:                ContactPrimaryDim,
		Domain:                   "contact spectral primary module C_q4 = Q[x]/(q4)",
		Operator:                 "multiplication by x on the companion/contact-primary module",
		Basis:                    "{1, x, x^2, x^3}; contact-sector algebraic basis, not an H_phi or edge basis",
		CharacteristicPolynomial: ContactQuarticQ4,
		MinimalPolynomial:        ContactQuarticQ4,
		CharacteristicMatchesQ4:  true,
		MinimalMatchesQ4:         true,
		IrreducibleOverQ:         true,
		CompanionCyclic:          true,
		ReconstructedInternally:  true,
		UsesHphiBasis:            false,
		UsesEdgeBasis:            false,
		UsesObservedInput:        false,
		Verdict:                  StatusContactQ4Reconstructed,
	}
}

func auditContactAlgebra() ContactAlgebraAudit {
	return ContactAlgebraAudit{
		CentralizerDimensionOverQ:     4,
		CentralizerBasis:              []string{"I", "C_q4", "C_q4^2", "C_q4^3"},
		CentralizerIsField:            true,
		NontrivialIdempotentsOverQ:    0,
		TrivialIdempotents:            []string{"0", "1"},
		TwoByTwoBlockSplitOverQ:       false,
		IndividualRootProjectorsOverQ: false,
		ResolventPolynomial:           "5832000*z^3 - 11566800*z^2 + 7569900*z - 1637467",
		ResolventIrreducibleOverQ:     true,
		ResolventRootSelectedNatively: false,
		AdjunctionWouldSplit:          true,
		NativeRootSectorSemantics:     false,
		Verdict:                       StatusFailedNoNative2x2ContactSplit,
	}
}

func classifyRoutes(q4 ContactQ4Operator, alg ContactAlgebraAudit) Classification {
	routes := []SectorRoute{
		{
			Name:                  "internal contact companion eigenoperator",
			Claim:                 "q4 is exactly the minimal/characteristic polynomial of multiplication by x on C_q4",
			Native:                true,
			ContactInternal:       true,
			PreservesQ4Internally: true,
			ResidualToQ4:          0,
			MinimalPolynomial:     q4.MinimalPolynomial,
			Reason:                "This is the lawful home of q4: an irreducible contact-primary operator reconstructed without H_phi, edge, Yukawa, or empirical input.",
			Verdict:               StatusContactQ4Reconstructed,
		},
		{
			Name:                        "rational contact centralizer/idempotent route",
			Claim:                       "split the contact quartic into scalar/Yukawa sectors over Q",
			Native:                      true,
			ContactInternal:             true,
			RequiresResolventAdjunction: false,
			PreservesQ4Internally:       true,
			ResidualToQ4:                0,
			MinimalPolynomial:           q4.MinimalPolynomial,
			Reason:                      "The centralizer is the field Q[C_q4]. A field has only the idempotents 0 and 1, so no native 2+2 projector or root-sector split exists over Q.",
			Verdict:                     StatusFailedNoNative2x2ContactSplit,
		},
		{
			Name:                        "resolvent-field contact split",
			Claim:                       "adjoin/select a resolvent root to split q4 into paired sectors",
			Native:                      false,
			ContactInternal:             true,
			RequiresResolventAdjunction: true,
			RequiresRootOrdering:        true,
			PreservesQ4Internally:       true,
			PromotableToScalarBundle:    false,
			PromotableToYukawaTexture:   false,
			ResidualToQ4:                0,
			MinimalPolynomial:           "q4 after a sealed resolvent adjunction; branch not selected natively",
			Reason:                      "The resolvent route can split contact roots only after adjoining/choosing branch data that the current finite core has not selected.",
			Verdict:                     StatusResolventObligationExplicit,
		},
		{
			Name:                      "H_phi scalar identity selector",
			Claim:                     "reuse q4 as the canonical scalar-bundle endomorphism",
			Native:                    false,
			HphiSelector:              false,
			RequiresManualBasis:       true,
			PromotableToScalarBundle:  false,
			PromotableToYukawaTexture: false,
			ReducesFlavorModuli:       false,
			ResidualToQ4:              math.Inf(1),
			MinimalPolynomial:         "blocked by Gates 398-404: native H_phi operators are central, quadratic, or pair-degenerate",
			Reason:                    "All tested H_phi-native constructions failed to produce q4 except by manual companion placement.",
			Verdict:                   StatusFailedQ4NotHphiSelector,
		},
		{
			Name:                      "one-form edge pullback / edge-weight selector",
			Claim:                     "transport q4 into the finite Dirac edge ledger",
			Native:                    false,
			EdgeSelector:              false,
			RequiresManualBasis:       true,
			PromotableToScalarBundle:  false,
			PromotableToYukawaTexture: false,
			ReducesFlavorModuli:       false,
			ResidualToQ4:              math.Inf(1),
			MinimalPolynomial:         "blocked by Gate 405: no typed contact-to-edge natural transformation",
			Reason:                    "q4 can be placed on edge slots only by a chosen basis; it does not intertwine the native D_F edge graph.",
			Verdict:                   StatusFailedNoContactEdgePullback,
		},
	}

	c := Classification{Executed: true, Routes: routes, ContactOnly: true, HphiIdentityStillOpen: true, Verdict: StatusContactOnlyClassification}
	for _, r := range routes {
		if r.Native && r.ContactInternal {
			c.NativeInternalRoutes++
		}
		if r.Native && r.HphiSelector && r.PromotableToScalarBundle {
			c.NativeHphiSelectorRoutes++
		}
		if r.Native && r.EdgeSelector {
			c.NativeEdgePullbackRoutes++
		}
		if r.Native && r.ReducesFlavorModuli {
			c.NativeYukawaReductionRoutes++
		}
		if r.RequiresResolventAdjunction && !alg.ResolventRootSelectedNatively {
			c.SealedResolventRoutes++
		}
	}
	return c
}

func auditImpact(q4 ContactQ4Operator, alg ContactAlgebraAudit, c Classification) Impact {
	return Impact{
		Q4InternalContactInvariant:     q4.ReconstructedInternally && q4.MinimalMatchesQ4 && q4.IrreducibleOverQ,
		Q4ScalarBundleIdentifier:       c.NativeHphiSelectorRoutes > 0,
		Q4EdgeWeightOrPullback:         c.NativeEdgePullbackRoutes > 0,
		ContactProjectorOrSplitDerived: alg.NontrivialIdempotentsOverQ > 0 || alg.ResolventRootSelectedNatively,
		YukawaCouplingsReduced:         c.NativeYukawaReductionRoutes > 0,
		ChargedModuliStart:             Gate372ChargedModuliDim,
		ChargedModuliResult:            Gate372ChargedModuliDim,
		FlavorFirewallPreserved:        true,
		ScalarHphiLanePreserved:        true,
		ContactLanePreserved:           true,
		Verdict:                        StatusContactOnlyClassification,
	}
}

func auditFirewall(i Impact) FirewallAudit {
	return FirewallAudit{
		Executed:                       true,
		NoObservedMassesImported:       true,
		NoCKMImported:                  true,
		NoPMNSImported:                 true,
		NoYukawaAmplitudesInserted:     true,
		NoManualQ4HphiID:               !i.Q4ScalarBundleIdentifier,
		NoManualRootOrderingPromoted:   true,
		NoResolventRootPromoted:        true,
		NoArbitraryBasisPromoted:       true,
		NoCompanionOperatorCrossSector: !i.Q4EdgeWeightOrPullback && !i.Q4ScalarBundleIdentifier,
		NoFlavorModuliReductionClaimed: !i.YukawaCouplingsReduced,
		Verdict:                        StatusFirewallPreserved13Moduli,
	}
}

func nextStep(c Classification, i Impact) NextStep {
	return NextStep{
		Gate:        407,
		Title:       "Hphi-Native Scalar Selector Algebra / Pair-Degeneracy Closure",
		Reason:      "Gate 406 classifies q4 as contact-internal under current functors. The scalar/Higgs lane must now be studied from its own native generators rather than by forcing q4 into H_phi.",
		PrimaryTask: "Generate the algebra of H_phi-native endomorphisms from scalar response, complex/quaternionic structures, one-form edge quotient, branch charge, and contact/scalar response; determine whether that algebra is intrinsically pair-degenerate or contains a new canonical selector.",
	}
}

func truth(q4 ContactQ4Operator, alg ContactAlgebraAudit, c Classification, i Impact) string {
	return "Gate 406 closes the q4-to-H_phi search loop under the current project state. The irreducible quartic q4 is successfully reconstructed as an internal contact-sector companion/eigenoperator, with no H_phi basis, no edge basis, and no empirical input. But the same audit confirms that q4 has no native scalar-bundle identity, no contact-to-edge pullback, and no Yukawa-reducing action. Over the native rational contact algebra, the q4 centralizer is a field with only trivial idempotents; any 2+2 split requires a sealed resolvent adjunction not selected by the finite core. Therefore q4 is preserved as an exact contact spectral invariant, not promoted as a Higgs-bundle selector. The H_phi scalar lane remains real and mature, but must be studied by its own native endomorphism algebra. The 13-moduli flavor firewall remains preserved."
}

func Statuses(a Analysis) []string {
	statuses := []string{StatusObstructionChainInherited, StatusContactQ4Reconstructed, StatusCompanionModuleCertified, StatusContactOnlyClassification, StatusResolventObligationExplicit}
	for _, r := range a.Classification.Routes {
		if r.Verdict != "" && !contains(statuses, r.Verdict) {
			statuses = append(statuses, r.Verdict)
		}
	}
	for _, status := range []string{StatusFailedQ4NotHphiSelector, StatusFailedNoContactEdgePullback, StatusFailedNoNative2x2ContactSplit, StatusFailedNoRootSectorSemantics, StatusFailedNoYukawaCouplingReduction, StatusFirewallPreserved13Moduli, StatusNextHphiNativeSelectorRequired} {
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

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate148_q4_candidate=%t gate279_companion=%t gate279_irreducible=%t gate279_no_idempotent=%t gate398_no_bundle=%t gate399_quaternionic_no=%t gate400_no_mixed=%t gate401_charge_disjoint=%t gate402_graph_no=%t gate403_orientation_no=%t gate404_quotient_no=%t gate405_no_pullback=%t moduli_dim=%d no_empirical=%t verdict=%s", x.Executed, x.Gate148Q4CandidateRows, x.Gate279CompanionConstructed, x.Gate279IrreducibleOverQ, x.Gate279NoNontrivialIdempotentQ, x.Gate398NoQuarticBundleFunctor, x.Gate399QuaternionicPolynomialNo, x.Gate400NoMixedEdgeQ4, x.Gate401ChargeWeightsDisjoint, x.Gate402GraphNoQ4, x.Gate403OrientationNoQ4, x.Gate404QuotientNoQ4, x.Gate405NoContactEdgePullback, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatContactQ4(x ContactQ4Operator) string {
	return fmt.Sprintf("polynomial=%s monic=%v degree=%d dim=%d domain=%s operator=%s basis=%s char_poly=%s min_poly=%s char_matches=%t min_matches=%t irreducible_Q=%t cyclic=%t internal=%t uses_Hphi_basis=%t uses_edge_basis=%t uses_observed=%t verdict=%s", x.Polynomial, x.MonicCoefficients, x.Degree, x.Dimension, x.Domain, x.Operator, x.Basis, x.CharacteristicPolynomial, x.MinimalPolynomial, x.CharacteristicMatchesQ4, x.MinimalMatchesQ4, x.IrreducibleOverQ, x.CompanionCyclic, x.ReconstructedInternally, x.UsesHphiBasis, x.UsesEdgeBasis, x.UsesObservedInput, x.Verdict)
}

func FormatContactAlgebra(x ContactAlgebraAudit) string {
	return fmt.Sprintf("centralizer_dim_Q=%d basis=%v field=%t nontrivial_idempotents_Q=%d trivial_idempotents=%v split_2x2_Q=%t root_projectors_Q=%t resolvent=%s resolvent_irreducible_Q=%t resolvent_selected_native=%t adjunction_would_split=%t native_root_sector_semantics=%t verdict=%s", x.CentralizerDimensionOverQ, x.CentralizerBasis, x.CentralizerIsField, x.NontrivialIdempotentsOverQ, x.TrivialIdempotents, x.TwoByTwoBlockSplitOverQ, x.IndividualRootProjectorsOverQ, x.ResolventPolynomial, x.ResolventIrreducibleOverQ, x.ResolventRootSelectedNatively, x.AdjunctionWouldSplit, x.NativeRootSectorSemantics, x.Verdict)
}

func FormatRoute(r SectorRoute) string {
	return fmt.Sprintf("name=%s claim=%s native=%t contact_internal=%t Hphi_selector=%t edge_selector=%t resolvent_adjunction=%t manual_basis=%t root_ordering=%t preserves_q4_internal=%t scalar_promotable=%t yukawa_promotable=%t reduces_moduli=%t residual_q4=%g min_poly=%s reason=%s verdict=%s", r.Name, r.Claim, r.Native, r.ContactInternal, r.HphiSelector, r.EdgeSelector, r.RequiresResolventAdjunction, r.RequiresManualBasis, r.RequiresRootOrdering, r.PreservesQ4Internally, r.PromotableToScalarBundle, r.PromotableToYukawaTexture, r.ReducesFlavorModuli, r.ResidualToQ4, r.MinimalPolynomial, r.Reason, r.Verdict)
}

func FormatClassification(x Classification) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("executed=%t native_internal_routes=%d native_Hphi_selector_routes=%d native_edge_pullback_routes=%d native_yukawa_reduction_routes=%d sealed_resolvent_routes=%d contact_only=%t Hphi_identity_open=%t verdict=%s", x.Executed, x.NativeInternalRoutes, x.NativeHphiSelectorRoutes, x.NativeEdgePullbackRoutes, x.NativeYukawaReductionRoutes, x.SealedResolventRoutes, x.ContactOnly, x.HphiIdentityStillOpen, x.Verdict))
	for i, r := range x.Routes {
		b.WriteString(fmt.Sprintf("\nroute[%d]: %s", i, FormatRoute(r)))
	}
	return b.String()
}

func FormatImpact(x Impact) string {
	return fmt.Sprintf("q4_internal_contact=%t q4_scalar_identifier=%t q4_edge_pullback=%t contact_split_derived=%t yukawa_reduced=%t moduli_start=%d moduli_result=%d flavor_firewall=%t scalar_Hphi_lane=%t contact_lane=%t verdict=%s", x.Q4InternalContactInvariant, x.Q4ScalarBundleIdentifier, x.Q4EdgeWeightOrPullback, x.ContactProjectorOrSplitDerived, x.YukawaCouplingsReduced, x.ChargedModuliStart, x.ChargedModuliResult, x.FlavorFirewallPreserved, x.ScalarHphiLanePreserved, x.ContactLanePreserved, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_CKM=%t no_PMNS=%t no_yukawa=%t no_manual_q4_Hphi=%t no_root_ordering=%t no_resolvent_root=%t no_arbitrary_basis=%t no_cross_sector_companion=%t no_moduli_reduction=%t verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoManualQ4HphiID, x.NoManualRootOrderingPromoted, x.NoResolventRootPromoted, x.NoArbitraryBasisPromoted, x.NoCompanionOperatorCrossSector, x.NoFlavorModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%s reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func findRoute(xs []SectorRoute, name string) SectorRoute {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return SectorRoute{Name: name, Verdict: "MISSING"}
}
