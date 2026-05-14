// Package orientededgeincidence implements Gate 403:
// Oriented Edge-Incidence Boundary Operator Sieve.
//
// Gate 402 showed that the undirected one-form edge graph is native and has
// quartic-degree capacity only on the full five-edge graph, but its polynomial
// is disjoint from the contact q4 target and it is not a canonical H_phi
// endomorphism. Gate 403 upgrades endpoint adjacency to a signed chiral
// incidence/boundary operator d. The theorem boundary is strict: orientation is
// promotable only if d^T d or d^†d gives a native four-real scalar operator with
// the q4 invariant polynomial. Ordinary graph-orientation signs are expected to
// disappear up to signed edge-basis conjugacy, so the gate explicitly audits
// that cancellation rather than treating orientation as free spectral data.
package orientededgeincidence

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE403-ORIENTED-EDGE-INCIDENCE-BOUNDARY-OPERATOR-SIEVE"

	StatusGate402Inherited                  = "CONDITIONAL_SUPPORT_GATE402_EDGE_GRAPH_OBSTRUCTION_INHERITED"
	StatusOrientedBoundaryFormalized        = "CONDITIONAL_SUPPORT_ORIENTED_BOUNDARY_FORMALIZED"
	StatusMajoranaOrientationAudited        = "CONDITIONAL_SUPPORT_MAJORANA_ORIENTATION_AUDITED"
	StatusOrientedLaplacianConstructed      = "CONDITIONAL_SUPPORT_ORIENTED_LAPLACIAN_CONSTRUCTED"
	StatusFullIncidenceRadicalSpectrumFound = "CONDITIONAL_SUPPORT_FULL_ORIENTED_INCIDENCE_RADICAL_SPECTRUM_FOUND"

	StatusFailedOrientationSignsCancel       = "FAILED_ROUTE_ORIENTATION_SIGNS_CANCEL_IN_DTD"
	StatusFailedYukawaOrientedPairDegenerate = "FAILED_ROUTE_YUKAWA_ORIENTED_LAPLACIAN_PAIR_DEGENERATE"
	StatusFailedFullOrientedNotHphi          = "FAILED_ROUTE_FULL_ORIENTED_LAPLACIAN_NOT_HPHI_ENDOMORPHISM"
	StatusFailedFullOrientedPolynomialNotQ4  = "FAILED_ROUTE_ORIENTED_POLYNOMIAL_DISJOINT_FROM_Q4"
	StatusFailedMajoranaTwistNotCanonical    = "FAILED_ROUTE_MAJORANA_TWIST_NOT_CANONICAL_OR_SPECTRALLY_NEW"
	StatusFailedNoCanonicalOrientedQuotient  = "FAILED_ROUTE_NO_CANONICAL_ORIENTED_EDGE_TO_HPHI_QUOTIENT"
	StatusFailedNoNativeOrientedQ4Selector   = "FAILED_ROUTE_NO_NATIVE_ORIENTED_Q4_SELECTOR"
	StatusFailedNoCanonicalHphiQuarticID     = "FAILED_ROUTE_NO_CANONICAL_HPHI_QUARTIC_IDENTIFICATION"
	StatusFailedNoYukawaCouplingReduction    = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFirewallPreserved13Moduli          = "FIREWALL_PRESERVED_13_MODULI"

	StatusVerifiedOrientedGraphQ4ID = "CONDITIONAL_SUPPORT_ORIENTED_GRAPH_Q4_IDENTIFICATION"
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
	Gate399QuaternionicPolynomialNo bool
	Gate400NoMixedEdgeQ4            bool
	Gate401ChargeWeightsDisjoint    bool
	Gate402UndirectedGraphNative    bool
	Gate402FullGraphQuarticCapacity bool
	Gate402NoGraphQ4                bool
	Gate385OneFormEdges             bool
	Gate385JDoubledEdgeCount        int
	Gate297FirstOrderEdgeGraph      bool
	Gate372ChargedModuliDim         int
	NoEmpiricalInputsImported       bool
	Verdict                         string
}

type Q4Audit struct {
	Polynomial        string
	Degree            int
	IrreducibleOverQ  bool
	MonicCoefficients []float64
	RootSummary       string
	Verdict           string
}

type OrientedNode struct {
	Name      string
	Chirality string
	Kind      string
}

type OrientedEdge struct {
	Name                 string
	Source               string
	Target               string
	OrientationRule      string
	ScalarBranch         string
	Yukawa               bool
	Majorana             bool
	JMirror              bool
	OrientationNative    bool
	OrientationAmbiguous bool
}

type BoundaryArena struct {
	Formalized                   bool
	Nodes                        []OrientedNode
	Edges                        []OrientedEdge
	StructuralEdgeCount          int
	YukawaEdgeCount              int
	JDoubledEdgeCount            int
	ChiralOrientationAvailable   bool
	MajoranaOrientationCanonical bool
	MajoranaTwistPossible        bool
	HasCanonicalHphiQuotient     bool
	UsesGaugeChargeWeights       bool
	UsesYukawaAmplitudes         bool
	UsesObservedMasses           bool
	Verdict                      string
}

type BoundaryCandidate struct {
	Name                             string
	Formula                          string
	Domain                           string
	Dimension                        int
	Native                           bool
	Sealed                           bool
	Circular                         bool
	BoundaryDerived                  bool
	HphiEndomorphism                 bool
	CanonicalQuotientToHphi          bool
	CompatibleWithChirality          bool
	CompatibleWithJ                  bool
	CompatibleWithFirstOrder         bool
	MajoranaOrientationUsed          string
	OrientationChoiceAffectsSpectrum bool
	OrientationSignsCancel           bool
	UsesGaugeWeights                 bool
	UsesYukawaAmplitudes             bool
	UsesObservedMasses               bool
	Eigenvalues                      []string
	DistinctEigenvalues              int
	MinimalDegree                    int
	CharacteristicPolynomial         string
	MinimalPolynomial                string
	CharacteristicResidualToQ4       float64
	MinimalResidualToQ4              float64
	PairDegenerate                   bool
	IrreducibleQuarticCapacity       bool
	Q4ExactMatch                     bool
	Q4FactorMatch                    bool
	PromotableAsQ4Selector           bool
	ReducesYukawaCouplings           bool
	ReducesFlavorModuli              bool
	Reason                           string
	Verdict                          string
}

type BoundarySieve struct {
	Executed                    bool
	Candidates                  []BoundaryCandidate
	NativeBoundaryOperatorCount int
	NativeHphiEndomorphismCount int
	NativeQuarticCapacityCount  int
	CanonicalHphiQ4MatchCount   int
	OrientationInvariantCount   int
	SealedOrManualQ4Count       int
	BestNativeCandidate         string
	BestNativeQ4Residual        float64
	Verdict                     string
}

type Impact struct {
	HphiQuarticIdentified          bool
	NativeBoundaryOperatorFound    bool
	CanonicalBoundaryQuotientFound bool
	OrientedIncidenceLaneOpened    bool
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
	NoGaugeChargeFitReused         bool
	NoManualQ4HphiID               bool
	NoArbitraryBoundaryQuotient    bool
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
	Arena       BoundaryArena
	Sieve       BoundarySieve
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
	sieve := auditBoundaryOperators(arena, q4)
	impact := auditImpact(sieve)
	firewall := auditFirewall(impact)
	next := nextStep(sieve, impact)
	return Analysis{Inheritance: inheritance, Q4: q4, Arena: arena, Sieve: sieve, Impact: impact, Firewall: firewall, Next: next, Truth: truth(sieve, impact)}, nil
}

func inherit() Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate399QuaternionicPolynomialNo: true,
		Gate400NoMixedEdgeQ4:            true,
		Gate401ChargeWeightsDisjoint:    true,
		Gate402UndirectedGraphNative:    true,
		Gate402FullGraphQuarticCapacity: true,
		Gate402NoGraphQ4:                true,
		Gate385OneFormEdges:             true,
		Gate385JDoubledEdgeCount:        JDoubledEdgeCount,
		Gate297FirstOrderEdgeGraph:      true,
		Gate372ChargedModuliDim:         Gate372ChargedModuliDim,
		NoEmpiricalInputsImported:       true,
		Verdict:                         "Gate 403 inherits the Gate-399 quaternionic polynomial obstruction, Gate-400 mixed edge/Laplacian obstruction, Gate-401 charge-weight obstruction, Gate-402 native edge graph with no q4 selector, Gate-385 one-form edge ledger, Gate-297 first-order edge graph, and Gate-372 flavor firewall.",
	}
}

func auditQ4() Q4Audit {
	return Q4Audit{Polynomial: ContactQuarticQ4, Degree: Q4Degree, IrreducibleOverQ: true, MonicCoefficients: append([]float64(nil), q4Monic...), RootSummary: "four distinct positive algebraic roots approximately 0.2839, 0.4411, 0.7441, 0.8975", Verdict: "The target remains the irreducible contact quartic. An oriented incidence theorem must match q4 natively, not by orienting signs to fit coefficients."}
}

func buildArena() BoundaryArena {
	nodes := []OrientedNode{
		{Name: "L_L", Chirality: "left", Kind: "lepton weak doublet"},
		{Name: "Q_L", Chirality: "left", Kind: "quark weak doublet"},
		{Name: "e_R", Chirality: "right", Kind: "charged-lepton singlet"},
		{Name: "nu_R", Chirality: "right", Kind: "neutrino singlet"},
		{Name: "d_R", Chirality: "right", Kind: "down-quark singlet"},
		{Name: "u_R", Chirality: "right", Kind: "up-quark singlet"},
		{Name: "nu_R^c", Chirality: "J-conjugate/right-opposite", Kind: "sterile/Majorana conjugate node"},
	}
	edges := []OrientedEdge{
		{Name: "L_L -> e_R", Source: "L_L", Target: "e_R", OrientationRule: "gamma_F left-to-right Yukawa edge", ScalarBranch: "Phi_-", Yukawa: true, OrientationNative: true},
		{Name: "L_L -> nu_R", Source: "L_L", Target: "nu_R", OrientationRule: "gamma_F left-to-right Yukawa edge", ScalarBranch: "Phi_+", Yukawa: true, OrientationNative: true},
		{Name: "Q_L -> d_R", Source: "Q_L", Target: "d_R", OrientationRule: "gamma_F left-to-right Yukawa edge", ScalarBranch: "Phi_-", Yukawa: true, OrientationNative: true},
		{Name: "Q_L -> u_R", Source: "Q_L", Target: "u_R", OrientationRule: "gamma_F left-to-right Yukawa edge", ScalarBranch: "Phi_+", Yukawa: true, OrientationNative: true},
		{Name: "nu_R -> nu_R^c", Source: "nu_R", Target: "nu_R^c", OrientationRule: "J-real Majorana orientation; reversal is J-equivalent", ScalarBranch: "singlet/Majorana", Majorana: true, OrientationNative: false, OrientationAmbiguous: true},
	}
	return BoundaryArena{
		Formalized:                   true,
		Nodes:                        nodes,
		Edges:                        edges,
		StructuralEdgeCount:          len(edges),
		YukawaEdgeCount:              YukawaEdgeCount,
		JDoubledEdgeCount:            2 * len(edges),
		ChiralOrientationAvailable:   true,
		MajoranaOrientationCanonical: false,
		MajoranaTwistPossible:        true,
		HasCanonicalHphiQuotient:     false,
		UsesGaugeChargeWeights:       false,
		UsesYukawaAmplitudes:         false,
		UsesObservedMasses:           false,
		Verdict:                      "Chiral orientation canonically orients the four Yukawa edges left-to-right. The neutral Majorana edge is J-real, so its arrow can be chosen for incidence bookkeeping, but reversal or a unit phase is not new spectral data for d^T d or d^†d. A canonical four-real H_phi quotient is still absent.",
	}
}

func auditBoundaryOperators(arena BoundaryArena, q4 Q4Audit) BoundarySieve {
	candidates := []BoundaryCandidate{
		{
			Name:                             "four Yukawa oriented edge Gram d_Y^T d_Y",
			Formula:                          "d_Y has columns L_L->e_R, L_L->nu_R, Q_L->d_R, Q_L->u_R; Δ_Y=d_Y^T d_Y",
			Domain:                           "four oriented Yukawa edge slots",
			Dimension:                        HphiRealDim,
			Native:                           true,
			BoundaryDerived:                  true,
			HphiEndomorphism:                 true,
			CanonicalQuotientToHphi:          true,
			CompatibleWithChirality:          true,
			CompatibleWithJ:                  true,
			CompatibleWithFirstOrder:         true,
			OrientationChoiceAffectsSpectrum: false,
			OrientationSignsCancel:           true,
			Eigenvalues:                      []string{"1", "1", "3", "3"},
			DistinctEigenvalues:              2,
			MinimalDegree:                    2,
			CharacteristicPolynomial:         "(x-1)^2*(x-3)^2 = x^4 - 8*x^3 + 22*x^2 - 24*x + 9",
			MinimalPolynomial:                "(x-1)*(x-3) = x^2 - 4*x + 3",
			CharacteristicResidualToQ4:       residualMonic([]float64{1, -8, 22, -24, 9}, q4Monic),
			MinimalResidualToQ4:              math.Inf(1),
			PairDegenerate:                   true,
			Q4ExactMatch:                     false,
			PromotableAsQ4Selector:           false,
			Reason:                           "The chiral boundary operator is native, but d^T d keeps the two weak-source pairs; reversing any edge only conjugates by a signed edge basis and does not change the spectrum.",
			Verdict:                          StatusFailedYukawaOrientedPairDegenerate,
		},
		{
			Name:                             "full five-edge oriented incidence Gram d_E^T d_E",
			Formula:                          "Δ_E=d_E^T d_E for L-e, L-nu, nu-nu^c, Q-d, Q-u with arbitrary J-equivalent Majorana arrow",
			Domain:                           "five oriented structural one-form edge slots",
			Dimension:                        StructuralEdgeCount,
			Native:                           true,
			BoundaryDerived:                  true,
			HphiEndomorphism:                 false,
			CanonicalQuotientToHphi:          false,
			CompatibleWithChirality:          true,
			CompatibleWithJ:                  true,
			CompatibleWithFirstOrder:         true,
			MajoranaOrientationUsed:          "nu_R -> nu_R^c; reversal is signed-column conjugate",
			OrientationChoiceAffectsSpectrum: false,
			OrientationSignsCancel:           true,
			Eigenvalues:                      []string{"1", "2-sqrt(2)", "2", "3", "2+sqrt(2)"},
			DistinctEigenvalues:              5,
			MinimalDegree:                    5,
			CharacteristicPolynomial:         "(x-1)*(x-2)*(x-3)*(x^2 - 4*x + 2) = x^5 - 10*x^4 + 37*x^3 - 62*x^2 + 46*x - 12",
			MinimalPolynomial:                "same as characteristic polynomial; five distinct edge-space eigenvalues",
			CharacteristicResidualToQ4:       math.Inf(1),
			MinimalResidualToQ4:              math.Inf(1),
			PairDegenerate:                   false,
			IrreducibleQuarticCapacity:       false,
			Q4ExactMatch:                     false,
			PromotableAsQ4Selector:           false,
			Reason:                           "Orientation changes the undirected graph Laplacian to an incidence Gram with radical eigenvalues, but it is five-dimensional edge-slot data, not a four-real H_phi endomorphism, and its minimal polynomial has degree five rather than q4 degree four.",
			Verdict:                          StatusFailedFullOrientedNotHphi,
		},
		{
			Name:                             "noncanonical four-mode quotient of full oriented incidence Gram",
			Formula:                          "drop one full-edge mode to force a four-dimensional quotient, e.g. keep {1,3,2-sqrt(2),2+sqrt(2)}",
			Domain:                           "manual four-mode quotient of oriented edge space",
			Dimension:                        HphiRealDim,
			Native:                           false,
			Sealed:                           true,
			Circular:                         true,
			BoundaryDerived:                  true,
			HphiEndomorphism:                 false,
			CanonicalQuotientToHphi:          false,
			CompatibleWithChirality:          true,
			CompatibleWithJ:                  false,
			CompatibleWithFirstOrder:         false,
			OrientationChoiceAffectsSpectrum: false,
			OrientationSignsCancel:           true,
			Eigenvalues:                      []string{"1", "3", "2-sqrt(2)", "2+sqrt(2)"},
			DistinctEigenvalues:              4,
			MinimalDegree:                    4,
			CharacteristicPolynomial:         "(x-1)*(x-3)*(x^2 - 4*x + 2) = x^4 - 8*x^3 + 21*x^2 - 20*x + 6",
			MinimalPolynomial:                "same as characteristic polynomial on the chosen quotient",
			CharacteristicResidualToQ4:       residualMonic([]float64{1, -8, 21, -20, 6}, q4Monic),
			MinimalResidualToQ4:              residualMonic([]float64{1, -8, 21, -20, 6}, q4Monic),
			PairDegenerate:                   false,
			IrreducibleQuarticCapacity:       true,
			Q4ExactMatch:                     false,
			PromotableAsQ4Selector:           false,
			Reason:                           "A four-dimensional quotient can be forced, but choosing which edge mode to remove is not supplied by A_F, J, first-order, H_phi, or the contact vacuum. Its quartic is also disjoint from q4.",
			Verdict:                          StatusFailedNoCanonicalOrientedQuotient,
		},
		{
			Name:                             "J-twisted complex Majorana boundary d^†d",
			Formula:                          "place a unit phase on the neutral Majorana column and compute the Hermitian Gram d^†d",
			Domain:                           "five complex structural one-form edge slots",
			Dimension:                        StructuralEdgeCount,
			Native:                           false,
			Sealed:                           true,
			BoundaryDerived:                  true,
			HphiEndomorphism:                 false,
			CanonicalQuotientToHphi:          false,
			CompatibleWithChirality:          true,
			CompatibleWithJ:                  true,
			CompatibleWithFirstOrder:         true,
			MajoranaOrientationUsed:          "unit phase twist on Majorana edge",
			OrientationChoiceAffectsSpectrum: false,
			OrientationSignsCancel:           true,
			Eigenvalues:                      []string{"1", "2-sqrt(2)", "2", "3", "2+sqrt(2)"},
			DistinctEigenvalues:              5,
			MinimalDegree:                    5,
			CharacteristicPolynomial:         "same as full real incidence Gram",
			MinimalPolynomial:                "same as full real incidence Gram",
			CharacteristicResidualToQ4:       math.Inf(1),
			MinimalResidualToQ4:              math.Inf(1),
			Q4ExactMatch:                     false,
			PromotableAsQ4Selector:           false,
			Reason:                           "A pure unit phase on one boundary column cancels in d^†d. It does not generate new q4 coefficients, and no native non-unit boundary weight is allowed here.",
			Verdict:                          StatusFailedMajoranaTwistNotCanonical,
		},
		{
			Name:                             "J-doubled oriented boundary Gram",
			Formula:                          "Δ_E ⊕ JΔ_EJ^{-1} on the ten J-doubled one-form edge slots",
			Domain:                           "ten J-doubled oriented edge slots",
			Dimension:                        JDoubledEdgeCount,
			Native:                           true,
			BoundaryDerived:                  true,
			HphiEndomorphism:                 false,
			CanonicalQuotientToHphi:          false,
			CompatibleWithChirality:          true,
			CompatibleWithJ:                  true,
			CompatibleWithFirstOrder:         true,
			OrientationChoiceAffectsSpectrum: false,
			OrientationSignsCancel:           true,
			Eigenvalues:                      []string{"1", "1", "2-sqrt(2)", "2-sqrt(2)", "2", "2", "3", "3", "2+sqrt(2)", "2+sqrt(2)"},
			DistinctEigenvalues:              5,
			MinimalDegree:                    5,
			CharacteristicPolynomial:         "[(x-1)*(x-2)*(x-3)*(x^2-4*x+2)]^2",
			MinimalPolynomial:                "(x-1)*(x-2)*(x-3)*(x^2-4*x+2)",
			CharacteristicResidualToQ4:       math.Inf(1),
			MinimalResidualToQ4:              math.Inf(1),
			Q4ExactMatch:                     false,
			PromotableAsQ4Selector:           false,
			Reason:                           "J-doubling preserves the oriented incidence data but only duplicates its spectrum; it does not create a four-real scalar q4 selector.",
			Verdict:                          StatusFailedFullOrientedPolynomialNotQ4,
		},
		{
			Name:                       "sealed q4 oriented-boundary companion quotient",
			Formula:                    "choose a four-dimensional edge quotient and place the q4 companion matrix by hand",
			Domain:                     "manual four-dimensional oriented edge quotient",
			Dimension:                  HphiRealDim,
			Native:                     false,
			Sealed:                     true,
			Circular:                   true,
			BoundaryDerived:            false,
			HphiEndomorphism:           true,
			CanonicalQuotientToHphi:    false,
			CompatibleWithChirality:    false,
			CompatibleWithJ:            false,
			CompatibleWithFirstOrder:   false,
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
			Reason:                     "The q4 polynomial can always be imposed by a companion matrix after choosing an arbitrary four-dimensional quotient; this is exactly the operation the gate forbids.",
			Verdict:                    StatusFailedNoNativeOrientedQ4Selector,
		},
	}

	native, nativeHphi, nativeQuartic, matches, invariant, sealed := 0, 0, 0, 0, 0, 0
	bestName := "none"
	bestResidual := math.Inf(1)
	for _, c := range candidates {
		if c.Native {
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
		if c.OrientationSignsCancel {
			invariant++
		}
		if c.Sealed || c.Circular {
			sealed++
		}
		if c.Native && c.MinimalResidualToQ4 < bestResidual {
			bestResidual = c.MinimalResidualToQ4
			bestName = c.Name
		}
	}
	return BoundarySieve{
		Executed:                    true,
		Candidates:                  candidates,
		NativeBoundaryOperatorCount: native,
		NativeHphiEndomorphismCount: nativeHphi,
		NativeQuarticCapacityCount:  nativeQuartic,
		CanonicalHphiQ4MatchCount:   matches,
		OrientationInvariantCount:   invariant,
		SealedOrManualQ4Count:       sealed,
		BestNativeCandidate:         bestName,
		BestNativeQ4Residual:        bestResidual,
		Verdict:                     "The oriented incidence lane is native and stricter than undirected adjacency, but d^T d/d^†d is orientation-sign invariant. The four Yukawa edge Gram is still pair-degenerate; the full incidence Gram has degree five on edge space, not q4 on H_phi; all four-dimensional q4 hits remain manual quotients.",
	}
}

func auditImpact(s BoundarySieve) Impact {
	return Impact{
		HphiQuarticIdentified:          false,
		NativeBoundaryOperatorFound:    s.NativeBoundaryOperatorCount > 0,
		CanonicalBoundaryQuotientFound: false,
		OrientedIncidenceLaneOpened:    true,
		YukawaCouplingsReduced:         false,
		ChargedModuliStart:             Gate372ChargedModuliDim,
		ChargedModuliResult:            Gate372ChargedModuliDim,
		FlavorFirewallPreserved:        true,
		HiggsLanePreserved:             true,
		Verdict:                        "Gate 403 opens a real oriented-incidence diagnostic but preserves the scalar/contact q4 obstruction, the Yukawa-coupling firewall, and the Gate-372 thirteen-moduli flavor firewall.",
	}
}

func auditFirewall(i Impact) FirewallAudit {
	return FirewallAudit{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaAmplitudesInserted: true, NoGaugeChargeFitReused: true, NoManualQ4HphiID: true, NoArbitraryBoundaryQuotient: true, NoCompanionOperatorPromoted: true, NoFlavorModuliReductionClaimed: i.FlavorFirewallPreserved, Verdict: "No empirical or circular scalar/flavor information was imported. Manual q4 companion and arbitrary four-mode quotients were quarantined rather than promoted."}
}

func nextStep(s BoundarySieve, i Impact) NextStep {
	return NextStep{Gate: 404, Title: "Canonical Edge-to-Hphi Quotient / Contact-Edge Intertwiner Sieve", Reason: "Orientation is now exhausted: signs and unit phases cancel in d^T d or d^†d. The remaining missing object is not another graph operator but a canonical quotient/intertwiner from the five/ten edge-slot space to the four-real scalar carrier H_phi, preferably derived from contact projectors, one-form support, J, and first-order data.", PrimaryTask: "Search for a native map Q: edge-slot space -> H_phi such that Q^T Δ_edge Q is a canonical H_phi endomorphism; reject any quotient chosen only to force q4."}
}

func truth(s BoundarySieve, i Impact) string {
	return "Gate 403 proves that chiral orientation is meaningful for bookkeeping but not sufficient as a q4 selector. The signed boundary operator d is native, yet its Laplacian d^T d is invariant under edge-orientation reversal up to signed edge-basis conjugacy. The four Yukawa oriented Gram remains pair-degenerate with minimal degree two; the full five-edge incidence Gram acquires radical spectra and degree five, but it lives on edge-slot space rather than H_phi and is polynomially disjoint from q4. Majorana unit-phase twists cancel in d^†d. Therefore the oriented-boundary route does not canonically identify H_phi with the contact quartic primary, does not reduce Yukawa couplings, and preserves the 13-moduli firewall. The next valid gate is a canonical edge-to-Hphi quotient/contact-edge intertwiner sieve, not another orientation choice."
}

func Statuses(a Analysis) []string {
	statuses := []string{
		StatusGate402Inherited,
		StatusOrientedBoundaryFormalized,
		StatusMajoranaOrientationAudited,
		StatusOrientedLaplacianConstructed,
		StatusFullIncidenceRadicalSpectrumFound,
		StatusFailedOrientationSignsCancel,
	}
	for _, c := range a.Sieve.Candidates {
		if c.Verdict != "" && !contains(statuses, c.Verdict) {
			statuses = append(statuses, c.Verdict)
		}
	}
	for _, status := range []string{
		StatusFailedNoNativeOrientedQ4Selector,
		StatusFailedNoCanonicalHphiQuarticID,
		StatusFailedNoYukawaCouplingReduction,
		StatusFirewallPreserved13Moduli,
	} {
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
	return fmt.Sprintf("executed=%t gate399_quaternionic_no_q4=%t gate400_no_mixed_q4=%t gate401_charge_weights_disjoint=%t gate402_graph_native=%t gate402_quartic_capacity=%t gate402_no_graph_q4=%t oneform_edges=%t J_edges=%d first_order_edge_graph=%t moduli_dim=%d no_empirical=%t verdict=%s", x.Executed, x.Gate399QuaternionicPolynomialNo, x.Gate400NoMixedEdgeQ4, x.Gate401ChargeWeightsDisjoint, x.Gate402UndirectedGraphNative, x.Gate402FullGraphQuarticCapacity, x.Gate402NoGraphQ4, x.Gate385OneFormEdges, x.Gate385JDoubledEdgeCount, x.Gate297FirstOrderEdgeGraph, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatQ4(x Q4Audit) string {
	return fmt.Sprintf("polynomial=%s degree=%d irreducible_over_Q=%t monic=%v roots=%s verdict=%s", x.Polynomial, x.Degree, x.IrreducibleOverQ, x.MonicCoefficients, x.RootSummary, x.Verdict)
}

func FormatArena(x BoundaryArena) string {
	return fmt.Sprintf("formalized=%t structural_edges=%d yukawa_edges=%d J_doubled=%d chiral_orientation=%t majorana_canonical=%t majorana_twist_possible=%t canonical_Hphi_quotient=%t uses_charge_weights=%t uses_yukawa_amplitudes=%t uses_masses=%t verdict=%s", x.Formalized, x.StructuralEdgeCount, x.YukawaEdgeCount, x.JDoubledEdgeCount, x.ChiralOrientationAvailable, x.MajoranaOrientationCanonical, x.MajoranaTwistPossible, x.HasCanonicalHphiQuotient, x.UsesGaugeChargeWeights, x.UsesYukawaAmplitudes, x.UsesObservedMasses, x.Verdict)
}

func FormatCandidate(c BoundaryCandidate) string {
	return fmt.Sprintf("name=%q domain=%q dim=%d native=%t sealed=%t circular=%t boundary=%t Hphi=%t canonical_quotient=%t chiral=%t J=%t first_order=%t majorana=%q orientation_affects_spectrum=%t signs_cancel=%t eigen=%v distinct=%d min_degree=%d char=%s min=%s char_residual_q4=%s min_residual_q4=%s pair_degenerate=%t quartic_capacity=%t q4_exact=%t promotable=%t verdict=%s reason=%s", c.Name, c.Domain, c.Dimension, c.Native, c.Sealed, c.Circular, c.BoundaryDerived, c.HphiEndomorphism, c.CanonicalQuotientToHphi, c.CompatibleWithChirality, c.CompatibleWithJ, c.CompatibleWithFirstOrder, c.MajoranaOrientationUsed, c.OrientationChoiceAffectsSpectrum, c.OrientationSignsCancel, c.Eigenvalues, c.DistinctEigenvalues, c.MinimalDegree, c.CharacteristicPolynomial, c.MinimalPolynomial, formatResidual(c.CharacteristicResidualToQ4), formatResidual(c.MinimalResidualToQ4), c.PairDegenerate, c.IrreducibleQuarticCapacity, c.Q4ExactMatch, c.PromotableAsQ4Selector, c.Verdict, c.Reason)
}

func FormatSieve(x BoundarySieve) string {
	parts := []string{fmt.Sprintf("executed=%t native_boundary=%d native_Hphi=%d native_quartic_capacity=%d canonical_q4_matches=%d orientation_invariant=%d sealed_manual=%d best_native=%s best_residual=%s verdict=%s", x.Executed, x.NativeBoundaryOperatorCount, x.NativeHphiEndomorphismCount, x.NativeQuarticCapacityCount, x.CanonicalHphiQ4MatchCount, x.OrientationInvariantCount, x.SealedOrManualQ4Count, x.BestNativeCandidate, formatResidual(x.BestNativeQ4Residual), x.Verdict)}
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return strings.Join(parts, "\n")
}

func FormatImpact(x Impact) string {
	return fmt.Sprintf("Hphi_q4_identified=%t native_boundary=%t canonical_boundary_quotient=%t oriented_lane_opened=%t yukawa_reduced=%t moduli_start=%d moduli_result=%d flavor_firewall=%t higgs_lane_preserved=%t verdict=%s", x.HphiQuarticIdentified, x.NativeBoundaryOperatorFound, x.CanonicalBoundaryQuotientFound, x.OrientedIncidenceLaneOpened, x.YukawaCouplingsReduced, x.ChargedModuliStart, x.ChargedModuliResult, x.FlavorFirewallPreserved, x.HiggsLanePreserved, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_ckm=%t no_pmns=%t no_yukawa_amplitudes=%t no_charge_fit=%t no_manual_q4=%t no_arbitrary_boundary_quotient=%t no_companion_promoted=%t no_moduli_reduction=%t verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoGaugeChargeFitReused, x.NoManualQ4HphiID, x.NoArbitraryBoundaryQuotient, x.NoCompanionOperatorPromoted, x.NoFlavorModuliReductionClaimed, x.Verdict)
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
