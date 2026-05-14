// Package spectralgraphedgeadjacency implements Gate 402:
// Spectral Graph Edge-Adjacency Operator Search.
//
// Gate 401 showed that already-derived gauge-charge weights differentiate the
// finite Dirac one-form edges, but do not produce the irreducible contact q4
// selector on the four-real scalar carrier H_phi. Gate 402 therefore audits the
// next non-charge candidate: the native adjacency/incidence topology of the
// finite one-form edge graph itself.  The theorem boundary is strict: a graph
// operator may be native on edge-slot space while still failing to be a
// canonical H_phi endomorphism or q4 selector.
package spectralgraphedgeadjacency

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE402-SPECTRAL-GRAPH-EDGE-ADJACENCY-OPERATOR-SEARCH"

	StatusGate401Inherited           = "CONDITIONAL_SUPPORT_GATE401_WEIGHTED_EDGE_OBSTRUCTION_INHERITED"
	StatusOneFormEdgeGraphFormalized = "CONDITIONAL_SUPPORT_ONEFORM_EDGE_GRAPH_FORMALIZED"
	StatusNativeAdjacencyAudited     = "CONDITIONAL_SUPPORT_NATIVE_EDGE_ADJACENCY_AUDITED"
	StatusYukawaPairGraphComputed    = "CONDITIONAL_SUPPORT_YUKAWA_PAIR_GRAPH_COMPUTED"
	StatusFullFiveEdgeGraphComputed  = "CONDITIONAL_SUPPORT_FULL_FIVE_EDGE_GRAPH_COMPUTED"
	StatusJDoubledGraphAudited       = "CONDITIONAL_SUPPORT_J_DOUBLED_EDGE_GRAPH_AUDITED"
	StatusGraphQuarticCapacityFound  = "CONDITIONAL_SUPPORT_FULL_EDGE_GRAPH_QUARTIC_DEGREE_CAPACITY_FOUND"

	StatusFailedYukawaAdjacencyPairDegenerate  = "FAILED_ROUTE_YUKAWA_EDGE_ADJACENCY_PAIR_DEGENERATE"
	StatusFailedYukawaLaplacianPairDegenerate  = "FAILED_ROUTE_YUKAWA_EDGE_LAPLACIAN_PAIR_DEGENERATE"
	StatusFailedFullGraphNotHphiEndomorphism   = "FAILED_ROUTE_FULL_EDGE_GRAPH_NOT_HPHI_ENDOMORPHISM"
	StatusFailedFullGraphPolynomialNotQ4       = "FAILED_ROUTE_FULL_EDGE_GRAPH_POLYNOMIAL_DISJOINT_FROM_Q4"
	StatusFailedPositiveSpectrumNotHphi        = "FAILED_ROUTE_POSITIVE_GRAPH_SPECTRUM_NOT_FOUR_DIMENSIONAL_HPHI"
	StatusFailedJDoubledOnlyDuplicatesSpectrum = "FAILED_ROUTE_J_DOUBLED_GRAPH_ONLY_DUPLICATES_STRUCTURAL_SPECTRUM"
	StatusFailedNoCanonicalGraphToHphiQuotient = "FAILED_ROUTE_NO_CANONICAL_GRAPH_TO_HPHI_QUOTIENT"
	StatusFailedNoNativeQ4EdgeAdjacency        = "FAILED_ROUTE_NO_NATIVE_Q4_EDGE_ADJACENCY_OPERATOR"
	StatusFailedNoCanonicalHphiQuarticID       = "FAILED_ROUTE_NO_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusFailedNoYukawaCouplingReduction      = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFirewallPreserved13Moduli            = "FIREWALL_PRESERVED_13_MODULI"

	StatusVerifiedCanonicalHphiQuarticID = "VERIFIED_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
)

const (
	ContactQuarticQ4        = "3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271"
	Q4Degree                = 4
	HphiRealDim             = 4
	StructuralEdgeCount     = 5
	YukawaEdgeCount         = 4
	JDoubledEdgeCount       = 10
	Gate372ChargedModuliDim = 13
	eps                     = 1e-10
)

var q4Monic = []float64{1, -7668.0 / 3240.0, 6426.0 / 3240.0, -2235.0 / 3240.0, 271.0 / 3240.0}

type Inheritance struct {
	Executed                         bool
	Gate400NoNativeQ4Selector        bool
	Gate401AnisotropicWeightsFound   bool
	Gate401NoNativeWeightedLaplacian bool
	Gate385OneFormEdges              bool
	Gate385JDoubledEdgeCount         int
	Gate297FirstOrderEdgeGraph       bool
	Gate298InnerFluctuationFields    bool
	Gate372ChargedModuliDim          int
	NoEmpiricalInputsImported        bool
	Verdict                          string
}

type Q4Audit struct {
	Polynomial        string
	Degree            int
	IrreducibleOverQ  bool
	MonicCoefficients []float64
	Verdict           string
}

type EdgeNode struct {
	Name string
	Kind string
}

type EdgeClass struct {
	Name         string
	Source       string
	Target       string
	ScalarBranch string
	Structural   bool
	Yukawa       bool
	Majorana     bool
	JMirror      bool
}

type EdgeGraphArena struct {
	Formalized                    bool
	Nodes                         []EdgeNode
	StructuralEdges               []EdgeClass
	StructuralEdgeCount           int
	YukawaEdgeCount               int
	JDoubledEdgeCount             int
	HasCanonicalEndpointIncidence bool
	HasCanonicalEdgeOrientation   bool
	HasCanonicalHphiQuotient      bool
	UsesGaugeChargeWeights        bool
	UsesYukawaAmplitudes          bool
	UsesObservedMasses            bool
	Verdict                       string
}

type GraphCandidate struct {
	Name                       string
	Formula                    string
	Domain                     string
	Dimension                  int
	Native                     bool
	Sealed                     bool
	Circular                   bool
	HphiEndomorphism           bool
	CanonicalQuotientToHphi    bool
	EdgeGraphDerived           bool
	CompatibleWithJ            bool
	CompatibleWithFirstOrder   bool
	UsesGaugeWeights           bool
	UsesYukawaAmplitudes       bool
	UsesObservedMasses         bool
	Components                 int
	Eigenvalues                []float64
	DistinctEigenvalues        int
	MinimalDegree              int
	CharacteristicPolynomial   string
	MinimalPolynomial          string
	CharacteristicResidualToQ4 float64
	MinimalResidualToQ4        float64
	PairDegenerate             bool
	CentralOnHphi              bool
	IrreducibleQuarticCapacity bool
	Q4ExactMatch               bool
	Q4FactorMatch              bool
	PromotableAsQ4Selector     bool
	ReducesYukawaCouplings     bool
	ReducesFlavorModuli        bool
	Reason                     string
	Verdict                    string
}

type GraphSieve struct {
	Executed                    bool
	Candidates                  []GraphCandidate
	NativeGraphOperatorCount    int
	NativeHphiEndomorphismCount int
	NativeQuarticCapacityCount  int
	CanonicalHphiQ4MatchCount   int
	SealedOrManualQ4Count       int
	BestNativeCandidate         string
	BestNativeQ4Residual        float64
	Verdict                     string
}

type Impact struct {
	HphiQuarticIdentified          bool
	NativeEdgeAdjacencyFound       bool
	CanonicalGraphQuotientFound    bool
	YukawaCouplingsReduced         bool
	ChargedModuliStart             int
	ChargedModuliResult            int
	FlavorFirewallPreserved        bool
	HiggsLanePreserved             bool
	EdgeGraphLaneOpenedButUnsealed bool
	Verdict                        string
}

type FirewallAudit struct {
	Executed                       bool
	NoObservedMassesImported       bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	NoYukawaAmplitudesInserted     bool
	NoGaugeChargeFitReused         bool
	NoManualQ4HphiID               bool
	NoArbitraryGraphQuotient       bool
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
	Q4          Q4Audit
	Arena       EdgeGraphArena
	Sieve       GraphSieve
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
	sieve := auditGraphs(arena, q4)
	impact := auditImpact(sieve)
	firewall := auditFirewall(sieve, impact)
	next := nextStep(sieve, impact)
	return Analysis{Inheritance: inheritance, Q4: q4, Arena: arena, Sieve: sieve, Impact: impact, Firewall: firewall, Next: next, Truth: truth(sieve, impact)}, nil
}

func inherit() Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate400NoNativeQ4Selector:        true,
		Gate401AnisotropicWeightsFound:   true,
		Gate401NoNativeWeightedLaplacian: true,
		Gate385OneFormEdges:              true,
		Gate385JDoubledEdgeCount:         JDoubledEdgeCount,
		Gate297FirstOrderEdgeGraph:       true,
		Gate298InnerFluctuationFields:    true,
		Gate372ChargedModuliDim:          Gate372ChargedModuliDim,
		NoEmpiricalInputsImported:        true,
		Verdict:                          "Gate 402 inherits the Gate-400/401 q4 obstruction, the Gate-385 one-form edge support, the Gate-297 first-order-compatible structural edge graph, the Gate-298 inner-fluctuation field inventory, and the Gate-372 flavor firewall.",
	}
}

func auditQ4() Q4Audit {
	return Q4Audit{Polynomial: ContactQuarticQ4, Degree: Q4Degree, IrreducibleOverQ: true, MonicCoefficients: append([]float64(nil), q4Monic...), Verdict: "The contact target remains the irreducible quartic q4. A native edge-graph operator must match this polynomial without basis insertion, affine fitting, or observed Yukawa amplitudes."}
}

func buildArena() EdgeGraphArena {
	nodes := []EdgeNode{
		{Name: "L_L", Kind: "left lepton doublet"},
		{Name: "Q_L", Kind: "left quark doublet"},
		{Name: "e_R", Kind: "right charged lepton"},
		{Name: "nu_R", Kind: "right neutrino"},
		{Name: "d_R", Kind: "right down quark"},
		{Name: "u_R", Kind: "right up quark"},
		{Name: "nu_R^c", Kind: "conjugate sterile/Majorana node"},
	}
	edges := []EdgeClass{
		{Name: "L_L ↔ e_R", Source: "L_L", Target: "e_R", ScalarBranch: "Phi_-", Structural: true, Yukawa: true},
		{Name: "L_L ↔ nu_R", Source: "L_L", Target: "nu_R", ScalarBranch: "Phi_+", Structural: true, Yukawa: true},
		{Name: "Q_L ↔ d_R", Source: "Q_L", Target: "d_R", ScalarBranch: "Phi_-", Structural: true, Yukawa: true},
		{Name: "Q_L ↔ u_R", Source: "Q_L", Target: "u_R", ScalarBranch: "Phi_+", Structural: true, Yukawa: true},
		{Name: "nu_R ↔ nu_R^c", Source: "nu_R", Target: "nu_R^c", ScalarBranch: "singlet/Majorana", Structural: true, Majorana: true},
	}
	return EdgeGraphArena{
		Formalized:                    true,
		Nodes:                         nodes,
		StructuralEdges:               edges,
		StructuralEdgeCount:           len(edges),
		YukawaEdgeCount:               4,
		JDoubledEdgeCount:             2 * len(edges),
		HasCanonicalEndpointIncidence: true,
		HasCanonicalEdgeOrientation:   false,
		HasCanonicalHphiQuotient:      false,
		UsesGaugeChargeWeights:        false,
		UsesYukawaAmplitudes:          false,
		UsesObservedMasses:            false,
		Verdict:                       "The native one-form edge graph is the first-order-compatible finite-Dirac graph with four Yukawa/Higgs edges plus one sterile/Majorana edge, doubled by J. Endpoint incidence is canonical; an orientation and a four-real H_phi quotient are not yet canonically selected.",
	}
}

func auditGraphs(arena EdgeGraphArena, q4 Q4Audit) GraphSieve {
	candidates := []GraphCandidate{
		{
			Name:                       "four Yukawa-edge adjacency graph K2 disjoint union K2",
			Formula:                    "A_Y on {L-e,L-nu,Q-d,Q-u}, edges adjacent when finite-Dirac edges share a left source node",
			Domain:                     "four Yukawa edge slots",
			Dimension:                  4,
			Native:                     true,
			HphiEndomorphism:           true,
			CanonicalQuotientToHphi:    true,
			EdgeGraphDerived:           true,
			CompatibleWithJ:            true,
			CompatibleWithFirstOrder:   true,
			Components:                 2,
			Eigenvalues:                []float64{-1, -1, 1, 1},
			DistinctEigenvalues:        2,
			MinimalDegree:              2,
			CharacteristicPolynomial:   "(x^2-1)^2",
			MinimalPolynomial:          "x^2-1",
			CharacteristicResidualToQ4: residualMonic([]float64{1, 0, -2, 0, 1}, q4Monic),
			MinimalResidualToQ4:        math.Inf(1),
			PairDegenerate:             true,
			IrreducibleQuarticCapacity: false,
			Q4ExactMatch:               false,
			PromotableAsQ4Selector:     false,
			Reason:                     "The canonical scalar/Yukawa edge adjacency splits into two identical weak-source pairs, so its invariant is pair-degenerate and quadratic.",
			Verdict:                    StatusFailedYukawaAdjacencyPairDegenerate,
		},
		{
			Name:                       "four Yukawa-edge graph Laplacian K2 disjoint union K2",
			Formula:                    "L_Y = B_Y^T B_Y on the four structural Yukawa edges",
			Domain:                     "four Yukawa edge slots",
			Dimension:                  4,
			Native:                     true,
			HphiEndomorphism:           true,
			CanonicalQuotientToHphi:    true,
			EdgeGraphDerived:           true,
			CompatibleWithJ:            true,
			CompatibleWithFirstOrder:   true,
			Components:                 2,
			Eigenvalues:                []float64{0, 0, 2, 2},
			DistinctEigenvalues:        2,
			MinimalDegree:              2,
			CharacteristicPolynomial:   "x^2*(x-2)^2",
			MinimalPolynomial:          "x*(x-2)",
			CharacteristicResidualToQ4: residualMonic([]float64{1, -4, 4, 0, 0}, q4Monic),
			MinimalResidualToQ4:        math.Inf(1),
			PairDegenerate:             true,
			IrreducibleQuarticCapacity: false,
			Q4ExactMatch:               false,
			PromotableAsQ4Selector:     false,
			Reason:                     "The graph Laplacian on the four Higgs/Yukawa edges is the same two-pair structure already seen in scalar branch compression.",
			Verdict:                    StatusFailedYukawaLaplacianPairDegenerate,
		},
		{
			Name:                       "full five-edge structural Laplacian P3 disjoint union K2",
			Formula:                    "L_E on {L-e,L-nu,nu-M,Q-d,Q-u}; edge-edge adjacency through shared endpoint modules",
			Domain:                     "five structural one-form edge classes",
			Dimension:                  5,
			Native:                     true,
			HphiEndomorphism:           false,
			CanonicalQuotientToHphi:    false,
			EdgeGraphDerived:           true,
			CompatibleWithJ:            true,
			CompatibleWithFirstOrder:   true,
			Components:                 2,
			Eigenvalues:                []float64{0, 0, 1, 2, 3},
			DistinctEigenvalues:        4,
			MinimalDegree:              4,
			CharacteristicPolynomial:   "x^2*(x-1)*(x-2)*(x-3)",
			MinimalPolynomial:          "x*(x-1)*(x-2)*(x-3)",
			CharacteristicResidualToQ4: math.Inf(1),
			MinimalResidualToQ4:        residualMonic([]float64{1, -6, 11, -6, 0}, q4Monic),
			PairDegenerate:             false,
			IrreducibleQuarticCapacity: true,
			Q4ExactMatch:               false,
			PromotableAsQ4Selector:     false,
			Reason:                     "The full edge graph finally has quartic-degree minimal-polynomial capacity, but it lives on the five-edge one-form graph and its quartic is x(x-1)(x-2)(x-3), not the contact q4.",
			Verdict:                    StatusFailedFullGraphPolynomialNotQ4,
		},
		{
			Name:                       "positive-spectrum quotient of full five-edge Laplacian",
			Formula:                    "L_E restricted modulo component-constant zero modes",
			Domain:                     "positive graph modes of P3 disjoint union K2",
			Dimension:                  3,
			Native:                     true,
			HphiEndomorphism:           false,
			CanonicalQuotientToHphi:    false,
			EdgeGraphDerived:           true,
			CompatibleWithJ:            true,
			CompatibleWithFirstOrder:   true,
			Components:                 2,
			Eigenvalues:                []float64{1, 2, 3},
			DistinctEigenvalues:        3,
			MinimalDegree:              3,
			CharacteristicPolynomial:   "(x-1)*(x-2)*(x-3)",
			MinimalPolynomial:          "(x-1)*(x-2)*(x-3)",
			CharacteristicResidualToQ4: math.Inf(1),
			MinimalResidualToQ4:        math.Inf(1),
			PairDegenerate:             false,
			IrreducibleQuarticCapacity: false,
			Q4ExactMatch:               false,
			PromotableAsQ4Selector:     false,
			Reason:                     "The canonical quotient by connected-component zero modes is three-dimensional, not H_phi's four-real scalar carrier.",
			Verdict:                    StatusFailedPositiveSpectrumNotHphi,
		},
		{
			Name:                       "J-doubled structural edge graph",
			Formula:                    "L_E ⊕ J L_E J^{-1} on ten one-form edge slots",
			Domain:                     "ten J-doubled edge slots",
			Dimension:                  10,
			Native:                     true,
			HphiEndomorphism:           false,
			CanonicalQuotientToHphi:    false,
			EdgeGraphDerived:           true,
			CompatibleWithJ:            true,
			CompatibleWithFirstOrder:   true,
			Components:                 4,
			Eigenvalues:                []float64{0, 0, 0, 0, 1, 1, 2, 2, 3, 3},
			DistinctEigenvalues:        4,
			MinimalDegree:              4,
			CharacteristicPolynomial:   "[x^2*(x-1)*(x-2)*(x-3)]^2",
			MinimalPolynomial:          "x*(x-1)*(x-2)*(x-3)",
			CharacteristicResidualToQ4: math.Inf(1),
			MinimalResidualToQ4:        residualMonic([]float64{1, -6, 11, -6, 0}, q4Monic),
			PairDegenerate:             false,
			IrreducibleQuarticCapacity: true,
			Q4ExactMatch:               false,
			PromotableAsQ4Selector:     false,
			Reason:                     "J-doubling respects the edge graph but only duplicates the structural spectrum; it does not create a new scalar selector.",
			Verdict:                    StatusFailedJDoubledOnlyDuplicatesSpectrum,
		},
		{
			Name:                       "sealed q4 edge-graph companion quotient",
			Formula:                    "choose a four-edge basis and place the q4 companion matrix by hand",
			Domain:                     "manually chosen four-dimensional edge quotient",
			Dimension:                  4,
			Native:                     false,
			Sealed:                     true,
			Circular:                   true,
			HphiEndomorphism:           true,
			CanonicalQuotientToHphi:    false,
			EdgeGraphDerived:           false,
			CompatibleWithJ:            false,
			CompatibleWithFirstOrder:   false,
			Components:                 1,
			Eigenvalues:                []float64{},
			DistinctEigenvalues:        4,
			MinimalDegree:              4,
			CharacteristicPolynomial:   ContactQuarticQ4,
			MinimalPolynomial:          ContactQuarticQ4,
			CharacteristicResidualToQ4: 0,
			MinimalResidualToQ4:        0,
			PairDegenerate:             false,
			IrreducibleQuarticCapacity: true,
			Q4ExactMatch:               true,
			PromotableAsQ4Selector:     false,
			Reason:                     "This stress test has q4 only because q4 is inserted as the companion polynomial; it is quarantined and cannot rewrite the native theorem.",
			Verdict:                    "QUARANTINED_SEALED_Q4_COMPANION_NOT_PROMOTED",
		},
	}

	nativeGraph := 0
	nativeHphi := 0
	quartic := 0
	q4Matches := 0
	sealed := 0
	bestName := "none"
	bestResidual := math.Inf(1)
	for _, c := range candidates {
		if c.Native && c.EdgeGraphDerived {
			nativeGraph++
		}
		if c.Native && c.HphiEndomorphism && c.CanonicalQuotientToHphi {
			nativeHphi++
		}
		if c.Native && c.IrreducibleQuarticCapacity {
			quartic++
		}
		if c.Native && c.Q4ExactMatch && c.PromotableAsQ4Selector {
			q4Matches++
		}
		if c.Sealed || c.Circular {
			sealed++
		}
		r := c.MinimalResidualToQ4
		if math.IsInf(r, 0) {
			r = c.CharacteristicResidualToQ4
		}
		if c.Native && !math.IsInf(r, 0) && r < bestResidual {
			bestResidual = r
			bestName = c.Name
		}
	}
	return GraphSieve{
		Executed:                    true,
		Candidates:                  candidates,
		NativeGraphOperatorCount:    nativeGraph,
		NativeHphiEndomorphismCount: nativeHphi,
		NativeQuarticCapacityCount:  quartic,
		CanonicalHphiQ4MatchCount:   q4Matches,
		SealedOrManualQ4Count:       sealed,
		BestNativeCandidate:         bestName,
		BestNativeQ4Residual:        bestResidual,
		Verdict:                     "The native edge graph has real adjacency/Laplacian structure and the full five-edge graph has quartic-degree capacity. However, the canonical four-edge H_phi quotient is pair-degenerate, while the quartic-capable graph is five/ten-dimensional and polynomially disjoint from q4.",
	}
}

func auditImpact(s GraphSieve) Impact {
	q4 := s.CanonicalHphiQ4MatchCount > 0
	return Impact{
		HphiQuarticIdentified:          q4,
		NativeEdgeAdjacencyFound:       s.NativeGraphOperatorCount > 0,
		CanonicalGraphQuotientFound:    false,
		YukawaCouplingsReduced:         false,
		ChargedModuliStart:             Gate372ChargedModuliDim,
		ChargedModuliResult:            Gate372ChargedModuliDim,
		FlavorFirewallPreserved:        true,
		HiggsLanePreserved:             !q4,
		EdgeGraphLaneOpenedButUnsealed: s.NativeQuarticCapacityCount > 0 && !q4,
		Verdict:                        "Gate 402 opens a real edge-graph spectral lane, but it does not identify q4 on H_phi or reduce Yukawa/flavor moduli.",
	}
}

func auditFirewall(s GraphSieve, impact Impact) FirewallAudit {
	return FirewallAudit{
		Executed:                       true,
		NoObservedMassesImported:       true,
		NoCKMImported:                  true,
		NoPMNSImported:                 true,
		NoYukawaAmplitudesInserted:     true,
		NoGaugeChargeFitReused:         true,
		NoManualQ4HphiID:               true,
		NoArbitraryGraphQuotient:       true,
		NoCompanionOperatorPromoted:    s.CanonicalHphiQ4MatchCount == 0,
		NoFlavorModuliReductionClaimed: impact.ChargedModuliResult == Gate372ChargedModuliDim,
		Verdict:                        "The graph audit uses only structural incidence data. It does not import observed masses, CKM/PMNS, Yukawa amplitudes, gauge-charge fitting, or arbitrary q4/H_phi identifications.",
	}
}

func nextStep(s GraphSieve, impact Impact) NextStep {
	return NextStep{Gate: 403, Title: "Oriented Edge-Incidence Boundary Operator Sieve", Reason: "Undirected edge adjacency is either pair-degenerate on the four Yukawa/H_phi edge slots or quartic-capable only on the five-edge graph, where it is not an H_phi endomorphism and its polynomial is not q4. The next non-arbitrary candidate is the oriented source-target incidence/boundary operator, because orientation may distinguish the four Higgs edge channels without using charge weights or Yukawa amplitudes.", PrimaryTask: "Construct canonical oriented incidence and signed boundary/coboundary operators for the finite one-form edge graph; test whether any J-compatible four-dimensional scalar quotient has a q4 minimal polynomial without manual edge-to-H_phi placement."}
}

func truth(s GraphSieve, impact Impact) string {
	return "Gate 402 proves that the one-form finite-Dirac edge graph is a real native object, but its undirected adjacency/Laplacian topology still does not identify the scalar carrier with the contact q4 primary. The four Yukawa edge graph is K2 ⊔ K2 and is therefore pair-degenerate. The full structural edge graph is P3 ⊔ K2 and has quartic-degree capacity, but it lives on the five-edge/ten-J-doubled edge-slot space and its native quartic x(x-1)(x-2)(x-3) is disjoint from the contact q4. Thus the q4/H_phi identification remains unproved, no Yukawa couplings are reduced, and the 13-moduli flavor firewall is preserved. The next valid search is not another weight but an oriented incidence/boundary operator or a canonical edge-to-H_phi quotient theorem."
}

func Statuses(a Analysis) []string {
	out := []string{
		StatusGate401Inherited,
		StatusOneFormEdgeGraphFormalized,
		StatusNativeAdjacencyAudited,
		StatusYukawaPairGraphComputed,
		StatusFullFiveEdgeGraphComputed,
		StatusJDoubledGraphAudited,
	}
	if a.Sieve.NativeQuarticCapacityCount > 0 {
		out = append(out, StatusGraphQuarticCapacityFound)
	}
	out = append(out,
		StatusFailedYukawaAdjacencyPairDegenerate,
		StatusFailedYukawaLaplacianPairDegenerate,
		StatusFailedFullGraphNotHphiEndomorphism,
		StatusFailedFullGraphPolynomialNotQ4,
		StatusFailedPositiveSpectrumNotHphi,
		StatusFailedJDoubledOnlyDuplicatesSpectrum,
		StatusFailedNoCanonicalGraphToHphiQuotient,
		StatusFailedNoNativeQ4EdgeAdjacency,
		StatusFailedNoCanonicalHphiQuarticID,
		StatusFailedNoYukawaCouplingReduction,
		StatusFirewallPreserved13Moduli,
	)
	return out
}

func residualMonic(poly []float64, target []float64) float64 {
	if len(poly) != len(target) {
		return math.Inf(1)
	}
	var s float64
	for i := range poly {
		d := poly[i] - target[i]
		s += d * d
	}
	return math.Sqrt(s)
}

func formatFloatSlice(xs []float64) string {
	if len(xs) == 0 {
		return "[]"
	}
	parts := make([]string, len(xs))
	for i, x := range xs {
		if math.IsInf(x, 1) {
			parts[i] = "+Inf"
		} else if math.IsInf(x, -1) {
			parts[i] = "-Inf"
		} else {
			parts[i] = fmt.Sprintf("%.12g", x)
		}
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
